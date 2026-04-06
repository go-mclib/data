#!/usr/bin/env python3
"""record and replay keyboard/mouse macros for Minecraft packet capture.

requires: pip install pynput
note: on macOS, grant Accessibility permissions to your terminal app
(System Settings > Privacy & Security > Accessibility)

limitations:
  - mouse position is recorded as absolute screen coordinates;
    replay only works at the same resolution/window position
  - in-game camera control (raw mouse delta) may not replay correctly
    because Minecraft reads CGEvent deltas, not absolute positions;
    use keyboard controls or /tp commands for camera orientation

setup checklist:
  - minecraft in borderless fullscreen (maximized window)
  - GUI scale 4 in minecraft video settings
  - same screen resolution for record and replay
  - place a command block at spawn that runs: /say [macro] start
  - place a command block at the end that runs: /say [macro] stop
"""

import argparse
import json
import os
import subprocess
import sys
import threading
import time


# markers in server log (from /say commands in command blocks)
START_MARKER = "[macro] start"
STOP_MARKER = "[macro] stop"


def _check_accessibility():
    """check macOS Accessibility permission; exit with instructions if missing."""
    if sys.platform != "darwin":
        return
    try:
        import ctypes
        import ctypes.util

        path = ctypes.util.find_library("ApplicationServices")
        if not path:
            return
        lib = ctypes.cdll.LoadLibrary(path)
        lib.AXIsProcessTrusted.restype = ctypes.c_bool
        if lib.AXIsProcessTrusted():
            return
    except Exception:
        return

    print(
        "error: macOS Accessibility permission required\n"
        "\n"
        "  System Settings > Privacy & Security > Accessibility\n"
        "  -> add your terminal app, then restart it\n",
        file=sys.stderr,
    )
    subprocess.run(
        [
            "open",
            "x-apple.systempreferences:"
            "com.apple.preference.security?Privacy_Accessibility",
        ],
        capture_output=True,
    )
    sys.exit(1)


def _wait_for_log(log_path, marker, timeout=300):
    """tail log_path from current end until a line containing marker appears."""
    deadline = time.monotonic() + timeout

    while not os.path.exists(log_path):
        if time.monotonic() > deadline:
            return False
        time.sleep(0.5)

    with open(log_path) as f:
        f.seek(0, 2)
        while time.monotonic() < deadline:
            line = f.readline()
            if line:
                if marker in line:
                    return True
            else:
                time.sleep(0.1)

    return False


def _send_server_cmd(server_input, cmd):
    """send a command to the Minecraft server via stdin FIFO."""
    if not server_input:
        return
    try:
        with open(server_input, "w") as f:
            f.write(cmd + "\n")
    except OSError:
        pass


def _focus_game_window():
    """activate the Minecraft (java) window on macOS."""
    if sys.platform != "darwin":
        return
    subprocess.run(
        [
            "osascript", "-e",
            'tell application "System Events" to set frontmost of '
            '(first process whose name contains "java") to true',
        ],
        capture_output=True,
    )
    time.sleep(0.3)


_suppress_tap = None
_event_source = None

# magic value to distinguish our synthetic events from real hardware
_MAGIC_USER_DATA = 0x474F4D43  # "GOMC"


def _start_input_suppression(stopped_event):
    """start an active CGEventTap that suppresses real input during replay.

    blocks real hardware mouse AND keyboard events so they don't contaminate
    the replay. our synthetic events pass through: mouse events are marked
    with magic user data, keyboard events are identified by source PID
    (real hardware has PID 0, our synthetic events have our PID).

    F10 (vk=109) from real hardware stops the replay via stopped_event.
    """
    import Quartz

    global _suppress_tap
    info = {}
    ready = threading.Event()

    def run():
        def suppress_callback(_proxy, event_type, event, _refcon):
            if event_type == Quartz.kCGEventTapDisabledByTimeout:
                Quartz.CGEventTapEnable(info.get("tap"), True)
                return event

            # let our synthetic events through (marked with magic value)
            if Quartz.CGEventGetIntegerValueField(
                event, Quartz.kCGEventSourceUserData
            ) == _MAGIC_USER_DATA:
                return event

            # F10 (vk=109) from real keyboard → stop the replay
            if event_type == Quartz.kCGEventKeyDown:
                vk = Quartz.CGEventGetIntegerValueField(
                    event, Quartz.kCGKeyboardEventKeycode
                )
                if vk == 109:  # F10
                    stopped_event.set()
                    return None

            return None  # drop all other real hardware input

        mask = (
            Quartz.CGEventMaskBit(Quartz.kCGEventMouseMoved)
            | Quartz.CGEventMaskBit(Quartz.kCGEventLeftMouseDragged)
            | Quartz.CGEventMaskBit(Quartz.kCGEventRightMouseDragged)
            | Quartz.CGEventMaskBit(Quartz.kCGEventOtherMouseDragged)
            | Quartz.CGEventMaskBit(Quartz.kCGEventKeyDown)
            | Quartz.CGEventMaskBit(Quartz.kCGEventKeyUp)
            | Quartz.CGEventMaskBit(Quartz.kCGEventFlagsChanged)
        )

        tap = Quartz.CGEventTapCreate(
            Quartz.kCGHIDEventTap,
            Quartz.kCGHeadInsertEventTap,
            Quartz.kCGEventTapOptionDefault,
            mask,
            suppress_callback,
            None,
        )
        if tap is None:
            print("warning: could not create suppression tap", file=sys.stderr)
            ready.set()
            return

        source = Quartz.CFMachPortCreateRunLoopSource(None, tap, 0)
        loop = Quartz.CFRunLoopGetCurrent()
        Quartz.CFRunLoopAddSource(loop, source, Quartz.kCFRunLoopCommonModes)
        Quartz.CGEventTapEnable(tap, True)

        info["tap"] = tap
        info["loop"] = loop
        ready.set()

        Quartz.CFRunLoopRun()

    t = threading.Thread(target=run, daemon=True)
    t.start()
    ready.wait()
    _suppress_tap = info.get("tap")
    return info.get("loop")


def _stop_input_suppression(loop):
    """stop the suppression tap."""
    import Quartz

    global _suppress_tap
    if _suppress_tap is not None:
        Quartz.CGEventTapEnable(_suppress_tap, False)
        _suppress_tap = None
    if loop is not None:
        Quartz.CFRunLoopStop(loop)


def _post_mouse_move(x, y, dx, dy):
    """post a mouse move through the full window server pipeline.

    sets both absolute position (for GUI/inventory) and deltas (for in-game
    camera). the recorded x/y is correct for both modes: frozen center when
    cursor is captured, actual position when cursor is free.
    """
    import Quartz

    global _event_source
    if _event_source is None:
        _event_source = Quartz.CGEventSourceCreate(
            Quartz.kCGEventSourceStateHIDSystemState
        )
        Quartz.CGEventSourceSetUserData(_event_source, _MAGIC_USER_DATA)

    event = Quartz.CGEventCreateMouseEvent(
        _event_source, Quartz.kCGEventMouseMoved, (x, y),
        Quartz.kCGMouseButtonLeft,
    )

    # delta fields — what GLFW reads via [NSEvent deltaX/Y]
    Quartz.CGEventSetIntegerValueField(
        event, Quartz.kCGMouseEventDeltaX, int(dx)
    )
    Quartz.CGEventSetIntegerValueField(
        event, Quartz.kCGMouseEventDeltaY, int(dy)
    )
    # unaccelerated fields (170/171)
    Quartz.CGEventSetIntegerValueField(event, 170, int(dx))
    Quartz.CGEventSetIntegerValueField(event, 171, int(dy))

    # prevent event coalescing — each event delivered individually
    flags = Quartz.CGEventGetFlags(event)
    Quartz.CGEventSetFlags(event, flags | Quartz.kCGEventFlagMaskNonCoalesced)

    # post at HID level — goes through full window server pipeline
    Quartz.CGEventPost(Quartz.kCGHIDEventTap, event)


def _post_key_event(vk, key_down):
    """post a keyboard event through the full pipeline with magic marker."""
    import Quartz

    global _event_source
    if _event_source is None:
        _event_source = Quartz.CGEventSourceCreate(
            Quartz.kCGEventSourceStateHIDSystemState
        )
        Quartz.CGEventSourceSetUserData(_event_source, _MAGIC_USER_DATA)

    event = Quartz.CGEventCreateKeyboardEvent(_event_source, vk, key_down)
    Quartz.CGEventPost(Quartz.kCGHIDEventTap, event)


def _start_mouse_tap(events, elapsed_fn):
    """start a CGEventTap that records mouse events with raw deltas.

    pynput's on_move only gives absolute position, which is frozen when
    Minecraft captures the cursor. a CGEventTap lets us read the raw
    kCGMouseEventDeltaX/Y fields directly.

    the tap, source, and run loop must all live on the SAME thread.
    """
    import Quartz

    info = {}
    ready = threading.Event()

    def run():
        def callback(_proxy, event_type, event, _refcon):
            if event_type == Quartz.kCGEventTapDisabledByTimeout:
                Quartz.CGEventTapEnable(info.get("tap"), True)
                return event

            if event_type in (
                Quartz.kCGEventMouseMoved,
                Quartz.kCGEventLeftMouseDragged,
                Quartz.kCGEventRightMouseDragged,
                Quartz.kCGEventOtherMouseDragged,
            ):
                dx = Quartz.CGEventGetDoubleValueField(
                    event, Quartz.kCGMouseEventDeltaX
                )
                dy = Quartz.CGEventGetDoubleValueField(
                    event, Quartz.kCGMouseEventDeltaY
                )
                pos = Quartz.CGEventGetLocation(event)
                events.append(
                    {
                        "t": elapsed_fn(),
                        "type": "mm",
                        "x": pos.x,
                        "y": pos.y,
                        "dx": dx,
                        "dy": dy,
                    }
                )

            elif event_type in (
                Quartz.kCGEventLeftMouseDown,
                Quartz.kCGEventLeftMouseUp,
                Quartz.kCGEventRightMouseDown,
                Quartz.kCGEventRightMouseUp,
                Quartz.kCGEventOtherMouseDown,
                Quartz.kCGEventOtherMouseUp,
            ):
                pos = Quartz.CGEventGetLocation(event)
                pressed = event_type in (
                    Quartz.kCGEventLeftMouseDown,
                    Quartz.kCGEventRightMouseDown,
                    Quartz.kCGEventOtherMouseDown,
                )
                btn_map = {
                    Quartz.kCGEventLeftMouseDown: "left",
                    Quartz.kCGEventLeftMouseUp: "left",
                    Quartz.kCGEventRightMouseDown: "right",
                    Quartz.kCGEventRightMouseUp: "right",
                    Quartz.kCGEventOtherMouseDown: "middle",
                    Quartz.kCGEventOtherMouseUp: "middle",
                }
                events.append(
                    {
                        "t": elapsed_fn(),
                        "type": "mp" if pressed else "mr",
                        "x": pos.x,
                        "y": pos.y,
                        "btn": btn_map.get(event_type, "left"),
                    }
                )

            elif event_type == Quartz.kCGEventScrollWheel:
                pos = Quartz.CGEventGetLocation(event)
                sdx = Quartz.CGEventGetIntegerValueField(
                    event, Quartz.kCGScrollWheelEventDeltaAxis2
                )
                sdy = Quartz.CGEventGetIntegerValueField(
                    event, Quartz.kCGScrollWheelEventDeltaAxis1
                )
                events.append(
                    {
                        "t": elapsed_fn(),
                        "type": "ms",
                        "x": pos.x,
                        "y": pos.y,
                        "dx": sdx,
                        "dy": sdy,
                    }
                )

            return event

        mask = (
            (1 << Quartz.kCGEventMouseMoved)
            | (1 << Quartz.kCGEventLeftMouseDown)
            | (1 << Quartz.kCGEventLeftMouseUp)
            | (1 << Quartz.kCGEventRightMouseDown)
            | (1 << Quartz.kCGEventRightMouseUp)
            | (1 << Quartz.kCGEventOtherMouseDown)
            | (1 << Quartz.kCGEventOtherMouseUp)
            | (1 << Quartz.kCGEventLeftMouseDragged)
            | (1 << Quartz.kCGEventRightMouseDragged)
            | (1 << Quartz.kCGEventOtherMouseDragged)
            | (1 << Quartz.kCGEventScrollWheel)
        )
        # kCGEventTapDisabledByTimeout is delivered automatically — NOT in mask

        tap = Quartz.CGEventTapCreate(
            Quartz.kCGHIDEventTap,
            Quartz.kCGHeadInsertEventTap,
            Quartz.kCGEventTapOptionListenOnly,
            mask,
            callback,
            None,
        )
        if tap is None:
            print(
                "warning: could not create CGEventTap for mouse", file=sys.stderr
            )
            ready.set()
            return

        source = Quartz.CFMachPortCreateRunLoopSource(None, tap, 0)
        loop = Quartz.CFRunLoopGetCurrent()
        Quartz.CFRunLoopAddSource(loop, source, Quartz.kCFRunLoopCommonModes)
        Quartz.CGEventTapEnable(tap, True)

        info["tap"] = tap
        info["source"] = source
        info["loop"] = loop
        ready.set()

        Quartz.CFRunLoopRun()

    t = threading.Thread(target=run, daemon=True)
    t.start()
    ready.wait()

    return info.get("tap"), info.get("source"), info.get("loop")


def _trigger_capture(capture_trigger):
    """create the trigger file so the proxy starts capturing."""
    if not capture_trigger:
        return
    try:
        os.makedirs(os.path.dirname(capture_trigger), exist_ok=True)
        open(capture_trigger, "w").close()
    except OSError:
        pass


def record(output_file, stop_key="f10", server_log=None, server_input=None, capture_trigger=None):
    """record keyboard and mouse events."""
    _check_accessibility()
    try:
        from pynput import keyboard
    except ImportError:
        print("pynput is required: pip install pynput", file=sys.stderr)
        sys.exit(1)

    if server_log:
        print(f"waiting for '{START_MARKER}' in server log...")
        if not _wait_for_log(server_log, START_MARKER):
            print("error: timed out waiting for start marker", file=sys.stderr)
            sys.exit(1)
        _trigger_capture(capture_trigger)
        _send_server_cmd(server_input, "say macro started")
        print("recording!")
    else:
        print(f"recording... press {stop_key.upper()} to stop")

    events = []
    start_time = time.monotonic()
    stop_key_name = stop_key.lower()

    def elapsed():
        return round(time.monotonic() - start_time, 4)

    def key_str(key):
        if isinstance(key, keyboard.Key):
            return key.name
        if isinstance(key, keyboard.KeyCode):
            if key.char is not None:
                return key.char
            return f"<{key.vk}>"
        return str(key)

    def key_event(etype, key):
        name = key_str(key)
        ev = {"t": elapsed(), "type": etype, "key": name}
        # extract vk from KeyCode directly, or from Key.value for special keys
        vk = getattr(key, "vk", None)
        if vk is None and isinstance(key, keyboard.Key):
            vk = getattr(key.value, "vk", None)
        if vk is not None:
            ev["vk"] = vk
        return ev

    def on_key_press(key):
        name = key_str(key)
        if name.lower() == stop_key_name:
            return False
        events.append(key_event("kp", key))

    def on_key_release(key):
        name = key_str(key)
        if name.lower() == stop_key_name:
            return
        events.append(key_event("kr", key))

    # on macOS, use a CGEventTap for mouse events to capture raw deltas
    # (pynput only gives absolute position, which is frozen when cursor is captured)
    mouse_tap = None
    mouse_loop = None
    if sys.platform == "darwin":
        mouse_tap, _mouse_source, mouse_loop = _start_mouse_tap(events, elapsed)

    kb_listener = keyboard.Listener(on_press=on_key_press, on_release=on_key_release)
    kb_listener.start()

    # watch for stop marker in server log
    if server_log:

        def watch_stop():
            if _wait_for_log(server_log, STOP_MARKER):
                print(f"'{STOP_MARKER}' received")
                _send_server_cmd(server_input, "say macro stopped")
                kb_listener.stop()

        threading.Thread(target=watch_stop, daemon=True).start()

    # join with timeout so ctrl+c (KeyboardInterrupt) is not swallowed
    try:
        while kb_listener.is_alive():
            kb_listener.join(timeout=0.5)
    except KeyboardInterrupt:
        pass

    kb_listener.stop()
    if mouse_loop is not None:
        import Quartz

        Quartz.CFRunLoopStop(mouse_loop)
        if mouse_tap is not None:
            Quartz.CGEventTapEnable(mouse_tap, False)

    with open(output_file, "w") as f:
        json.dump(events, f, separators=(",", ":"))

    duration = elapsed()
    print(f"recorded {len(events)} events ({duration:.1f}s) -> {output_file}")


def replay(input_file, speed=1.0, server_log=None, server_input=None, capture_trigger=None):
    """replay recorded events from file."""
    _check_accessibility()
    try:
        from pynput import keyboard, mouse
    except ImportError:
        print("pynput is required: pip install pynput", file=sys.stderr)
        sys.exit(1)

    with open(input_file) as f:
        events = json.load(f)

    if not events:
        print("no events to replay")
        return

    if server_log:
        print(f"waiting for '{START_MARKER}' in server log...")
        if not _wait_for_log(server_log, START_MARKER):
            print("error: timed out waiting for start marker", file=sys.stderr)
            sys.exit(1)
        _trigger_capture(capture_trigger)
        _send_server_cmd(server_input, "say macro started")
        print("replaying!")

    def str_to_key(s, vk=None):
        # prefer virtual keycode when available (stored during recording)
        if vk is not None:
            return keyboard.KeyCode.from_vk(vk)
        try:
            return keyboard.Key[s]
        except (KeyError, ValueError):
            pass
        if len(s) == 1:
            return keyboard.KeyCode.from_char(s)
        if s.startswith("<") and s.endswith(">"):
            return keyboard.KeyCode.from_vk(int(s[1:-1]))
        return keyboard.KeyCode.from_char(s)

    def str_to_button(s):
        return mouse.Button[s]

    kb = keyboard.Controller()
    ms = mouse.Controller()

    # on macOS, use Quartz CGEvents for mouse moves so Minecraft gets
    # delta fields (needed for camera control when cursor is captured)
    use_quartz = sys.platform == "darwin"
    if use_quartz:
        try:
            import Quartz  # noqa: F401 — verify import, actual use in _post_mouse_move
        except ImportError:
            use_quartz = False

    duration = events[-1]["t"] / speed
    print(
        f"replaying {len(events)} events ({duration:.1f}s at {speed}x)... "
        "F10 or ctrl+c to abort"
    )

    _focus_game_window()

    # watch for stop marker in server log, or F10 from real keyboard
    stopped = threading.Event()

    suppress_loop = None
    if use_quartz:
        import Quartz

        # suppress real mouse + keyboard so they don't contaminate the replay;
        # our synthetic events pass through; F10 triggers stopped
        suppress_loop = _start_input_suppression(stopped)
    if server_log:

        def watch_stop():
            if _wait_for_log(server_log, STOP_MARKER):
                print(f"'{STOP_MARKER}' received")
                _send_server_cmd(server_input, "say macro stopped")
                stopped.set()

        threading.Thread(target=watch_stop, daemon=True).start()

    start = time.monotonic()

    def precise_wait(deadline):
        """sleep + busy-spin for sub-ms accuracy."""
        remaining = deadline - time.monotonic()
        if remaining > 0.002:
            time.sleep(remaining - 0.001)  # sleep most of it
        while time.monotonic() < deadline:  # busy-spin the rest
            pass

    try:
        for ev in events:
            if stopped.is_set():
                break
            target = start + ev["t"] / speed
            precise_wait(target)

            t = ev["type"]
            if t == "kp":
                vk = ev.get("vk")
                if use_quartz and vk is not None:
                    _post_key_event(vk, True)
                else:
                    kb.press(str_to_key(ev["key"], vk))
            elif t == "kr":
                vk = ev.get("vk")
                if use_quartz and vk is not None:
                    _post_key_event(vk, False)
                else:
                    kb.release(str_to_key(ev["key"], vk))
            elif t == "mm":
                if use_quartz and "dx" in ev:
                    _post_mouse_move(
                        ev.get("x", 0), ev.get("y", 0),
                        ev["dx"], ev["dy"],
                    )
                else:
                    ms.position = (ev["x"], ev["y"])
            elif t == "mp":
                ms.position = (ev["x"], ev["y"])
                ms.press(str_to_button(ev["btn"]))
            elif t == "mr":
                ms.position = (ev["x"], ev["y"])
                ms.release(str_to_button(ev["btn"]))
            elif t == "ms":
                ms.position = (ev["x"], ev["y"])
                ms.scroll(ev["dx"], ev["dy"])
    except KeyboardInterrupt:
        print("\naborted")
    finally:
        if suppress_loop is not None:
            _stop_input_suppression(suppress_loop)

    if not stopped.is_set():
        _send_server_cmd(server_input, "say macro stopped")
    print("replay complete")


def main():
    parser = argparse.ArgumentParser(description="record/replay keyboard+mouse macros")
    sub = parser.add_subparsers(dest="cmd", required=True)

    log_help = "server log to watch for [macro] start/stop markers"
    input_help = "server stdin FIFO for sending commands (say macro started/stopped)"
    trigger_help = "file to create when macro starts (signals proxy to begin capture)"

    rec = sub.add_parser("record", help="record a macro")
    rec.add_argument("output", help="output file (.json)")
    rec.add_argument(
        "--stop-key", default="f10", help="key to stop recording (default: f10)"
    )
    rec.add_argument("--server-log", metavar="LOG", help=log_help)
    rec.add_argument("--server-input", metavar="FIFO", help=input_help)
    rec.add_argument("--capture-trigger", metavar="PATH", help=trigger_help)

    rep = sub.add_parser("replay", help="replay a macro")
    rep.add_argument("input", help="input file (.json)")
    rep.add_argument(
        "--speed", type=float, default=1.0, help="playback speed (default: 1.0)"
    )
    rep.add_argument("--server-log", metavar="LOG", help=log_help)
    rep.add_argument("--server-input", metavar="FIFO", help=input_help)
    rep.add_argument("--capture-trigger", metavar="PATH", help=trigger_help)

    sub.add_parser("check", help="verify macOS Accessibility permissions")

    args = parser.parse_args()
    if args.cmd == "check":
        _check_accessibility()
        print("ok: Accessibility permissions granted")
    elif args.cmd == "record":
        record(
            args.output, args.stop_key,
            args.server_log, args.server_input, args.capture_trigger,
        )
    else:
        replay(
            args.input, args.speed,
            args.server_log, args.server_input, args.capture_trigger,
        )


if __name__ == "__main__":
    main()

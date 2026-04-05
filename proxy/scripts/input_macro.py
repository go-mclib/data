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
"""

import abc
import argparse
import json
import sys
import time

# minimum interval between mouse move events (seconds)
MOUSE_MOVE_INTERVAL = 0.01


def record(output_file, stop_key="f10"):
    """record keyboard and mouse events until stop key is pressed."""
    try:
        from pynput import keyboard, mouse
    except ImportError:
        print("pynput is required: pip install pynput", file=sys.stderr)
        sys.exit(1)

    events = []
    start_time = time.monotonic()
    last_mouse_move = 0.0
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

    def on_key_press(key):
        name = key_str(key)
        if name.lower() == stop_key_name:
            return False
        events.append({"t": elapsed(), "type": "kp", "key": name})

    def on_key_release(key):
        name = key_str(key)
        if name.lower() == stop_key_name:
            return
        events.append({"t": elapsed(), "type": "kr", "key": name})

    def on_move(x, y):
        nonlocal last_mouse_move
        now = time.monotonic()
        if now - last_mouse_move < MOUSE_MOVE_INTERVAL:
            return
        last_mouse_move = now
        events.append({"t": elapsed(), "type": "mm", "x": x, "y": y})

    def on_click(x, y, button, pressed):
        events.append(
            {
                "t": elapsed(),
                "type": "mp" if pressed else "mr",
                "x": x,
                "y": y,
                "btn": button.name,
            }
        )

    def on_scroll(x, y, dx, dy):
        events.append(
            {
                "t": elapsed(),
                "type": "ms",
                "x": x,
                "y": y,
                "dx": dx,
                "dy": dy,
            }
        )

    print(f"recording... press {stop_key.upper()} to stop")

    kb_listener = keyboard.Listener(on_press=on_key_press, on_release=on_key_release)
    ms_listener = mouse.Listener(
        on_move=on_move, on_click=on_click, on_scroll=on_scroll
    )

    kb_listener.start()
    ms_listener.start()
    kb_listener.join()
    ms_listener.stop()

    with open(output_file, "w") as f:
        json.dump(events, f, separators=(",", ":"))

    duration = elapsed()
    print(f"recorded {len(events)} events ({duration:.1f}s) -> {output_file}")


def replay(input_file, speed=1.0):
    """replay recorded events from file."""
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

    def str_to_key(s):
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

    duration = events[-1]["t"] / speed
    print(
        f"replaying {len(events)} events ({duration:.1f}s at {speed}x)... "
        "ctrl+c to abort"
    )

    start = time.monotonic()

    try:
        for ev in events:
            target = ev["t"] / speed
            wait = target - (time.monotonic() - start)
            if wait > 0:
                time.sleep(wait)

            t = ev["type"]
            if t == "kp":
                kb.press(str_to_key(ev["key"]))
            elif t == "kr":
                kb.release(str_to_key(ev["key"]))
            elif t == "mm":
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
        return

    print("replay complete")


def main():
    parser = argparse.ArgumentParser(description="record/replay keyboard+mouse macros")
    sub = parser.add_subparsers(dest="cmd", required=True)

    rec = sub.add_parser("record", help="record a macro")
    rec.add_argument("output", help="output file (.json)")
    rec.add_argument(
        "--stop-key", default="f10", help="key to stop recording (default: f10)"
    )

    rep = sub.add_parser("replay", help="replay a macro")
    rep.add_argument("input", help="input file (.json)")
    rep.add_argument(
        "--speed", type=float, default=1.0, help="playback speed (default: 1.0)"
    )

    args = parser.parse_args()
    if args.cmd == "record":
        record(args.output, args.stop_key)
    else:
        replay(args.input, args.speed)


if __name__ == "__main__":
    main()

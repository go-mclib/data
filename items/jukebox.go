package items

import (
	"encoding/json"
	"fmt"

	ns "github.com/go-mclib/protocol/javaproto/wire"

	"github.com/go-mclib/data/registries"
)

// JukeboxSong is the behavior-bearing data selected by a jukebox_playable component.
type JukeboxSong struct {
	SoundEvent       string
	LengthInSeconds  float32
	ComparatorOutput int32
	RegistryID       int32
}

// JukeboxSongFromSlot resolves the jukebox song selected by a slot's effective components.
func JukeboxSongFromSlot(slot ns.Slot) (JukeboxSong, bool, error) {
	stack, err := FromSlot(slot)
	if err != nil {
		return JukeboxSong{}, false, err
	}
	if stack.Components == nil {
		return JukeboxSong{}, false, nil
	}
	if stack.Components.JukeboxPlayable != "" {
		song, err := registeredJukeboxSong(stack.Components.JukeboxPlayable)
		return song, err == nil, err
	}
	raw, ok := stack.Components.rawPassthrough[ComponentJukeboxPlayable]
	if !ok {
		return JukeboxSong{}, false, nil
	}
	song, err := directJukeboxSong(raw)
	return song, err == nil, err
}

func registeredJukeboxSong(name string) (JukeboxSong, error) {
	raw := registries.SynchronizedRegistryData["minecraft:jukebox_song"][name]
	if len(raw) == 0 {
		return JukeboxSong{}, fmt.Errorf("unknown jukebox song %q", name)
	}
	var value struct {
		SoundEvent       string  `json:"sound_event"`
		LengthInSeconds  float32 `json:"length_in_seconds"`
		ComparatorOutput int32   `json:"comparator_output"`
	}
	if err := json.Unmarshal(raw, &value); err != nil {
		return JukeboxSong{}, fmt.Errorf("decode jukebox song %q: %w", name, err)
	}
	registryID := registries.SynchronizedEntries["minecraft:jukebox_song"].Get(name)
	return JukeboxSong{SoundEvent: value.SoundEvent, LengthInSeconds: value.LengthInSeconds, ComparatorOutput: value.ComparatorOutput, RegistryID: registryID}, nil
}

func directJukeboxSong(raw []byte) (JukeboxSong, error) {
	reader := ns.NewReader(raw)
	holder, err := reader.ReadVarInt()
	if err != nil {
		return JukeboxSong{}, err
	}
	if holder != 0 {
		return JukeboxSong{}, fmt.Errorf("invalid direct jukebox song holder %d", holder)
	}
	sound, err := readSoundHolder(reader)
	if err != nil {
		return JukeboxSong{}, err
	}
	if _, err := reader.ReadNBT(); err != nil {
		return JukeboxSong{}, fmt.Errorf("decode jukebox song description: %w", err)
	}
	length, err := reader.ReadFloat32()
	if err != nil {
		return JukeboxSong{}, err
	}
	output, err := reader.ReadVarInt()
	if err != nil {
		return JukeboxSong{}, err
	}
	return JukeboxSong{SoundEvent: sound.name, LengthInSeconds: float32(length), ComparatorOutput: int32(output), RegistryID: -1}, nil
}

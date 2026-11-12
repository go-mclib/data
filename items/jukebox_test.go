package items_test

import (
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestJukeboxSongFromSlot(t *testing.T) {
	registered := ns.NewSlot(ns.VarInt(items.ID("minecraft:music_disc_cat")), 1)
	song, ok, err := items.JukeboxSongFromSlot(registered)
	if err != nil || !ok || song.SoundEvent != "minecraft:music_disc.cat" || song.LengthInSeconds != 185 || song.ComparatorOutput != 2 {
		t.Fatalf("registered song = %+v, %v, %v", song, ok, err)
	}

	writer := ns.NewWriter()
	writer.WriteVarInt(0)
	writer.WriteVarInt(ns.VarInt(registries.SoundEvent.HolderID("minecraft:music_disc.cat")))
	if err := nbt.NewWriterTo(writer).WriteTag(nbt.String("Custom song"), "", nbt.FormatNetwork); err != nil {
		t.Fatal(err)
	}
	writer.WriteFloat32(2.5)
	writer.WriteVarInt(7)
	custom := ns.NewSlot(ns.VarInt(items.ID("minecraft:stick")), 1)
	custom.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentJukeboxPlayable, Data: writer.Bytes()}}
	song, ok, err = items.JukeboxSongFromSlot(custom)
	if err != nil || !ok || song.SoundEvent != "minecraft:music_disc.cat" || song.LengthInSeconds != 2.5 || song.ComparatorOutput != 7 {
		t.Fatalf("direct song = %+v, %v, %v", song, ok, err)
	}
}

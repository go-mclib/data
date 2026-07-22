package items_test

import (
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestJukeboxSongFromSlot(t *testing.T) {
	registered := ns.NewSlot(ns.VarInt(items.ItemID("minecraft:music_disc_cat")), 1)
	song, ok, err := items.JukeboxSongFromSlot(registered)
	if err != nil || !ok || song.SoundEvent != "minecraft:music_disc.cat" || song.LengthInSeconds != 185 || song.ComparatorOutput != 2 {
		t.Fatalf("registered song = %+v, %v, %v", song, ok, err)
	}

	writer := ns.NewWriter()
	writer.WriteVarInt(0)
	writer.WriteVarInt(ns.VarInt(registries.SoundEvent.Get("minecraft:music_disc.cat") + 1))
	if err := nbt.NewWriterTo(writer).WriteTag(nbt.String("Custom song"), "", true); err != nil {
		t.Fatal(err)
	}
	writer.WriteFloat32(2.5)
	writer.WriteVarInt(7)
	custom := ns.NewSlot(ns.VarInt(items.ItemID("minecraft:stick")), 1)
	custom.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentJukeboxPlayable, Data: writer.Bytes()}}
	song, ok, err = items.JukeboxSongFromSlot(custom)
	if err != nil || !ok || song.SoundEvent != "minecraft:music_disc.cat" || song.LengthInSeconds != 2.5 || song.ComparatorOutput != 7 {
		t.Fatalf("direct song = %+v, %v, %v", song, ok, err)
	}
}

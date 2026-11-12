package packets_test

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/packets"
	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestSoundPacketsRoundTrip(t *testing.T) {
	roundTrip(t, &packets.S2CSound{
		SoundEvent: ns.RegistrySoundEvent(42), SoundCategory: 4,
		EffectPositionX: 80, EffectPositionY: 512, EffectPositionZ: -24,
		Volume: 1, Pitch: 0.8, Seed: 123,
	}, &packets.S2CSound{})

	roundTrip(t, &packets.S2CSoundEntity{
		SoundEvent: ns.SoundEvent{
			Direct: true, SoundID: "minecraft:custom.sound",
			FixedRange: ns.Some(ns.Float32(32)),
		},
		SoundCategory: 7, EntityID: 99, Volume: 2, Pitch: 1, Seed: 456,
	}, &packets.S2CSoundEntity{})
}

func TestExplosionFieldsAreTrailingData(t *testing.T) {
	roundTrip(t, &packets.S2CExplode{
		X: 1, Y: 2, Z: 3,
		Data:      []byte{0x40, 0x80, 0x00, 0x00, 0, 0, 0, 0, 0, 0x1d, 0xbc, 0x05, 0},
		PlaySound: true,
	}, &packets.S2CExplode{})
}

func TestItemParticleOptionsPreserveFollowingFields(t *testing.T) {
	// An unframed item template with a variable-length component payload must
	// end before the limiter flags and coordinates, not at packet EOF.
	options := ns.NewWriter()
	options.WriteVarInt(ns.VarInt(items.ID("minecraft:diamond_sword")))
	options.WriteVarInt(2)
	options.WriteVarInt(1)
	options.WriteVarInt(1)
	options.WriteVarInt(items.ComponentDamage)
	options.WriteVarInt(300)
	options.WriteVarInt(items.ComponentCustomName)

	wire := ns.NewWriter()
	id := ns.VarInt(registries.ParticleType.Get("minecraft:item"))
	wire.WriteVarInt(id)
	wire.WriteFixedByteArray(options.Bytes())
	wire.WriteBool(true)
	wire.WriteBool(true)
	wire.WriteFloat64(1)
	wire.WriteFloat64(2)
	wire.WriteFloat64(3)
	wire.WriteFloat32(4)
	wire.WriteFloat32(5)
	wire.WriteFloat32(6)
	wire.WriteFloat32(7)
	wire.WriteFloat32(8)
	wire.WriteFloat32(9)
	wire.WriteVarInt(300)
	wire.WriteVarInt(1)
	reader := bytes.NewReader(append(append([]byte(nil), wire.Bytes()...), 0x5a))
	var packet packets.S2CLevelParticles
	if err := packet.Read(ns.NewReaderFrom(reader)); err != nil {
		t.Fatal(err)
	}
	if packet.ParticleID != id || !bytes.Equal(packet.Data, options.Bytes()) || !bool(packet.LongDistance) || !bool(packet.AlwaysVisible) || packet.X != 1 || packet.Y != 2 || packet.Z != 3 || packet.OffsetX != 4 || packet.OffsetY != 5 || packet.OffsetZ != 6 || packet.MaxSpeedX != 7 || packet.MaxSpeedY != 8 || packet.MaxSpeedZ != 9 || packet.ParticleCount != 300 || packet.RandomizationType != packets.ParticleRandomizationAlternative {
		t.Fatalf("item particle fields = %+v", packet)
	}
	if sentinel, err := reader.ReadByte(); err != nil || sentinel != 0x5a || reader.Len() != 0 {
		t.Fatalf("particle payload boundary: sentinel=%x, err=%v", sentinel, err)
	}
	out := ns.NewWriter()
	if err := packet.Write(out); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out.Bytes(), wire.Bytes()) {
		t.Fatalf("particle wire = %x; want %x", out.Bytes(), wire.Bytes())
	}
}

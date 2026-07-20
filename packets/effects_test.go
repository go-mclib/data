package packets_test

import (
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
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
		SoundCategory: 7, EntityId: 99, Volume: 2, Pitch: 1, Seed: 456,
	}, &packets.S2CSoundEntity{})
}

func TestExplosionFieldsAreTrailingData(t *testing.T) {
	roundTrip(t, &packets.S2CExplode{
		X: 1, Y: 2, Z: 3,
		Data: []byte{0x40, 0x80, 0x00, 0x00, 0, 0, 0, 0, 0, 0x1d, 0xbc, 0x05, 0},
	}, &packets.S2CExplode{})
}

func TestParticleOptionsAreTrailingData(t *testing.T) {
	in := &packets.S2CLevelParticles{
		LongDistance: true, X: 1, Y: 2, Z: 3,
		OffsetX: 0.1, OffsetY: 0.2, OffsetZ: 0.3,
		MaxSpeed: 0.5, ParticleCount: 4, ParticleId: 7,
		Data: []byte{0x3f, 0x80, 0x00, 0x00},
	}
	roundTrip(t, in, &packets.S2CLevelParticles{})
}

package items_test

import (
	"math"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// slotWithComponent encodes a one-component slot patch with raw component bytes.
func slotWithComponent(itemID, componentID int32, data []byte) []byte {
	w := ns.NewWriter()
	w.WriteVarInt(1)
	w.WriteVarInt(ns.VarInt(itemID))
	w.WriteVarInt(1)
	w.WriteVarInt(0)
	w.WriteVarInt(ns.VarInt(componentID))
	w.Write(data)
	return w.Bytes()
}

func decodeSlot(data []byte) (ns.Slot, error) {
	return ns.NewReader(data).ReadSlot(items.Decoder())
}

func TestHostileHolderSetAndListCountsAreRejected(t *testing.T) {
	stone := items.ID("minecraft:stone")
	huge := ns.VarInt(math.MaxInt32)

	mobVisibility := ns.NewWriter()
	mobVisibility.WriteVarInt(huge) // entity type holder set marker: huge direct list
	blocksAttacks := ns.NewWriter()
	blocksAttacks.WriteFloat32(0)
	blocksAttacks.WriteFloat32(1)
	blocksAttacks.WriteVarInt(huge) // damage reduction count
	toolRule := ns.NewWriter()
	toolRule.WriteVarInt(1) // one rule
	toolRule.WriteVarInt(-5)
	lore := ns.NewWriter()
	lore.WriteVarInt(-1)
	enchantments := ns.NewWriter()
	enchantments.WriteVarInt(-1)
	bees := ns.NewWriter()
	bees.WriteVarInt(-1)

	for name, test := range map[string]struct {
		id   int32
		data []byte
	}{
		"mob_visibility": {items.ComponentMobVisibility, mobVisibility.Bytes()},
		"blocks_attacks": {items.ComponentBlocksAttacks, blocksAttacks.Bytes()},
		"tool rule":      {items.ComponentTool, toolRule.Bytes()},
		"negative lore":  {items.ComponentLore, lore.Bytes()},
		"negative enchs": {items.ComponentEnchantments, enchantments.Bytes()},
		"negative bees":  {items.ComponentBees, bees.Bytes()},
	} {
		if _, err := decodeSlot(slotWithComponent(stone, test.id, test.data)); err == nil {
			t.Errorf("%s: hostile payload decoded without error", name)
		}
	}
}

func TestDelimitedComponentLengthIsBounded(t *testing.T) {
	for _, length := range []ns.VarInt{-1, math.MaxInt32} {
		w := ns.NewWriter()
		w.WriteVarInt(1)
		w.WriteVarInt(ns.VarInt(items.ID("minecraft:stone")))
		w.WriteVarInt(1)
		w.WriteVarInt(0)
		w.WriteVarInt(ns.VarInt(items.ComponentDamage))
		w.WriteVarInt(length)
		if _, err := ns.NewReader(w.Bytes()).ReadSlot(items.DecoderDelimited()); err == nil {
			t.Errorf("component length %d accepted", length)
		}
	}
}

func TestOutOfRangeRarityDecodesAsCommon(t *testing.T) {
	w := ns.NewWriter()
	w.WriteVarInt(-1)
	slot, err := decodeSlot(slotWithComponent(items.ID("minecraft:stone"), items.ComponentRarity, w.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatal(err)
	}
	if stack.Components.Rarity != "common" {
		t.Fatalf("rarity = %q, want common (vanilla OutOfBoundsStrategy.ZERO)", stack.Components.Rarity)
	}
}

func TestStackEquippableSlotUnknownItem(t *testing.T) {
	if slot, ok := items.StackEquippableSlot(ns.NewSlot(math.MaxInt32, 1)); ok {
		t.Fatalf("unknown item equippable slot = %q", slot)
	}
}

func TestResolvableProfileWireFormat(t *testing.T) {
	w := ns.NewWriter()
	w.WriteBool(true) // complete GameProfile
	w.WriteUUID(ns.UUID{1, 2, 3})
	w.WriteString("Notch")
	w.WriteVarInt(1) // properties
	w.WriteString("textures")
	w.WriteString("value")
	w.WriteBool(true)
	w.WriteString("signature")
	w.WriteBool(true) // body texture
	w.WriteIdentifier("minecraft:entity/player/wide/steve")
	w.WriteBool(false) // cape
	w.WriteBool(false) // elytra
	w.WriteBool(true)  // model type
	w.WriteBool(true)  // slim
	profile := w.Bytes()

	data := append(slotWithComponent(items.ID("minecraft:player_head"), items.ComponentProfile, profile), 0x7f)
	reader := ns.NewReader(data)
	slot, err := reader.ReadSlot(items.Decoder())
	if err != nil {
		t.Fatal(err)
	}
	if trailer, err := reader.ReadByte(); err != nil || trailer != 0x7f {
		t.Fatalf("profile decode consumed the wrong number of bytes (trailer %#x, %v)", trailer, err)
	}
	if len(slot.Components.Add) != 1 || string(slot.Components.Add[0].Data) != string(profile) {
		t.Fatalf("profile payload = %x, want %x", slot.Components.Add[0].Data, profile)
	}

	partial := ns.NewWriter()
	partial.WriteBool(false) // partial profile
	partial.WriteBool(true)
	partial.WriteString("Notch")
	partial.WriteBool(false) // no UUID
	partial.WriteVarInt(17)  // more properties than vanilla allows
	if _, err := decodeSlot(slotWithComponent(items.ID("minecraft:player_head"), items.ComponentProfile, partial.Bytes())); err == nil {
		t.Fatal("profile with 17 properties accepted")
	}
}

func TestStatusEffectAmplifierIsUnsignedByte(t *testing.T) {
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:potion")), 1)
	patch := nbt.Compound{"minecraft:potion_contents": nbt.Compound{
		"custom_effects": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"id": nbt.String("minecraft:speed"), "amplifier": nbt.Byte(-56), "duration": nbt.Int(100)},
			nbt.Compound{"id": nbt.String("minecraft:haste"), "amplifier": nbt.Byte(1), "duration": nbt.Int(100)},
		}},
	}}
	applied, err := items.ApplyNBTComponents(slot, patch)
	if err != nil {
		t.Fatal(err)
	}
	contents, err := items.ReadPotionContents(applied)
	if err != nil {
		t.Fatal(err)
	}
	if len(contents.Effects) != 2 || contents.Effects[0].Amplifier != 200 || contents.Effects[1].Amplifier != 1 {
		t.Fatalf("effects = %+v", contents.Effects)
	}
	effects := items.ComponentsToNBT(applied).GetCompound("minecraft:potion_contents").GetList("custom_effects")
	if amplifier, ok := effects.Elements[0].(nbt.Compound)["amplifier"].(nbt.Byte); !ok || uint8(amplifier) != 200 {
		t.Fatalf("written amplifier = %#v, want unsigned byte 200", effects.Elements[0].(nbt.Compound)["amplifier"])
	}
}

func TestSpearUseEffectsDefaults(t *testing.T) {
	defaults := items.DefaultComponents(items.ID("minecraft:iron_spear"))
	if defaults == nil || defaults.UseEffects == nil {
		t.Fatal("spear has no use_effects default")
	}
	if !defaults.UseEffects.CanSprint || defaults.UseEffects.SpeedMultiplier != 1 {
		t.Fatalf("spear use_effects = %+v", *defaults.UseEffects)
	}
	stack := items.NewStack(items.ID("minecraft:iron_spear"), 1)
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatal(err)
	}
	if len(slot.Components.Add) != 0 || len(slot.Components.Remove) != 0 {
		t.Fatalf("default spear encodes a patch: %+v", slot.Components)
	}
}

func TestPatchRemovalOverridesAddition(t *testing.T) {
	slot := items.WithStackDamage(ns.NewSlot(ns.VarInt(items.ID("minecraft:iron_sword")), 1), 7)
	slot.Components.Remove = []ns.VarInt{items.ComponentDamage}
	if damage := items.StackDamage(slot); damage != 0 {
		t.Fatalf("removed damage = %d", damage)
	}
	slot = items.WithStackDamage(slot, 3)
	if damage := items.StackDamage(slot); damage != 3 || len(slot.Components.Remove) != 0 {
		t.Fatalf("replaced damage = %d, removals %v", damage, slot.Components.Remove)
	}
}

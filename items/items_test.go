package items_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestFormattedLoreRetainsNestedTextAndWireStyle(t *testing.T) {
	writer := ns.NewWriter()
	writer.WriteVarInt(2)
	for _, line := range []ns.TextComponent{
		{Extra: []ns.TextComponent{{Text: "Spend ", Color: "gray"}, {Extra: []ns.TextComponent{{Text: "10 souls", Color: "gold"}}}}},
		{Text: "Right click ", Extra: []ns.TextComponent{{Text: "to redeem", Color: "yellow"}}},
	} {
		if err := writer.WriteTextComponent(line); err != nil {
			t.Fatal(err)
		}
	}
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:recovery_compass")), 1)
	slot.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentLore, Data: writer.Bytes()}}
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"Spend 10 souls", "Right click to redeem"}; !reflect.DeepEqual(stack.Components.Lore, want) {
		t.Fatalf("formatted lore = %q, want %q", stack.Components.Lore, want)
	}
	restored, err := stack.ToSlot()
	if err != nil || !items.SameStack(slot, restored) {
		t.Fatalf("styled lore changed on replay: %v, %v", restored, err)
	}
}

func TestReadPotionContentsCustomFields(t *testing.T) {
	writer := ns.NewWriter()
	writer.WriteBool(true)
	writer.WriteVarInt(ns.VarInt(registries.Potion.Get("minecraft:water")))
	writer.WriteBool(true)
	writer.WriteInt32(0x123456)
	writer.WriteVarInt(1)
	writer.WriteVarInt(ns.VarInt(registries.MobEffect.Get("minecraft:poison")))
	writer.WriteVarInt(1)
	writer.WriteVarInt(200)
	writer.WriteBool(false)
	writer.WriteBool(true)
	writer.WriteBool(true)
	writer.WriteBool(false)
	writer.WriteBool(false)
	stack := ns.NewSlot(ns.VarInt(items.ID("minecraft:splash_potion")), 1)
	stack.Components.Add = []ns.RawSlotComponent{{ID: ns.VarInt(items.ComponentPotionContents), Data: writer.Bytes()}}
	contents, err := items.ReadPotionContents(stack)
	if err != nil {
		t.Fatalf("ReadPotionContents: %v", err)
	}
	if contents.Potion != "minecraft:water" || contents.CustomColor == nil || *contents.CustomColor != 0x123456 {
		t.Fatalf("contents = %+v", contents)
	}
	if len(contents.Effects) != 1 || contents.Effects[0].ID != "minecraft:poison" || contents.Effects[0].Amplifier != 1 || contents.Effects[0].Duration != 200 {
		t.Fatalf("custom effects = %+v", contents.Effects)
	}
}

func TestPassthroughComponentsSurviveTypedRoundTrip(t *testing.T) {
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
	slot.Components.Add = []ns.RawSlotComponent{{ID: ns.VarInt(items.ComponentCreativeSlotLock)}}
	slot.Components.Remove = []ns.VarInt{ns.VarInt(items.ComponentCanBreak)}
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	roundTrip, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	if !items.SameStack(slot, roundTrip) {
		t.Fatalf("round trip = %+v, want %+v", roundTrip.Components, slot.Components)
	}
	clone := stack.Components.Clone()
	stack.Components = clone
	clonedSlot, err := stack.ToSlot()
	if err != nil || !items.SameStack(slot, clonedSlot) {
		t.Fatalf("cloned passthrough round trip = %+v, %v", clonedSlot.Components, err)
	}
}

func TestAttributeModifierWireLayout(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:stone"), 1)
	stack.Components.AttributeModifiers = []items.AttributeModifier{{
		Type: "minecraft:attack_damage", Amount: 1.5, ID: "minecraft:test",
		Operation: "add_multiplied_base", Slot: "mainhand",
		Display: items.AttributeModifierDisplay{Type: "hidden"},
	}}
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	var raw []byte
	for _, component := range slot.Components.Add {
		if int32(component.ID) == items.ComponentAttributeModifiers {
			raw = component.Data
			break
		}
	}
	const expected = "01030e6d696e6563726166743a746573743ff8000000000000010101"
	if got := hex.EncodeToString(raw); got != expected {
		t.Fatalf("attribute modifier wire = %s, want %s", got, expected)
	}
	decoded, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	modifier := decoded.Components.AttributeModifiers[0]
	if modifier.Slot != "mainhand" || modifier.Display.Type != "hidden" {
		t.Fatalf("decoded modifier = %+v", modifier)
	}
	defaults := items.DefaultComponents(items.ID("minecraft:carved_pumpkin"))
	if defaults == nil || len(defaults.AttributeModifiers) != 1 || defaults.AttributeModifiers[0].Display.Type != "hidden" {
		t.Fatalf("carved pumpkin modifiers = %+v", defaults)
	}
}

func TestEquipmentSlotGroupWireOrdinals(t *testing.T) {
	for ordinal, slotName := range []string{"any", "mainhand", "offhand", "hand", "feet", "legs", "chest", "head", "armor", "body", "saddle"} {
		stack := items.NewStack(items.ID("minecraft:stone"), 1)
		stack.Components.AttributeModifiers = []items.AttributeModifier{{
			Type: "minecraft:attack_damage", ID: "minecraft:test", Operation: "add_value", Slot: slotName,
		}}
		slot, err := stack.ToSlot()
		if err != nil {
			t.Fatalf("%s ToSlot: %v", slotName, err)
		}
		var reader *ns.PacketBuffer
		for _, component := range slot.Components.Add {
			if int32(component.ID) == items.ComponentAttributeModifiers {
				reader = ns.NewReader(component.Data)
				break
			}
		}
		if reader == nil {
			t.Fatalf("%s attribute modifier component missing", slotName)
		}
		if _, err := reader.ReadVarInt(); err != nil {
			t.Fatal(err)
		}
		if _, err := reader.ReadVarInt(); err != nil {
			t.Fatal(err)
		}
		if _, err := reader.ReadString(32767); err != nil {
			t.Fatal(err)
		}
		if _, err := reader.ReadFloat64(); err != nil {
			t.Fatal(err)
		}
		if _, err := reader.ReadVarInt(); err != nil {
			t.Fatal(err)
		}
		got, err := reader.ReadVarInt()
		if err != nil || int(got) != ordinal {
			t.Fatalf("%s ordinal = %d, %v; want %d", slotName, got, err, ordinal)
		}
	}
}

func TestAttributeModifierOverrideNBTRoundTrip(t *testing.T) {
	override := nbt.Compound{
		"text": nbt.String("Reach"), "color": nbt.String("gold"), "bold": nbt.Byte(1),
		"extra": nbt.List{ElementType: nbt.TagString, Elements: []nbt.Tag{nbt.String(" bonus")}},
	}
	patch := nbt.Compound{"minecraft:attribute_modifiers": nbt.List{
		ElementType: nbt.TagCompound,
		Elements: []nbt.Tag{nbt.Compound{
			"type": nbt.String("minecraft:block_interaction_range"), "id": nbt.String("minecraft:test_reach"),
			"amount": nbt.Double(2.5), "operation": nbt.String("add_value"), "slot": nbt.String("saddle"),
			"display": nbt.Compound{"type": nbt.String("override"), "value": override},
		}},
	}}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	roundTrip, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil {
		t.Fatalf("ReadSlot: %v", err)
	}
	decoded, err := items.FromSlot(roundTrip)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	modifier := decoded.Components.AttributeModifiers[0]
	if modifier.Slot != "saddle" || modifier.Display.Type != "override" || modifier.Display.Value == nil || len(modifier.Display.Value.Raw) == 0 {
		t.Fatalf("decoded override = %+v", modifier)
	}
	display := items.ComponentsToNBT(roundTrip).GetList("minecraft:attribute_modifiers").Elements[0].(nbt.Compound).GetCompound("display")
	value := display.GetCompound("value")
	if display.GetString("type") != "override" || value.GetString("text") != "Reach" || value.GetString("color") != "gold" || value.GetByte("bold") != 1 || value.GetList("extra").Len() != 1 {
		t.Fatalf("override NBT = %+v", display)
	}
}

func TestSameStackIgnoresPatchOrder(t *testing.T) {
	a := ns.NewSlot(1, 1)
	a.Components.Add = []ns.RawSlotComponent{{ID: 1, Data: []byte{1}}, {ID: 2, Data: []byte{2}}}
	a.Components.Remove = []ns.VarInt{3, 4}
	b := ns.NewSlot(1, 1)
	b.Components.Add = []ns.RawSlotComponent{{ID: 2, Data: []byte{2}}, {ID: 1, Data: []byte{1}}}
	b.Components.Remove = []ns.VarInt{4, 3}
	if !items.SameStack(a, b) {
		t.Fatal("equivalent reordered patches did not match")
	}
	b.Components.Add[0].Data[0] = 3
	if items.SameStack(a, b) {
		t.Fatal("different component values matched")
	}

	a.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentEnchantments, Data: []byte{2, 1, 1, 2, 2}}}
	b.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentEnchantments, Data: []byte{2, 2, 2, 1, 1}}}
	a.Components.Remove, b.Components.Remove = nil, nil
	if !items.SameStack(a, b) {
		t.Fatal("semantically equal enchantment maps did not match")
	}

	a.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentCustomData, Data: []byte{10, 3, 0, 1, 'a', 0, 0, 0, 1, 3, 0, 1, 'b', 0, 0, 0, 2, 0}}}
	b.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentCustomData, Data: []byte{10, 3, 0, 1, 'b', 0, 0, 0, 2, 3, 0, 1, 'a', 0, 0, 0, 1, 0}}}
	if !items.SameStack(a, b) {
		t.Fatal("semantically equal custom-data compounds did not match")
	}
}

func TestFireworksComponentsRoundTrip(t *testing.T) {
	rocket := ns.NewSlot(ns.VarInt(items.ID("minecraft:firework_rocket")), 1)
	restored, err := items.ApplyNBTComponents(rocket, nbt.Compound{
		"minecraft:fireworks": nbt.Compound{
			"flight_duration": nbt.Int(3),
			"explosions": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
				nbt.Compound{"shape": nbt.String("star"), "colors": nbt.IntArray{0xff0000}, "has_trail": nbt.Byte(1)},
			}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := items.FromSlot(restored)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Components.Fireworks == nil || decoded.Components.Fireworks.FlightDuration != 3 || len(decoded.Components.FireworkExplosions) != 1 || decoded.Components.FireworkExplosions[0].Shape != 2 {
		t.Fatalf("fireworks = %+v/%+v", decoded.Components.Fireworks, decoded.Components.FireworkExplosions)
	}
	if got := items.ComponentsToNBT(restored).GetCompound("minecraft:fireworks").GetInt("flight_duration"); got != 3 {
		t.Fatalf("firework flight NBT = %d", got)
	}
}

func TestPotDecorationsComponentRoundTrip(t *testing.T) {
	stack := ns.NewSlot(ns.VarInt(items.ID("minecraft:decorated_pot")), 1)
	restored, err := items.ApplyNBTComponents(stack, nbt.Compound{
		"minecraft:pot_decorations": nbt.Compound{
			"back": nbt.Compound{
				"id":    nbt.String("minecraft:angler_pottery_sherd"),
				"count": nbt.Int(2),
				"components": nbt.Compound{
					"minecraft:custom_name":  nbt.Compound{"text": nbt.String("etched"), "color": nbt.String("gold")},
					"!minecraft:break_sound": nbt.Compound{},
				},
			},
			"left": nbt.String("minecraft:brick"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	writer := ns.NewWriter()
	if err := writer.WriteSlot(restored); err != nil {
		t.Fatal(err)
	}
	writer.WriteVarInt(123456)
	reader := ns.NewReader(writer.Bytes())
	wire, err := reader.ReadSlot(items.Decoder())
	if err != nil {
		t.Fatal(err)
	}
	if sentinel, err := reader.ReadVarInt(); err != nil || sentinel != 123456 {
		t.Fatalf("pot payload alignment: %d, %v", sentinel, err)
	}
	decoded, err := items.FromSlot(wire)
	if err != nil {
		t.Fatal(err)
	}
	decorations := decoded.Components.PotDecorations
	if decorations == nil || decorations.Back == nil || decorations.Back.ID != items.ID("minecraft:angler_pottery_sherd") ||
		decorations.Back.Count != 2 || decorations.Left == nil || decorations.Right != nil || decorations.Front != nil {
		t.Fatalf("pot decorations = %+v", decorations)
	}
	clone := decoded.Components.Clone()
	clone.PotDecorations.Back.Count = 3
	clone.PotDecorations.Back.Components.CustomName.Raw[0] = 0
	if decorations.Back.Count != 2 || decorations.Back.Components.CustomName.Raw[0] == 0 {
		t.Fatal("decoration clone shares nested component state")
	}
	patch := items.ComponentsToNBT(wire).GetCompound("minecraft:pot_decorations")
	back := patch.GetCompound("back")
	if back.GetInt("count") != 2 || back.GetCompound("components").GetCompound("minecraft:custom_name").GetString("color") != "gold" {
		t.Fatalf("nested decoration patch lost: %+v", patch)
	}
	if _, present := back.GetCompound("components")["!minecraft:break_sound"]; !present {
		t.Fatal("nested component removal lost")
	}
	if _, present := patch["right"]; present {
		t.Fatal("absent right side coerced to an item")
	}
	empty, err := items.ApplyNBTComponents(stack, nbt.Compound{"minecraft:pot_decorations": nbt.Compound{}})
	if err != nil {
		t.Fatal(err)
	}
	if value, present := items.ComponentsToNBT(empty)["minecraft:pot_decorations"]; !present || len(value.(nbt.Compound)) != 0 {
		t.Fatalf("explicit empty decorations not preserved: %+v", value)
	}
}

func TestResolvableFuelWireConstantsAndReferences(t *testing.T) {
	writer := ns.NewWriter()
	writer.WriteBool(true)
	writer.WriteInt32(0x01020304)
	writer.WriteBool(false)
	writer.WriteString("minecraft:cooking/speed_default")
	writer.WriteVarInt(123456)
	reader := ns.NewReader(writer.Bytes())
	data, err := items.Decoder()(reader, items.ComponentCookingFuel)
	if err != nil {
		t.Fatal(err)
	}
	if sentinel, err := reader.ReadVarInt(); err != nil || sentinel != 123456 {
		t.Fatalf("fuel payload alignment: %d, %v", sentinel, err)
	}
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
	slot.AddComponent(items.ComponentCookingFuel, data)
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatal(err)
	}
	fuel := stack.Components.CookingFuel
	if fuel == nil || fuel.BurnTime.Value != 0x01020304 || fuel.SpeedMultiplier.Reference != "minecraft:cooking/speed_default" {
		t.Fatalf("resolved fuel = %+v", fuel)
	}
	patch := items.ComponentsToNBT(slot)
	roundTrip, err := items.ApplyNBTComponents(ns.NewSlot(slot.ItemID, 1), patch)
	if err != nil || !items.SameStack(slot, roundTrip) {
		t.Fatalf("fuel persistence: %+v, %v", roundTrip, err)
	}
	other, err := items.ApplyNBTComponents(ns.NewSlot(slot.ItemID, 1), nbt.Compound{
		"minecraft:brewing_fuel": nbt.Compound{"uses": nbt.String("minecraft:brewing/fuel_blaze_powder"), "speed_multiplier": nbt.Float(1.25)},
		"minecraft:compostable":  nbt.Compound{"layers": nbt.Int(-3)},
	})
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := items.FromSlot(other)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Components.BrewingFuel.Uses.Reference != "minecraft:brewing/fuel_blaze_powder" ||
		decoded.Components.BrewingFuel.SpeedMultiplier.Value != 1.25 || decoded.Components.Compostable.Layers.Value != -3 {
		t.Fatalf("number union variants lost: %+v", decoded.Components)
	}
}

func TestSignTextFixedFourLinesAndFilteredPersistence(t *testing.T) {
	writer := ns.NewWriter()
	for i := range 4 {
		if err := writer.WriteTextComponent(ns.TextComponent{Text: "line", Color: "gold", Extra: []ns.TextComponent{{Text: "child"}}}); err != nil {
			t.Fatal(i, err)
		}
	}
	writer.WriteBool(true)
	for range 4 {
		writer.WriteTextComponent(ns.TextComponent{Text: "filtered"})
	}
	writer.WriteVarInt(14)
	writer.WriteBool(true)
	writer.WriteVarInt(123456)
	reader := ns.NewReader(writer.Bytes())
	data, err := items.Decoder()(reader, items.ComponentSignTextFront)
	if err != nil {
		t.Fatal(err)
	}
	if sentinel, err := reader.ReadVarInt(); err != nil || sentinel != 123456 {
		t.Fatalf("sign payload alignment: %d, %v", sentinel, err)
	}
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
	slot.AddComponent(items.ComponentSignTextFront, data)
	patch := items.ComponentsToNBT(slot)
	sign := patch.GetCompound("minecraft:sign_text_front")
	if sign.GetString("color") != "red" || len(sign.GetList("filtered_messages").Elements) != 4 ||
		sign.GetList("messages").Elements[0].(nbt.Compound).GetString("color") != "gold" {
		t.Fatalf("styled or filtered sign text lost: %+v", sign)
	}
	roundTrip, err := items.ApplyNBTComponents(ns.NewSlot(slot.ItemID, 1), patch)
	if err != nil || !items.SameStack(slot, roundTrip) {
		t.Fatalf("sign persistence: %+v, %v", roundTrip, err)
	}
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatal(err)
	}
	clone := stack.Components.Clone()
	clone.SignTextFront.FilteredMessages[0].Raw[0] = 0
	if stack.Components.SignTextFront.FilteredMessages[0].Raw[0] == 0 {
		t.Fatal("filtered sign text clone aliases original")
	}
	sign["messages"] = nbt.List{ElementType: nbt.TagString, Elements: []nbt.Tag{nbt.String("one")}}
	if _, err := items.ApplyNBTComponents(ns.NewSlot(slot.ItemID, 1), patch); err == nil {
		t.Fatal("accepted a sign without four messages")
	}
}

func TestPlainRegistryComponentsUseZeroBasedHolders(t *testing.T) {
	for _, test := range []struct {
		component      int32
		name, registry string
	}{
		{items.ComponentBlockTransformer, "minecraft:block_transformer", "minecraft:block_transformer"},
		{items.ComponentProvidesPotteryPattern, "minecraft:provides_pottery_pattern", "minecraft:decorated_pot_pattern"},
	} {
		t.Run(test.name, func(t *testing.T) {
			reader := ns.NewReader([]byte{0, 42})
			data, err := items.Decoder()(reader, ns.VarInt(test.component))
			if err != nil {
				t.Fatal(err)
			}
			if next, err := reader.ReadVarInt(); err != nil || next != 42 {
				t.Fatalf("plain holder consumed inline data: %d, %v", next, err)
			}
			slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
			slot.AddComponent(ns.VarInt(test.component), data)
			patch := items.ComponentsToNBT(slot)
			if patch.GetString(test.name) != registries.SynchronizedEntries[test.registry].ByID(0) {
				t.Fatalf("holder zero mapped incorrectly: %+v", patch)
			}
			restored, err := items.ApplyNBTComponents(ns.NewSlot(slot.ItemID, 1), patch)
			if err != nil || !items.SameStack(slot, restored) {
				t.Fatalf("holder persistence: %+v, %v", restored, err)
			}
		})
	}
}

func TestBannerPatternsComponentRoundTrip(t *testing.T) {
	stack := ns.NewSlot(ns.VarInt(items.ID("minecraft:red_banner")), 1)
	restored, err := items.ApplyNBTComponents(stack, nbt.Compound{
		"minecraft:base_color": nbt.String("red"),
		"minecraft:banner_patterns": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"pattern": nbt.String("minecraft:stripe_top"), "color": nbt.String("white")},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	var component ns.RawSlotComponent
	for _, candidate := range restored.Components.Add {
		if int32(candidate.ID) == items.ComponentBannerPatterns {
			component = candidate
			break
		}
	}
	wire := ns.NewReader(component.Data)
	count, err := wire.ReadVarInt()
	if err != nil || count != 1 {
		t.Fatalf("banner pattern count = %d, %v", count, err)
	}
	holder, err := wire.ReadVarInt()
	if err != nil || registries.SynchronizedEntries["minecraft:banner_pattern"].HolderName(int32(holder)) != "minecraft:stripe_top" {
		t.Fatalf("banner pattern holder = %d, %v", holder, err)
	}
	decoded, err := items.FromSlot(restored)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Components.BannerPatterns == nil || len(*decoded.Components.BannerPatterns) != 1 || (*decoded.Components.BannerPatterns)[0].Pattern != "minecraft:stripe_top" {
		t.Fatalf("banner patterns = %+v", decoded.Components.BannerPatterns)
	}
	if decoded.Components.BaseColor == nil || *decoded.Components.BaseColor != 14 {
		t.Fatalf("base color = %v", decoded.Components.BaseColor)
	}
	patch := items.ComponentsToNBT(restored)
	if len(patch.GetList("minecraft:banner_patterns").Elements) != 1 {
		t.Fatalf("banner pattern nbt = %+v", patch)
	}
}

func TestDyeComponentsRoundTrip(t *testing.T) {
	for id, name := range []string{
		"white", "orange", "magenta", "light_blue", "yellow", "lime", "pink", "gray",
		"light_gray", "cyan", "purple", "blue", "brown", "green", "red", "black",
	} {
		defaults := items.DefaultComponents(items.ID("minecraft:" + name + "_dye"))
		if defaults == nil || defaults.Dye == nil || *defaults.Dye != int32(id) {
			t.Fatalf("%s dye default = %+v, want %d", name, defaults, id)
		}
	}
	stack := ns.NewSlot(ns.VarInt(items.ID("minecraft:leather_chestplate")), 1)
	restored, err := items.ApplyNBTComponents(stack, nbt.Compound{
		"minecraft:dye":        nbt.String("white"),
		"minecraft:dyed_color": nbt.Int(0x102030),
	})
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := items.FromSlot(restored)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Components.Dye == nil || *decoded.Components.Dye != 0 || decoded.Components.DyedColor == nil || *decoded.Components.DyedColor != 0x102030 {
		t.Fatalf("dye components = %v/%v", decoded.Components.Dye, decoded.Components.DyedColor)
	}
	patch := items.ComponentsToNBT(restored)
	if patch.GetString("minecraft:dye") != "white" || patch.GetInt("minecraft:dyed_color") != 0x102030 {
		t.Fatalf("dye nbt = %+v", patch)
	}
}

func TestPaintingVariantComponentRoundTrip(t *testing.T) {
	stack := ns.NewSlot(ns.VarInt(items.ID("minecraft:painting")), 1)
	restored, err := items.ApplyNBTComponents(stack, nbt.Compound{
		"minecraft:painting/variant": nbt.String("minecraft:kebab"),
	})
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := items.FromSlot(restored)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Components.PaintingVariant != "minecraft:kebab" {
		t.Fatalf("painting variant = %q", decoded.Components.PaintingVariant)
	}
	if got := items.ComponentsToNBT(restored).GetString("minecraft:painting/variant"); got != "minecraft:kebab" {
		t.Fatalf("painting variant nbt = %q", got)
	}
}

func TestItemStackTemplateRoundTrip(t *testing.T) {
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:diamond_sword")), 1)
	slot.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentDamage, Data: []byte{7}}}
	slot.Components.Remove = []ns.VarInt{items.ComponentCustomName}
	writer := ns.NewWriter()
	if err := items.WriteItemStackTemplate(writer, slot); err != nil {
		t.Fatal(err)
	}
	decoded, err := items.ReadItemStackTemplate(ns.NewReader(writer.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Count != slot.Count || !items.SameStack(decoded, slot) {
		t.Fatalf("decoded template = %+v, want %+v", decoded, slot)
	}
}

func TestCraftingRemainders(t *testing.T) {
	for ingredient, remainder := range map[string]string{
		"minecraft:water_bucket":  "minecraft:bucket",
		"minecraft:lava_bucket":   "minecraft:bucket",
		"minecraft:milk_bucket":   "minecraft:bucket",
		"minecraft:dragon_breath": "minecraft:glass_bottle",
		"minecraft:honey_bottle":  "minecraft:glass_bottle",
	} {
		got, ok := items.CraftingRemainder(items.ID(ingredient))
		if !ok || got != items.ID(remainder) {
			t.Errorf("CraftingRemainder(%s) = %s, %v", ingredient, items.Name(got), ok)
		}
	}
	if _, ok := items.CraftingRemainder(items.ID("minecraft:stick")); ok {
		t.Fatal("stick has a crafting remainder")
	}
}

func TestItemIDLookup(t *testing.T) {
	tests := []string{
		"minecraft:diamond_sword",
		"minecraft:iron_sword",
		"minecraft:apple",
		"minecraft:golden_apple",
		"minecraft:diamond_pickaxe",
		"minecraft:iron_pickaxe",
		"minecraft:stick",
		"minecraft:diamond",
	}

	for _, name := range tests {
		t.Run(name, func(t *testing.T) {
			id := items.ID(name)
			if id < 0 {
				t.Fatalf("ItemID(%q) = %d, want >= 0", name, id)
			}
			if got := items.Name(id); got != name {
				t.Errorf("ItemName(%d) = %q, want %q", id, got, name)
			}
		})
	}
}

func TestItemIDNotFound(t *testing.T) {
	if got := items.ID("minecraft:nonexistent_item"); got != -1 {
		t.Errorf("ItemID for nonexistent item = %d, want -1", got)
	}
	if got := items.Name(-999); got != "" {
		t.Errorf("ItemName for invalid ID = %q, want empty string", got)
	}
}

func TestConsumableComponentRoundTrip(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:golden_apple"), 1)
	stack.Components.Consumable.ConsumeSeconds = 0.5
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	if len(slot.Components.Add) != 1 || int32(slot.Components.Add[0].ID) != items.ComponentConsumable {
		t.Fatalf("component patch = %+v", slot.Components)
	}
	decoded, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	if decoded.Components.Consumable == nil || decoded.Components.Consumable.ConsumeSeconds != 0.5 {
		t.Fatalf("decoded consumable = %+v", decoded.Components.Consumable)
	}
	if effects := decoded.Components.Consumable.OnConsumeEffects; len(effects) != 1 || len(effects[0].Effects) != 2 || effects[0].Effects[0].ID != "minecraft:regeneration" {
		t.Fatalf("decoded consume effects = %+v", effects)
	}

	removed := ns.NewSlot(ns.VarInt(items.ID("minecraft:golden_apple")), 1)
	removed.Components.Remove = []ns.VarInt{ns.VarInt(items.ComponentConsumable)}
	decoded, err = items.FromSlot(removed)
	if err != nil {
		t.Fatalf("FromSlot removed consumable: %v", err)
	}
	if decoded.Components.Consumable != nil {
		t.Fatalf("removed consumable = %+v", decoded.Components.Consumable)
	}
}

func TestComponentNBTRoundTrip(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:diamond_chestplate"), 1)
	stack.Components.Damage = 17
	stack.Components.Trim = &items.ArmorTrim{Material: "minecraft:gold", Pattern: "minecraft:sentry"}
	stack.Components.Enchantments = map[string]int32{"minecraft:protection": 4}
	stack.Components.CustomName = &items.ItemNameComponent{Text: "Test Armor"}
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	components := items.ComponentsToNBT(slot)
	if components.GetInt("minecraft:damage") != 17 || components.GetCompound("minecraft:trim").GetString("material") != "minecraft:gold" {
		t.Fatalf("component NBT = %+v", components)
	}
	if name, ok := components["minecraft:custom_name"].(nbt.String); !ok || string(name) != "Test Armor" {
		t.Fatalf("custom name NBT = %+v", components["minecraft:custom_name"])
	}
	restored, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:diamond_chestplate")), 1), components)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	decoded, err := items.FromSlot(restored)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	if decoded.Components.Damage != 17 || decoded.Components.Trim == nil || decoded.Components.Trim.Pattern != "minecraft:sentry" || decoded.Components.Enchantments["minecraft:protection"] != 4 || decoded.Components.CustomName.Text != "Test Armor" {
		t.Fatalf("restored components = %+v", decoded.Components)
	}

	potion := items.PotionSlot(ns.VarInt(items.ID("minecraft:potion")), "minecraft:strong_swiftness")
	potionNBT := items.ComponentsToNBT(potion)
	restoredPotion, err := items.ApplyNBTComponents(ns.NewSlot(potion.ItemID, 1), potionNBT)
	if err != nil || items.PotionName(restoredPotion) != "minecraft:strong_swiftness" {
		t.Fatalf("restored potion = %+v, %v", restoredPotion, err)
	}
}

func TestWrittenBookContentNBTRoundTrip(t *testing.T) {
	patch := nbt.Compound{"minecraft:written_book_content": nbt.Compound{
		"title":  nbt.Compound{"raw": nbt.String("Title"), "filtered": nbt.String("Filtered")},
		"author": nbt.String("Author"), "generation": nbt.Int(2), "resolved": nbt.Byte(1),
		"pages": nbt.List{ElementType: nbt.TagString, Elements: []nbt.Tag{nbt.String("Page one"), nbt.String("Page two")}},
	}}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:written_book")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil || !items.SameStack(slot, decoded) {
		t.Fatalf("written book wire round trip = %+v, %v", decoded.Components, err)
	}
	book := items.ComponentsToNBT(decoded).GetCompound("minecraft:written_book_content")
	if book.GetString("author") != "Author" || book.GetInt("generation") != 2 || book.GetByte("resolved") != 1 || book.GetList("pages").Len() != 2 || book.GetCompound("title").GetString("filtered") != "Filtered" {
		t.Fatalf("written book = %+v", book)
	}
}

func TestWritableBookHasDefaultEditableContent(t *testing.T) {
	writable := ns.NewSlot(ns.VarInt(items.ID("minecraft:writable_book")), 1)
	written := ns.NewSlot(ns.VarInt(items.ID("minecraft:written_book")), 1)
	if !items.HasWritableBookContent(writable) || items.HasWritableBookContent(written) {
		t.Fatal("generated writable-book content default is incorrect")
	}
	removed, err := items.ApplyNBTComponents(writable, nbt.Compound{"!minecraft:writable_book_content": nbt.Compound{}})
	if err != nil || items.HasWritableBookContent(removed) {
		t.Fatalf("removed writable content remained editable: %v", err)
	}
}

func TestWritableBookContentNBTRoundTrip(t *testing.T) {
	patch := nbt.Compound{"minecraft:writable_book_content": nbt.Compound{
		"pages": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"raw": nbt.String("hello"), "filtered": nbt.String("filtered")},
			nbt.Compound{"raw": nbt.String("world")},
		}},
	}}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:writable_book")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	pages := items.ComponentsToNBT(slot).GetCompound("minecraft:writable_book_content").GetList("pages")
	if pages.Len() != 2 || pages.Elements[0].(nbt.Compound).GetString("filtered") != "filtered" || pages.Elements[1].(nbt.Compound).GetString("raw") != "world" {
		t.Fatalf("writable pages = %+v", pages)
	}
}

func TestItemContainerComponentNBTRoundTrip(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:sulfur_cube_content": nbt.Compound{"id": nbt.String("minecraft:stone"), "count": nbt.Int(4)},
		"minecraft:bundle_contents": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"id": nbt.String("minecraft:apple"), "count": nbt.Int(2)},
		}},
		"minecraft:container": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"slot": nbt.Int(2), "item": nbt.Compound{"id": nbt.String("minecraft:stone"), "count": nbt.Int(3)}},
		}},
	}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:bundle")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil || !items.SameStack(slot, decoded) {
		t.Fatalf("template-list wire round trip = %+v, %v", decoded.Components, err)
	}
	roundTrip := items.ComponentsToNBT(decoded)
	bundle := roundTrip.GetList("minecraft:bundle_contents")
	container := roundTrip.GetList("minecraft:container")
	if sulfur := roundTrip.GetCompound("minecraft:sulfur_cube_content"); sulfur.GetString("id") != "minecraft:stone" || sulfur.GetInt("count") != 4 {
		t.Fatalf("sulfur cube content = %+v", sulfur)
	}
	if bundle.Len() != 1 || bundle.Elements[0].(nbt.Compound).GetString("id") != "minecraft:apple" {
		t.Fatalf("bundle contents = %+v", bundle)
	}
	if container.Len() != 1 || container.Elements[0].(nbt.Compound).GetInt("slot") != 2 || container.Elements[0].(nbt.Compound).GetCompound("item").GetString("id") != "minecraft:stone" {
		t.Fatalf("container contents = %+v", container)
	}
}

func TestCustomPotionContentsNBTRoundTrip(t *testing.T) {
	hidden := nbt.Compound{"duration": nbt.Int(40), "amplifier": nbt.Int(0)}
	patch := nbt.Compound{"minecraft:potion_contents": nbt.Compound{
		"potion": nbt.String("minecraft:awkward"), "custom_color": nbt.Int(0x123456), "custom_name": nbt.String("test"),
		"custom_effects": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{nbt.Compound{
			"id": nbt.String("minecraft:poison"), "duration": nbt.Int(200), "amplifier": nbt.Int(1), "hidden_effect": hidden,
		}}},
	}}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:splash_potion")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	contents, err := items.ReadPotionContents(slot)
	if err != nil || contents.CustomColor == nil || *contents.CustomColor != 0x123456 || contents.CustomName != "test" || len(contents.Effects) != 1 || contents.Effects[0].Hidden == nil {
		t.Fatalf("potion contents = %+v, %v", contents, err)
	}
	roundTrip := items.ComponentsToNBT(slot).GetCompound("minecraft:potion_contents")
	if roundTrip.GetString("potion") != "minecraft:awkward" || roundTrip.GetString("custom_name") != "test" || roundTrip.GetList("custom_effects").Len() != 1 {
		t.Fatalf("potion NBT = %+v", roundTrip)
	}
}

func TestStyledTextComponentsRoundTrip(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:custom_name": nbt.Compound{"text": nbt.String("Styled"), "color": nbt.String("red"), "bold": nbt.Byte(1)},
		"minecraft:lore": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"text": nbt.String("Line"), "italic": nbt.Byte(0), "color": nbt.String("gold")},
		}},
	}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	roundTripSlot, err := stack.ToSlot()
	if err != nil || !items.SameStack(slot, roundTripSlot) {
		t.Fatalf("styled wire round trip = %+v, %v", roundTripSlot.Components, err)
	}
	roundTrip := items.ComponentsToNBT(roundTripSlot)
	if roundTrip.GetCompound("minecraft:custom_name").GetString("color") != "red" ||
		roundTrip.GetList("minecraft:lore").Elements[0].(nbt.Compound).GetString("color") != "gold" {
		t.Fatalf("styled components = %+v", roundTrip)
	}
}

func TestTypedEntityDataComponentRoundTrip(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:entity_data":        nbt.Compound{"id": nbt.String("minecraft:pig"), "CustomNameVisible": nbt.Byte(1)},
		"minecraft:block_entity_data":  nbt.Compound{"id": nbt.String("minecraft:chest"), "CustomName": nbt.String("Chest")},
		"minecraft:bucket_entity_data": nbt.Compound{"Health": nbt.Float(4)},
	}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil || !items.SameStack(slot, decoded) {
		t.Fatalf("typed entity data wire round trip = %+v, %v", decoded.Components, err)
	}
	stack, err := items.FromSlot(decoded)
	if err != nil || stack.Components.EntityData == nil || stack.Components.EntityData.Type != "minecraft:pig" {
		t.Fatalf("decoded entity data = %+v, %v", stack.Components.EntityData, err)
	}
	pigEgg := items.DefaultComponents(items.ID("minecraft:pig_spawn_egg"))
	if pigEgg == nil || pigEgg.EntityData == nil || pigEgg.EntityData.Type != "minecraft:pig" {
		t.Fatalf("pig spawn egg entity data = %+v", pigEgg)
	}
	roundTrip := items.ComponentsToNBT(decoded)
	if roundTrip.GetCompound("minecraft:entity_data").GetString("id") != "minecraft:pig" ||
		roundTrip.GetCompound("minecraft:block_entity_data").GetString("id") != "minecraft:chest" ||
		roundTrip.GetCompound("minecraft:bucket_entity_data").GetFloat("Health") != 4 {
		t.Fatalf("typed entity data = %+v", roundTrip)
	}
}

func TestConsumableDirectSoundRangesRoundTrip(t *testing.T) {
	writer := ns.NewWriter()
	writer.WriteFloat32(1.6)
	writer.WriteVarInt(1)
	writer.WriteVarInt(0)
	writer.WriteString("minecraft:test.consume")
	writer.WriteBool(true)
	writer.WriteFloat32(12.5)
	writer.WriteBool(true)
	writer.WriteVarInt(1)
	writer.WriteVarInt(4)
	writer.WriteVarInt(0)
	writer.WriteString("minecraft:test.effect")
	writer.WriteBool(true)
	writer.WriteFloat32(8)
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
	slot.Components.Add = []ns.RawSlotComponent{{ID: ns.VarInt(items.ComponentConsumable), Data: writer.Bytes()}}
	stack, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	roundTrip, err := stack.ToSlot()
	if err != nil || !items.SameStack(slot, roundTrip) {
		t.Fatalf("consumable sound round trip = %+v, %v", roundTrip.Components, err)
	}
}

func TestEquippableComponentWireAndNBTRoundTrip(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:iron_helmet"), 1)
	equipRange, shearRange := 12.5, 7.5
	stack.Components.Equippable = &items.Equippable{
		Slot:       "head",
		EquipSound: "minecraft:test.equip", EquipSoundDirect: true, EquipSoundRange: &equipRange,
		AssetID: "minecraft:test", CameraOverlay: "minecraft:misc/test",
		AllowedEntities: []string{"#minecraft:undead"},
		Dispensable:     false, Swappable: false, DamageOnHurt: false, EquipOnInteract: true, CanBeSheared: true,
		ShearingSound: "minecraft:test.shear", ShearingSoundDirect: true, ShearingSoundRange: &shearRange,
	}
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil || !items.SameStack(slot, decoded) {
		t.Fatalf("equippable wire round trip = %+v, %v", decoded.Components, err)
	}
	component, err := items.FromSlot(decoded)
	if err != nil || component.Components.Equippable == nil {
		t.Fatalf("decoded equippable = %+v, %v", component, err)
	}
	equippable := component.Components.Equippable
	if equippable.Dispensable || !equippable.EquipOnInteract || !equippable.CanBeSheared ||
		equippable.EquipSoundRange == nil || *equippable.EquipSoundRange != equipRange ||
		equippable.ShearingSoundRange == nil || *equippable.ShearingSoundRange != shearRange {
		t.Fatalf("decoded equippable = %+v", equippable)
	}
	patch := items.ComponentsToNBT(decoded)
	restored, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:iron_helmet")), 1), patch)
	if err != nil || !items.SameStack(decoded, restored) {
		t.Fatalf("equippable NBT round trip = %+v, %v", patch, err)
	}
}

func TestDirectHolderComponentNBTRoundTrip(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:break_sound": nbt.Compound{"sound_id": nbt.String("minecraft:test.break"), "range": nbt.Float(8)},
		"minecraft:instrument": nbt.Compound{
			"sound_event": nbt.String("minecraft:item.goat_horn.sound.0"), "use_duration": nbt.Float(7),
			"range": nbt.Float(32), "description": nbt.String("Horn"),
		},
		"minecraft:jukebox_playable": nbt.Compound{
			"sound_event": nbt.String("minecraft:music_disc.13"), "description": nbt.String("Song"),
			"length_in_seconds": nbt.Float(12), "comparator_output": nbt.Int(4),
		},
		"minecraft:provides_trim_material": nbt.Compound{
			"asset_name": nbt.String("test"), "override_armor_assets": nbt.Compound{"minecraft:iron": nbt.String("test_iron")},
			"description": nbt.String("Material"),
		},
	}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil || !items.SameStack(slot, decoded) {
		t.Fatalf("direct holder wire round trip = %+v, %v", decoded.Components, err)
	}
	roundTrip := items.ComponentsToNBT(decoded)
	if roundTrip.GetCompound("minecraft:break_sound").GetString("sound_id") != "minecraft:test.break" ||
		roundTrip.GetCompound("minecraft:instrument").GetFloat("use_duration") != 7 ||
		roundTrip.GetCompound("minecraft:jukebox_playable").GetInt("comparator_output") != 4 ||
		roundTrip.GetCompound("minecraft:provides_trim_material").GetString("asset_name") != "test" {
		t.Fatalf("direct holders = %+v", roundTrip)
	}
}

func TestHolderComponentReferencesUseVanillaHolderEncoding(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:break_sound":            nbt.String("minecraft:block.note_block.harp"),
		"minecraft:instrument":             nbt.String("minecraft:ponder_goat_horn"),
		"minecraft:jukebox_playable":       nbt.String("minecraft:13"),
		"minecraft:provides_trim_material": nbt.String("minecraft:quartz"),
	}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	registriesByComponent := map[int32]struct {
		registry string
		name     string
	}{
		items.ComponentInstrument:           {"minecraft:instrument", "minecraft:ponder_goat_horn"},
		items.ComponentJukeboxPlayable:      {"minecraft:jukebox_song", "minecraft:13"},
		items.ComponentProvidesTrimMaterial: {"minecraft:trim_material", "minecraft:quartz"},
	}
	for _, component := range slot.Components.Add {
		id, err := ns.NewReader(component.Data).ReadVarInt()
		if err != nil {
			t.Fatalf("component %d holder: %v", component.ID, err)
		}
		if int32(component.ID) == items.ComponentBreakSound {
			if want := registries.SoundEvent.HolderID("minecraft:block.note_block.harp"); int32(id) != want {
				t.Fatalf("break sound holder = %d, want %d", id, want)
			}
		} else if entry, ok := registriesByComponent[int32(component.ID)]; ok {
			if want := registries.SynchronizedEntries[entry.registry].HolderID(entry.name); int32(id) != want {
				t.Fatalf("component %d holder = %d, want %d", component.ID, id, want)
			}
		}
	}
	if _, err := items.FromSlot(slot); err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	roundTrip := items.ComponentsToNBT(slot)
	for name, value := range patch {
		if roundTrip.GetString(name) != string(value.(nbt.String)) {
			t.Fatalf("%s round trip = %+v", name, roundTrip[name])
		}
	}
}

func TestAquaticVariantComponentNBTRoundTrip(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:axolotl/variant":             nbt.String("blue"),
		"minecraft:tropical_fish/pattern":       nbt.String("clayfish"),
		"minecraft:tropical_fish/base_color":    nbt.String("cyan"),
		"minecraft:tropical_fish/pattern_color": nbt.String("magenta"),
	}
	slot, err := items.ApplyNBTComponents(
		ns.NewSlot(ns.VarInt(items.ID("minecraft:tropical_fish_bucket")), 1), patch,
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := items.FromSlot(slot); err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	roundTrip := items.ComponentsToNBT(slot)
	for name, value := range patch {
		if got := roundTrip.GetString(name); got != string(value.(nbt.String)) {
			t.Fatalf("%s = %q, want %q", name, got, value)
		}
	}
}

func TestBasicPositiveComponentNBTRoundTrip(t *testing.T) {
	patch := nbt.Compound{
		"minecraft:consumable": nbt.Compound{
			"consume_seconds": nbt.Float(0.8), "animation": nbt.String("drink"),
			"sound":                 nbt.Compound{"sound_id": nbt.String("minecraft:test.consume"), "range": nbt.Float(12)},
			"has_consume_particles": nbt.Byte(0),
			"on_consume_effects": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
				nbt.Compound{"type": nbt.String("minecraft:apply_effects"), "effects": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
					nbt.Compound{"id": nbt.String("minecraft:poison"), "duration": nbt.Int(100), "show_particles": nbt.Byte(1)},
				}}},
				nbt.Compound{"type": nbt.String("minecraft:play_sound"), "sound": nbt.String("minecraft:entity.generic.drink")},
			}},
		},
		"minecraft:food":            nbt.Compound{"nutrition": nbt.Int(4), "saturation": nbt.Float(2.4), "can_always_eat": nbt.Byte(1)},
		"minecraft:weapon":          nbt.Compound{"item_damage_per_attack": nbt.Int(2), "disable_blocking_for_seconds": nbt.Float(5)},
		"minecraft:enchantable":     nbt.Compound{"value": nbt.Int(15)},
		"minecraft:tooltip_display": nbt.Compound{"hide_tooltip": nbt.Byte(1), "hidden_components": nbt.List{ElementType: nbt.TagString, Elements: []nbt.Tag{nbt.String("minecraft:damage")}}},
		"minecraft:use_cooldown":    nbt.Compound{"seconds": nbt.Float(1.5), "cooldown_group": nbt.String("minecraft:test")},
		"minecraft:use_effects":     nbt.Compound{"can_sprint": nbt.Byte(1), "interact_vibrations": nbt.Byte(0), "speed_multiplier": nbt.Float(0.4)},
		"minecraft:recipes":         nbt.List{ElementType: nbt.TagString, Elements: []nbt.Tag{nbt.String("minecraft:stick")}},
		"minecraft:block_state":     nbt.Compound{"facing": nbt.String("north"), "powered": nbt.String("true")},
		"minecraft:max_stack_size":  nbt.Int(16), "minecraft:potion_duration_scale": nbt.Float(0.5),
		"minecraft:item_model": nbt.String("minecraft:apple"), "minecraft:enchantment_glint_override": nbt.Byte(1),
		"minecraft:tooltip_style": nbt.String("minecraft:test"), "minecraft:creative_slot_lock": nbt.Compound{},
		"minecraft:intangible_projectile": nbt.Compound{},
	}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil || !items.SameStack(slot, decoded) {
		t.Fatalf("basic component wire round trip = %+v, %v", decoded.Components, err)
	}
	roundTrip := items.ComponentsToNBT(decoded)
	if roundTrip.GetInt("minecraft:max_stack_size") != 16 || roundTrip.GetFloat("minecraft:potion_duration_scale") != 0.5 ||
		roundTrip.GetString("minecraft:item_model") != "minecraft:apple" || roundTrip.GetByte("minecraft:enchantment_glint_override") != 1 ||
		roundTrip.GetString("minecraft:tooltip_style") != "minecraft:test" ||
		roundTrip.GetCompound("minecraft:food").GetInt("nutrition") != 4 || roundTrip.GetCompound("minecraft:weapon").GetInt("item_damage_per_attack") != 2 ||
		roundTrip.GetCompound("minecraft:enchantable").GetInt("value") != 15 || roundTrip.GetCompound("minecraft:tooltip_display").GetByte("hide_tooltip") != 1 ||
		roundTrip.GetCompound("minecraft:use_cooldown").GetString("cooldown_group") != "minecraft:test" ||
		roundTrip.GetCompound("minecraft:use_effects").GetFloat("speed_multiplier") != float32(0.4) ||
		roundTrip.GetCompound("minecraft:consumable").GetString("animation") != "drink" ||
		roundTrip.GetCompound("minecraft:consumable").GetList("on_consume_effects").Len() != 2 ||
		roundTrip.GetList("minecraft:recipes").Len() != 1 || roundTrip.GetCompound("minecraft:block_state").GetString("facing") != "north" {
		t.Fatalf("basic component round trip = %+v", roundTrip)
	}
	if _, ok := roundTrip["minecraft:creative_slot_lock"].(nbt.Compound); !ok {
		t.Fatalf("creative slot lock = %+v", roundTrip["minecraft:creative_slot_lock"])
	}
}

func TestTrimComponentRoundTrip(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:diamond_chestplate"), 1)
	stack.Components.Trim = &items.ArmorTrim{Material: "minecraft:gold", Pattern: "minecraft:sentry"}
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	if len(slot.Components.Add) != 1 || int32(slot.Components.Add[0].ID) != items.ComponentTrim || hex.EncodeToString(slot.Components.Add[0].Data) != "0509" {
		t.Fatalf("trim patch = %+v", slot.Components)
	}
	decoded, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	if decoded.Components.Trim == nil || decoded.Components.Trim.Material != "minecraft:gold" || decoded.Components.Trim.Pattern != "minecraft:sentry" {
		t.Fatalf("decoded trim = %+v", decoded.Components.Trim)
	}
}

func TestUseCooldownGroupRoundTrip(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:chorus_fruit"), 1)
	stack.Components.UseCooldown.CooldownGroup = "minecraft:test_group"
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	decoded, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	if got := decoded.Components.UseCooldown.CooldownGroup; got != "minecraft:test_group" {
		t.Fatalf("cooldown group = %q", got)
	}
}

func TestDefaultComponents(t *testing.T) {
	appleID := items.ID("minecraft:apple")
	apple := items.DefaultComponents(appleID)
	if apple == nil {
		t.Fatal("DefaultComponents(apple) = nil")
	}
	if apple.Food == nil {
		t.Error("apple should have Food component")
	} else if apple.Food.Nutrition != 4 || apple.Food.CanAlwaysEat {
		t.Errorf("apple food = %+v, want nutrition 4 and can_always_eat false", apple.Food)
	}
	for _, name := range []string{"minecraft:golden_apple", "minecraft:enchanted_golden_apple", "minecraft:chorus_fruit", "minecraft:honey_bottle", "minecraft:suspicious_stew"} {
		component := items.DefaultComponents(items.ID(name))
		if component == nil || component.Food == nil || !component.Food.CanAlwaysEat {
			t.Errorf("%s should have can_always_eat food", name)
		}
	}
	goldenApple := items.DefaultComponents(items.ID("minecraft:golden_apple"))
	if goldenApple.Consumable == nil || len(goldenApple.Consumable.OnConsumeEffects) != 1 || len(goldenApple.Consumable.OnConsumeEffects[0].Effects) != 2 {
		t.Fatalf("golden apple consumable = %+v", goldenApple.Consumable)
	}
	honey := items.DefaultComponents(items.ID("minecraft:honey_bottle"))
	if honey.Consumable == nil || honey.Consumable.ConsumeSeconds != 2 || len(honey.Consumable.OnConsumeEffects[0].EffectIDs) != 1 {
		t.Fatalf("honey consumable = %+v", honey.Consumable)
	}

	swordID := items.ID("minecraft:diamond_sword")
	sword := items.DefaultComponents(swordID)
	if sword == nil {
		t.Fatal("DefaultComponents(diamond_sword) = nil")
	}
	if sword.MaxDamage != 1561 {
		t.Errorf("diamond_sword max damage = %d, want 1561", sword.MaxDamage)
	}
}

func TestBlocksAttacksComponentRoundTrip(t *testing.T) {
	shield := items.DefaultComponents(items.ID("minecraft:shield"))
	if shield == nil || shield.BlocksAttacks == nil {
		t.Fatal("shield blocks_attacks default is missing")
	}
	defaults := shield.BlocksAttacks
	if defaults.BlockDelaySeconds != 0.25 || defaults.DisableCooldownScale != 1 ||
		len(defaults.DamageReductions) != 1 || defaults.DamageReductions[0].HorizontalBlockingAngle != 90 ||
		defaults.ItemDamage != (items.DamageSpec{Threshold: 3, Base: 1, Factor: 1}) ||
		len(defaults.BypassedBy) != 1 || defaults.BypassedBy[0] != "#minecraft:bypasses_shield" {
		t.Fatalf("shield blocks_attacks = %+v", defaults)
	}

	patch := nbt.Compound{"minecraft:blocks_attacks": nbt.Compound{
		"block_delay_seconds":    nbt.Float(0.5),
		"disable_cooldown_scale": nbt.Float(2),
		"damage_reductions": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{
				"horizontal_blocking_angle": nbt.Float(60),
				"type":                      nbt.List{ElementType: nbt.TagString, Elements: []nbt.Tag{nbt.String("minecraft:generic")}},
				"base":                      nbt.Float(2),
				"factor":                    nbt.Float(0.5),
			},
		}},
		"item_damage": nbt.Compound{"threshold": nbt.Float(4), "base": nbt.Float(2), "factor": nbt.Float(0.5)},
		"bypassed_by": nbt.String("#minecraft:bypasses_armor"),
		"block_sound": nbt.Compound{"sound_id": nbt.String("minecraft:test.block"), "range": nbt.Float(8)},
	}}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil {
		t.Fatalf("ReadSlot: %v", err)
	}
	stack, err := items.FromSlot(decoded)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	component := stack.Components.BlocksAttacks
	if component == nil || len(component.DamageReductions) != 1 ||
		component.DamageReductions[0].Types[0] != "minecraft:generic" ||
		component.ItemDamage != (items.DamageSpec{Threshold: 4, Base: 2, Factor: 0.5}) ||
		component.BlockSound != "minecraft:test.block" || component.BlockSoundRange == nil || *component.BlockSoundRange != 8 {
		t.Fatalf("decoded blocks_attacks = %+v", component)
	}
	roundTrip := items.ComponentsToNBT(decoded).GetCompound("minecraft:blocks_attacks")
	if roundTrip.GetFloat("block_delay_seconds") != 0.5 ||
		roundTrip.GetCompound("item_damage").GetFloat("threshold") != 4 ||
		roundTrip.GetCompound("block_sound").GetFloat("range") != 8 {
		t.Fatalf("blocks_attacks NBT = %+v", roundTrip)
	}
}

func TestDeathProtectionComponentRoundTrip(t *testing.T) {
	defaults := items.DefaultComponents(items.ID("minecraft:totem_of_undying"))
	if defaults == nil || defaults.DeathProtection == nil || len(defaults.DeathProtection.DeathEffects) != 2 {
		t.Fatalf("totem death protection defaults = %+v", defaults)
	}
	apply := defaults.DeathProtection.DeathEffects[1]
	if apply.Type != "minecraft:apply_effects" || len(apply.Effects) != 3 || apply.Effects[0].ID != "minecraft:regeneration" {
		t.Fatalf("totem death effects = %+v", defaults.DeathProtection.DeathEffects)
	}

	patch := nbt.Compound{"minecraft:death_protection": nbt.Compound{
		"death_effects": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
			nbt.Compound{"type": nbt.String("minecraft:clear_all_effects")},
			nbt.Compound{"type": nbt.String("minecraft:apply_effects"), "effects": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
				nbt.Compound{"id": nbt.String("minecraft:regeneration"), "duration": nbt.Int(40), "amplifier": nbt.Int(2)},
			}}},
		}},
	}}
	slot, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil {
		t.Fatalf("ReadSlot: %v", err)
	}
	stack, err := items.FromSlot(decoded)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	protection := stack.Components.DeathProtection
	if protection == nil || len(protection.DeathEffects) != 2 || protection.DeathEffects[1].Effects[0].Amplifier != 2 {
		t.Fatalf("decoded death protection = %+v", protection)
	}
	roundTrip := items.ComponentsToNBT(decoded).GetCompound("minecraft:death_protection")
	if got := roundTrip.GetList("death_effects").Elements; len(got) != 2 {
		t.Fatalf("death protection NBT = %+v", roundTrip)
	}
}

func TestAttackRangeDefaultsWireAndNBTRoundTrip(t *testing.T) {
	itemID := items.ID("minecraft:diamond_spear")
	defaults := items.DefaultComponents(itemID)
	if defaults.AttackRange == nil || defaults.AttackRange.MinReach != 2 ||
		defaults.AttackRange.MaxReach != 4.5 || defaults.AttackRange.HitboxMargin != 0.125 {
		t.Fatalf("diamond spear attack range = %+v", defaults.AttackRange)
	}

	stack := items.NewStack(itemID, 1)
	stack.Components.AttackRange = &items.AttackRange{
		MinReach: 1, MaxReach: 8, MinCreativeReach: 1,
		MaxCreativeReach: 10, HitboxMargin: 0.25, MobFactor: 0.75,
	}
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	wire := ns.NewWriter()
	if err := wire.WriteSlot(slot); err != nil {
		t.Fatalf("WriteSlot: %v", err)
	}
	decoded, err := ns.NewReader(wire.Bytes()).ReadSlot(items.Decoder())
	if err != nil {
		t.Fatalf("ReadSlot: %v", err)
	}
	decodedStack, err := items.FromSlot(decoded)
	if err != nil || decodedStack.Components.AttackRange == nil ||
		decodedStack.Components.AttackRange.MaxReach != 8 ||
		decodedStack.Components.AttackRange.MobFactor != 0.75 {
		t.Fatalf("decoded attack range = %+v, %v", decodedStack.Components.AttackRange, err)
	}

	patch := items.ComponentsToNBT(decoded)
	restored, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(itemID), 1), patch)
	if err != nil {
		t.Fatalf("ApplyNBTComponents: %v", err)
	}
	restoredStack, err := items.FromSlot(restored)
	if err != nil || restoredStack.Components.AttackRange == nil ||
		restoredStack.Components.AttackRange.MaxCreativeReach != 10 ||
		restoredStack.Components.AttackRange.HitboxMargin != 0.25 {
		t.Fatalf("restored attack range = %+v, %v", restoredStack.Components.AttackRange, err)
	}
}

func TestToolComponentWireRoundTripPreservesDefaultsAndOptionals(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:stick"), 1)
	want := &items.Tool{
		DefaultMiningSpeed: 0.75,
		DamagePerBlock:     2,
		Rules: []items.ToolRule{
			{
				Blocks:          []string{"#minecraft:mineable/pickaxe"},
				CorrectForDrops: false, HasCorrectForDrops: true,
			},
			{
				Blocks: []string{"minecraft:stone", "minecraft:dirt"},
				Speed:  3, HasSpeed: true,
			},
		},
	}
	stack.Components.Tool = want
	stack.Components.SetPresent(items.ComponentTool)
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatalf("ToSlot: %v", err)
	}
	decoded, err := items.FromSlot(slot)
	if err != nil {
		t.Fatalf("FromSlot: %v", err)
	}
	if !reflect.DeepEqual(decoded.Components.Tool, want) {
		t.Fatalf("decoded tool = %#v, want %#v", decoded.Components.Tool, want)
	}
}

func TestComponentConstants(t *testing.T) {
	if items.ComponentDamage < 0 {
		t.Error("ComponentDamage should be non-negative")
	}
	if items.ComponentMaxDamage < 0 {
		t.Error("ComponentMaxDamage should be non-negative")
	}
	if items.ComponentFood < 0 {
		t.Error("ComponentFood should be non-negative")
	}
	if items.MaxComponentID < items.ComponentFood {
		t.Error("MaxComponentID should be >= ComponentFood")
	}
}

func TestToolComponentNBTRoundTripPreservesOptionalRules(t *testing.T) {
	stack := items.NewStack(items.ID("minecraft:diamond_pickaxe"), 1)
	stack.Components.Tool = &items.Tool{
		Rules: []items.ToolRule{
			{Blocks: []string{"#minecraft:mineable/pickaxe"}, Speed: 8, HasSpeed: true},
			{Blocks: []string{"minecraft:obsidian", "minecraft:crying_obsidian"}, CorrectForDrops: true, HasCorrectForDrops: true},
		},
		DefaultMiningSpeed:         1,
		DamagePerBlock:             2,
		CanDestroyBlocksInCreative: true,
	}
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatal(err)
	}
	component := items.ComponentsToNBT(slot).GetCompound("minecraft:tool")
	if _, present := component["default_mining_speed"]; present {
		t.Fatal("default mining speed should be omitted")
	}
	if _, present := component["can_destroy_blocks_in_creative"]; present {
		t.Fatal("default creative flag should be omitted")
	}
	rules := component.GetList("rules")
	if len(rules.Elements) != 2 {
		t.Fatalf("rule count = %d, want 2", len(rules.Elements))
	}
	first := rules.Elements[0].(nbt.Compound)
	if first.GetString("blocks") != "#minecraft:mineable/pickaxe" {
		t.Fatalf("tag holder = %q", first.GetString("blocks"))
	}
	if _, present := first["correct_for_drops"]; present {
		t.Fatal("absent correct-for-drops was encoded")
	}
	restored, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:diamond_pickaxe")), 1), nbt.Compound{
		"minecraft:tool": component,
	})
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := items.FromSlot(restored)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded.Components.Tool, stack.Components.Tool) {
		t.Fatalf("restored tool = %#v, want %#v", decoded.Components.Tool, stack.Components.Tool)
	}
}

func TestToolComponentNBTRejectsInvalidValues(t *testing.T) {
	tests := map[string]nbt.Compound{
		"negative damage": {
			"rules":            nbt.List{ElementType: nbt.TagCompound},
			"damage_per_block": nbt.Int(-1),
		},
		"zero rule speed": {
			"rules": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
				nbt.Compound{
					"blocks": nbt.String("#minecraft:mineable/pickaxe"),
					"speed":  nbt.Float(0),
				},
			}},
		},
		"empty holder set": {
			"rules": nbt.List{ElementType: nbt.TagCompound, Elements: []nbt.Tag{
				nbt.Compound{
					"blocks": nbt.List{ElementType: nbt.TagString},
				},
			}},
		},
	}
	for name, tool := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := items.ApplyNBTComponents(ns.NewSlot(ns.VarInt(items.ID("minecraft:diamond_pickaxe")), 1), nbt.Compound{
				"minecraft:tool": tool,
			})
			if err == nil {
				t.Fatal("invalid tool component was accepted")
			}
		})
	}
}

func TestSlotMaxStackSizeUsesComponentPatch(t *testing.T) {
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
	writer := ns.NewWriter()
	writer.WriteVarInt(7)
	slot.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentMaxStackSize, Data: writer.Bytes()}}
	if got := items.SlotMaxStackSize(slot); got != 7 {
		t.Fatalf("custom slot max = %d, want 7", got)
	}
	slot.Components.Add = nil
	slot.Components.Remove = []ns.VarInt{items.ComponentMaxStackSize}
	if got := items.SlotMaxStackSize(slot); got != 1 {
		t.Fatalf("removed slot max = %d, want fallback 1", got)
	}
}

func TestSlotResistsDamageUsesComponentPatch(t *testing.T) {
	netherite := ns.NewSlot(ns.VarInt(items.ID("minecraft:netherite_ingot")), 1)
	if !items.SlotResistsDamage(netherite, "minecraft:lava") {
		t.Fatal("netherite ingot should resist fire")
	}
	netherite.Components.Remove = []ns.VarInt{items.ComponentDamageResistant}
	if items.SlotResistsDamage(netherite, "minecraft:lava") {
		t.Fatal("removed resistance still applied")
	}

	stone := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 1)
	writer := ns.NewWriter()
	writer.WriteVarInt(0)
	if err := writer.WriteString("minecraft:is_fire"); err != nil {
		t.Fatal(err)
	}
	stone.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentDamageResistant, Data: writer.Bytes()}}
	if !items.SlotResistsDamage(stone, "minecraft:lava") {
		t.Fatal("patched stone should resist fire")
	}
}

func TestSlotNBTRoundTrip(t *testing.T) {
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:stone")), 3)
	writer := ns.NewWriter()
	writer.WriteVarInt(7)
	slot.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentMaxStackSize, Data: writer.Bytes()}}

	restored, err := items.SlotFromNBT(items.SlotToNBT(slot))
	if err != nil {
		t.Fatal(err)
	}
	if !items.SameStack(restored, slot) || restored.Count != slot.Count {
		t.Fatalf("restored slot = %+v, want %+v", restored, slot)
	}
}

func TestStackEquippableSlotUsesDefaultsAndPatch(t *testing.T) {
	helmet := ns.NewSlot(ns.VarInt(items.ID("minecraft:diamond_helmet")), 1)
	if slot, ok := items.StackEquippableSlot(helmet); !ok || slot != "head" {
		t.Fatalf("helmet equippable slot = %q/%v", slot, ok)
	}
	shield := ns.NewSlot(ns.VarInt(items.ID("minecraft:shield")), 1)
	if slot, ok := items.StackEquippableSlot(shield); !ok || slot != "offhand" {
		t.Fatalf("shield equippable slot = %q/%v", slot, ok)
	}
	helmet.Components.Remove = []ns.VarInt{items.ComponentEquippable}
	if slot, ok := items.StackEquippableSlot(helmet); ok {
		t.Fatalf("removed equippable slot = %q", slot)
	}
	stack := items.NewStack(items.ID("minecraft:stone"), 1)
	stack.Components.Equippable = &items.Equippable{Slot: "head"}
	stone, err := stack.ToSlot()
	if err != nil {
		t.Fatal(err)
	}
	if slot, ok := items.StackEquippableSlot(stone); !ok || slot != "head" {
		t.Fatalf("patched equippable slot = %q/%v", slot, ok)
	}
}

func TestApplyNBTComponentsAcceptsJSONDerivedLists(t *testing.T) {
	tag, err := nbt.FromJSON([]byte(`{"minecraft:lore":["heart","explorer"]}`))
	if err != nil {
		t.Fatal(err)
	}
	patch, ok := tag.(nbt.Compound)
	if !ok {
		t.Fatalf("JSON patch type = %T, want nbt.Compound", tag)
	}
	slot := ns.NewSlot(ns.VarInt(items.ID("minecraft:decorated_pot")), 1)
	applied, err := items.ApplyNBTComponents(slot, patch)
	if err != nil {
		t.Fatal(err)
	}
	lore := items.ComponentsToNBT(applied).GetList("minecraft:lore")
	if len(lore.Elements) != 2 ||
		lore.Elements[0] != nbt.String("heart") ||
		lore.Elements[1] != nbt.String("explorer") {
		t.Fatalf("lore = %#v", lore)
	}
}

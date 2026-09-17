package items

// Component codec registration - registers complex component codecs at init time.
// Simple codecs (varint, float32, string, empty) are auto-generated in item_components_codec_gen.go.

import (
	"fmt"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	// struct components with dedicated codecs
	RegisterCodec(ComponentCustomName, &customNameCodec{})
	RegisterCodec(ComponentItemName, &itemNameCodec{})
	RegisterCodec(ComponentAttributeModifiers, &attributeModifiersCodec{})
	RegisterCodec(ComponentRarity, &rarityCodec{})
	RegisterCodec(ComponentBreakSound, newHolderStringCodec(ComponentBreakSound, "minecraft:sound_event", func(c *Components) string { return c.BreakSound }, func(c *Components, value string) { c.BreakSound = value }, copySoundEventDirect))
	RegisterCodec(ComponentInstrument, newHolderStringCodec(ComponentInstrument, "minecraft:instrument", func(c *Components) string { return c.Instrument }, func(c *Components, value string) { c.Instrument = value }, copyInstrumentDirect))
	RegisterCodec(ComponentJukeboxPlayable, newHolderStringCodec(ComponentJukeboxPlayable, "minecraft:jukebox_song", func(c *Components) string { return c.JukeboxPlayable }, func(c *Components, value string) { c.JukeboxPlayable = value }, copyJukeboxSongDirect))
	RegisterCodec(ComponentProvidesTrimMaterial, newHolderStringCodec(ComponentProvidesTrimMaterial, "minecraft:trim_material", func(c *Components) string { return c.ProvidesTrimMaterial }, func(c *Components, value string) { c.ProvidesTrimMaterial = value }, copyTrimMaterialDirect))

	// lore, enchantments, and tool codecs (full implementations)
	RegisterCodec(ComponentLore, &loreCodec{})
	RegisterCodec(ComponentEnchantments, &enchantmentsCodec{
		get: func(c *Components) map[string]int32 { return c.Enchantments },
		set: func(c *Components, v map[string]int32) { c.Enchantments = v },
	})
	RegisterCodec(ComponentStoredEnchantments, &enchantmentsCodec{
		get: func(c *Components) map[string]int32 { return c.StoredEnchantments },
		set: func(c *Components, v map[string]int32) { c.StoredEnchantments = v },
	})
	RegisterCodec(ComponentTool, &toolCodec{})

	// complex passthrough codecs - these have custom decoders
	// simple passthroughs (varint, bool, string, empty, int32, nbt, holderSet, slot, slotList)
	// are registered in item_components_codec_gen.go
	RegisterCodec(ComponentCanBreak, &passthroughCodec{decode: decodeBlockPredicatesWire})
	RegisterCodec(ComponentCanPlaceOn, &passthroughCodec{decode: decodeBlockPredicatesWire})
	RegisterCodec(ComponentCustomModelData, &passthroughCodec{decode: decodeCustomModelDataWire})
	RegisterCodec(ComponentBundleContents, &passthroughCodec{decode: decodeTemplateListWire})
	RegisterCodec(ComponentChargedProjectiles, &passthroughCodec{decode: decodeTemplateListWire})
	RegisterCodec(ComponentContainer, &passthroughCodec{decode: decodeOptionalTemplateListWire})
	RegisterCodec(ComponentSulfurCubeContent, &passthroughCodec{decode: copyItemStackTemplate})
	RegisterCodec(ComponentEntityData, entityDataCodec{})
	RegisterCodec(ComponentBlockEntityData, &passthroughCodec{decode: decodeTypedEntityDataWire})
	RegisterCodec(ComponentUseRemainder, &passthroughCodec{decode: copyItemStackTemplate})
	RegisterCodec(ComponentConsumable, consumableCodec{})
	RegisterCodec(ComponentUseEffects, &passthroughCodec{decode: decodeUseEffectsWire})
	RegisterCodec(ComponentAttackRange, attackRangeCodec{})
	RegisterCodec(ComponentEquippable, equippableCodec{})
	RegisterCodec(ComponentDeathProtection, deathProtectionCodec{})
	RegisterCodec(ComponentBlocksAttacks, blocksAttacksCodec{})
	RegisterCodec(ComponentKineticWeapon, &passthroughCodec{decode: decodeKineticWeaponWire})
	RegisterCodec(ComponentPiercingWeapon, &passthroughCodec{decode: decodePiercingWeaponWire})
	RegisterCodec(ComponentPotionContents, &passthroughCodec{decode: decodePotionContentsWire})
	RegisterCodec(ComponentSuspiciousStewEffects, &passthroughCodec{decode: decodeSuspiciousStewWire})
	RegisterCodec(ComponentWritableBookContent, &passthroughCodec{decode: decodeWritableBookWire})
	RegisterCodec(ComponentWrittenBookContent, &passthroughCodec{decode: decodeWrittenBookWire})
	RegisterCodec(ComponentTrim, trimCodec{})
	RegisterCodec(ComponentPaintingVariant, paintingVariantCodec{})
	RegisterCodec(ComponentDye, dyeCodec{})
	RegisterCodec(ComponentDyedColor, dyedColorCodec{})
	RegisterCodec(ComponentBannerPatterns, bannerPatternsCodec{})
	RegisterCodec(ComponentPotDecorations, potDecorationsCodec{})
	RegisterCodec(ComponentFireworkExplosion, fireworkExplosionCodec{})
	RegisterCodec(ComponentFireworks, fireworksCodec{})
	RegisterCodec(ComponentBaseColor, baseColorCodec{})
	RegisterCodec(ComponentRecipes, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentLodestoneTracker, &passthroughCodec{decode: decodeLodestoneWire})
	RegisterCodec(ComponentProfile, &passthroughCodec{decode: decodeProfileWire})
	RegisterCodec(ComponentBlockState, &passthroughCodec{decode: decodeBlockStateWire})
	RegisterCodec(ComponentBees, &passthroughCodec{decode: decodeBeesWire})
}

type holderStringCodec struct {
	id       int32
	registry string
	plain    bool // registry holder, without the inline-value discriminator
	get      func(*Components) string
	set      func(*Components, string)
	direct   func(*ns.PacketBuffer, *ns.PacketBuffer) error
}

func newHolderStringCodec(id int32, registry string, get func(*Components) string, set func(*Components, string), direct func(*ns.PacketBuffer, *ns.PacketBuffer) error) *holderStringCodec {
	return &holderStringCodec{id: id, registry: registry, get: get, set: set, direct: direct}
}

func (codec *holderStringCodec) DecodeWire(reader *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	writer.WriteVarInt(holderID)
	if holderID == 0 && !codec.plain {
		if err := codec.direct(reader, writer); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}

func (codec *holderStringCodec) Apply(components *Components, data []byte) error {
	reader := ns.NewReader(data)
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	if holderID < 0 {
		return fmt.Errorf("invalid %s holder %d", codec.registry, holderID)
	}
	if holderID == 0 && !codec.plain {
		validated := ns.NewWriter()
		validated.WriteVarInt(0)
		if err := codec.direct(reader, validated); err != nil {
			return err
		}
		if components.rawPassthrough == nil {
			components.rawPassthrough = make(map[int32][]byte)
		}
		components.rawPassthrough[codec.id] = append([]byte(nil), data...)
		codec.set(components, "")
		return nil
	}
	entryID := int(holderID)
	if !codec.plain {
		entryID--
	}
	entries := registries.SynchronizedEntries[codec.registry]
	if codec.registry == "minecraft:sound_event" {
		name := registries.SoundEvent.ByID(int32(entryID))
		if name == "" {
			return fmt.Errorf("unknown %s holder %d", codec.registry, holderID)
		}
		codec.set(components, name)
	} else if entryID >= len(entries) {
		return fmt.Errorf("unknown %s holder %d", codec.registry, holderID)
	} else {
		codec.set(components, entries[entryID])
	}
	delete(components.rawPassthrough, codec.id)
	return nil
}

func (codec *holderStringCodec) Clear(components *Components) {
	codec.set(components, "")
	delete(components.rawPassthrough, codec.id)
}

func (codec *holderStringCodec) Differs(components, defaults *Components) (bool, bool) {
	if _, direct := components.rawPassthrough[codec.id]; direct {
		return true, true
	}
	value := codec.get(components)
	defaultValue := ""
	if defaults != nil {
		defaultValue = codec.get(defaults)
	}
	if value == defaultValue {
		return false, value != ""
	}
	return true, value != ""
}

func (codec *holderStringCodec) Encode(components *Components) ([]byte, error) {
	if direct, ok := components.rawPassthrough[codec.id]; ok {
		return append([]byte(nil), direct...), nil
	}
	value := codec.get(components)
	var entryID int
	if codec.registry == "minecraft:sound_event" {
		entryID = int(registries.SoundEvent.Get(value))
	} else {
		entryID = registryEntryID(codec.registry, value)
	}
	if entryID < 0 {
		return nil, fmt.Errorf("unknown %s entry %q", codec.registry, value)
	}
	writer := ns.NewWriter()
	if !codec.plain {
		entryID++
	}
	writer.WriteVarInt(ns.VarInt(entryID))
	return writer.Bytes(), nil
}

func copySoundEventDirect(reader, writer *ns.PacketBuffer) error {
	if err := writer.CopyString(reader, maxStringLen); err != nil {
		return err
	}
	hasRange, err := reader.ReadBool()
	if err != nil {
		return err
	}
	writer.WriteBool(hasRange)
	if hasRange {
		return writer.CopyFloat32(reader)
	}
	return nil
}

func copySoundEventHolder(reader, writer *ns.PacketBuffer) error {
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	writer.WriteVarInt(holderID)
	if holderID == 0 {
		return copySoundEventDirect(reader, writer)
	}
	return nil
}

func copyInstrumentDirect(reader, writer *ns.PacketBuffer) error {
	if err := copySoundEventHolder(reader, writer); err != nil {
		return err
	}
	if err := writer.CopyFloat32(reader); err != nil {
		return err
	}
	if err := writer.CopyFloat32(reader); err != nil {
		return err
	}
	return copyNBT(reader, writer)
}

func copyJukeboxSongDirect(reader, writer *ns.PacketBuffer) error {
	if err := copySoundEventHolder(reader, writer); err != nil {
		return err
	}
	if err := copyNBT(reader, writer); err != nil {
		return err
	}
	if err := writer.CopyFloat32(reader); err != nil {
		return err
	}
	return writer.CopyVarInt(reader)
}

func copyTrimMaterialDirect(reader, writer *ns.PacketBuffer) error {
	if err := writer.CopyString(reader, maxStringLen); err != nil {
		return err
	}
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return fmt.Errorf("invalid trim material override count %d", count)
	}
	writer.WriteVarInt(count)
	for range int(count) {
		if err := writer.CopyString(reader, maxStringLen); err != nil {
			return err
		}
		if err := writer.CopyString(reader, maxStringLen); err != nil {
			return err
		}
	}
	return copyNBT(reader, writer)
}

type fireworkExplosionCodec struct{}

func (fireworkExplosionCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := copyFireworkExplosion(buf, writer); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}
func (fireworkExplosionCodec) Apply(components *Components, data []byte) error {
	explosion, err := readFireworkExplosion(ns.NewReader(data))
	if err != nil {
		return err
	}
	components.FireworkExplosion = &explosion
	return nil
}
func (fireworkExplosionCodec) Clear(components *Components) { components.FireworkExplosion = nil }
func (fireworkExplosionCodec) Differs(components, defaults *Components) (bool, bool) {
	return !equalFireworkExplosionPointers(components.FireworkExplosion, defaults.FireworkExplosion), components.FireworkExplosion != nil
}
func (fireworkExplosionCodec) Encode(components *Components) ([]byte, error) {
	if components.FireworkExplosion == nil {
		return nil, fmt.Errorf("firework explosion is absent")
	}
	writer := ns.NewWriter()
	if err := writeFireworkExplosion(writer, *components.FireworkExplosion); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

type fireworksCodec struct{}

func (fireworksCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := writer.CopyVarInt(buf); err != nil {
		return nil, err
	}
	count, err := buf.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nil, fmt.Errorf("invalid firework explosion count %d", count)
	}
	if err := writer.WriteVarInt(count); err != nil {
		return nil, err
	}
	for range int(count) {
		if err := copyFireworkExplosion(buf, writer); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}
func (fireworksCodec) Apply(components *Components, data []byte) error {
	reader := ns.NewReader(data)
	flight, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return fmt.Errorf("invalid firework explosion count %d", count)
	}
	explosions := make([]FireworkExplosion, count)
	for i := range explosions {
		explosions[i], err = readFireworkExplosion(reader)
		if err != nil {
			return err
		}
	}
	components.Fireworks = &Fireworks{FlightDuration: int32(flight)}
	components.FireworkExplosions = explosions
	return nil
}
func (fireworksCodec) Clear(components *Components) {
	components.Fireworks = nil
	components.FireworkExplosions = nil
}
func (fireworksCodec) Differs(components, defaults *Components) (bool, bool) {
	if components.Fireworks == nil || defaults.Fireworks == nil {
		return components.Fireworks != nil || defaults.Fireworks != nil, components.Fireworks != nil
	}
	if components.Fireworks.FlightDuration != defaults.Fireworks.FlightDuration || len(components.FireworkExplosions) != len(defaults.FireworkExplosions) {
		return true, true
	}
	for i := range components.FireworkExplosions {
		if !equalFireworkExplosion(components.FireworkExplosions[i], defaults.FireworkExplosions[i]) {
			return true, true
		}
	}
	return false, true
}
func (fireworksCodec) Encode(components *Components) ([]byte, error) {
	if components.Fireworks == nil {
		return nil, fmt.Errorf("fireworks component is absent")
	}
	writer := ns.NewWriter()
	if err := writer.WriteVarInt(ns.VarInt(components.Fireworks.FlightDuration)); err != nil {
		return nil, err
	}
	if err := writer.WriteVarInt(ns.VarInt(len(components.FireworkExplosions))); err != nil {
		return nil, err
	}
	for _, explosion := range components.FireworkExplosions {
		if err := writeFireworkExplosion(writer, explosion); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}

func readFireworkExplosion(reader *ns.PacketBuffer) (FireworkExplosion, error) {
	shape, err := reader.ReadVarInt()
	if err != nil || shape < 0 || shape > 4 {
		return FireworkExplosion{}, fmt.Errorf("invalid firework shape %d", shape)
	}
	colors, err := readFireworkColors(reader)
	if err != nil {
		return FireworkExplosion{}, err
	}
	fadeColors, err := readFireworkColors(reader)
	if err != nil {
		return FireworkExplosion{}, err
	}
	trail, err := reader.ReadBool()
	if err != nil {
		return FireworkExplosion{}, err
	}
	twinkle, err := reader.ReadBool()
	if err != nil {
		return FireworkExplosion{}, err
	}
	return FireworkExplosion{Shape: int32(shape), Colors: colors, FadeColors: fadeColors, HasTrail: bool(trail), HasTwinkle: bool(twinkle)}, nil
}

func readFireworkColors(reader *ns.PacketBuffer) ([]int32, error) {
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nil, fmt.Errorf("invalid firework color count %d", count)
	}
	colors := make([]int32, count)
	for i := range colors {
		color, err := reader.ReadInt32()
		if err != nil {
			return nil, err
		}
		colors[i] = int32(color)
	}
	return colors, nil
}

func writeFireworkExplosion(writer *ns.PacketBuffer, explosion FireworkExplosion) error {
	if explosion.Shape < 0 || explosion.Shape > 4 || len(explosion.Colors) > 256 || len(explosion.FadeColors) > 256 {
		return fmt.Errorf("invalid firework explosion")
	}
	if err := writer.WriteVarInt(ns.VarInt(explosion.Shape)); err != nil {
		return err
	}
	for _, colors := range [][]int32{explosion.Colors, explosion.FadeColors} {
		if err := writer.WriteVarInt(ns.VarInt(len(colors))); err != nil {
			return err
		}
		for _, color := range colors {
			if err := writer.WriteInt32(ns.Int32(color)); err != nil {
				return err
			}
		}
	}
	if err := writer.WriteBool(ns.Boolean(explosion.HasTrail)); err != nil {
		return err
	}
	return writer.WriteBool(ns.Boolean(explosion.HasTwinkle))
}

func equalFireworkExplosionPointers(a, b *FireworkExplosion) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return equalFireworkExplosion(*a, *b)
}

func equalFireworkExplosion(a, b FireworkExplosion) bool {
	if a.Shape != b.Shape || a.HasTrail != b.HasTrail || a.HasTwinkle != b.HasTwinkle || len(a.Colors) != len(b.Colors) || len(a.FadeColors) != len(b.FadeColors) {
		return false
	}
	for i := range a.Colors {
		if a.Colors[i] != b.Colors[i] {
			return false
		}
	}
	for i := range a.FadeColors {
		if a.FadeColors[i] != b.FadeColors[i] {
			return false
		}
	}
	return true
}

type potDecorationsCodec struct{}

func (potDecorationsCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := decodePotDecorationsWire(buf, writer); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func (potDecorationsCodec) Apply(components *Components, data []byte) error {
	reader := ns.NewReader(data)
	decorations := &PotDecorations{}
	for _, side := range []**ItemStack{&decorations.Back, &decorations.Left, &decorations.Right, &decorations.Front} {
		present, err := reader.ReadBool()
		if err != nil {
			return err
		}
		if present {
			slot, err := ReadItemStackTemplate(reader)
			if err != nil {
				return err
			}
			*side, err = FromSlot(slot)
			if err != nil {
				return err
			}
		}
	}
	components.PotDecorations = decorations
	return nil
}

func (potDecorationsCodec) Clear(components *Components) { components.PotDecorations = nil }
func (potDecorationsCodec) Differs(components, defaults *Components) (bool, bool) {
	if components.PotDecorations == nil || defaults.PotDecorations == nil {
		return components.PotDecorations != nil || defaults.PotDecorations != nil, components.PotDecorations != nil
	}
	a, b := components.PotDecorations, defaults.PotDecorations
	return !equalDecoration(a.Back, b.Back) || !equalDecoration(a.Left, b.Left) ||
		!equalDecoration(a.Right, b.Right) || !equalDecoration(a.Front, b.Front), true
}
func (potDecorationsCodec) Encode(components *Components) ([]byte, error) {
	if components.PotDecorations == nil {
		return nil, fmt.Errorf("pot decorations are absent")
	}
	writer := ns.NewWriter()
	for _, side := range []*ItemStack{
		components.PotDecorations.Back, components.PotDecorations.Left,
		components.PotDecorations.Right, components.PotDecorations.Front,
	} {
		writer.WriteBool(ns.Boolean(side != nil))
		if side == nil {
			continue
		}
		slot, err := side.ToSlot()
		if err != nil {
			return nil, err
		}
		if err := WriteItemStackTemplate(writer, slot); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}

type baseColorCodec struct{}

func (baseColorCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := writer.CopyVarInt(buf); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}
func (baseColorCodec) Apply(components *Components, data []byte) error {
	color, err := ns.NewReader(data).ReadVarInt()
	if err != nil || color < 0 || color >= 16 {
		return fmt.Errorf("unknown base color id %d", color)
	}
	value := int32(color)
	components.BaseColor = &value
	return nil
}
func (baseColorCodec) Clear(components *Components) { components.BaseColor = nil }
func (baseColorCodec) Differs(components, defaults *Components) (bool, bool) {
	return !equalInt32Pointers(components.BaseColor, defaults.BaseColor), components.BaseColor != nil
}
func (baseColorCodec) Encode(components *Components) ([]byte, error) {
	if components.BaseColor == nil || *components.BaseColor < 0 || *components.BaseColor >= 16 {
		return nil, fmt.Errorf("invalid base color")
	}
	writer := ns.NewWriter()
	if err := writer.WriteVarInt(ns.VarInt(*components.BaseColor)); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

type bannerPatternsCodec struct{}

func (bannerPatternsCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := decodeBannerPatternsWire(buf, writer); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func (bannerPatternsCodec) Apply(components *Components, data []byte) error {
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return fmt.Errorf("invalid banner pattern count %d", count)
	}
	entries := registries.SynchronizedEntries["minecraft:banner_pattern"]
	layers := make([]BannerPatternLayer, count)
	for i := range layers {
		holder, err := reader.ReadVarInt()
		if err != nil || holder < 0 || holder > ns.VarInt(len(entries)) {
			return fmt.Errorf("unknown banner pattern holder %d", holder)
		}
		if holder == 0 {
			assetID, err := reader.ReadIdentifier()
			if err != nil {
				return err
			}
			translationKey, err := reader.ReadString(maxStringLen)
			if err != nil {
				return err
			}
			layers[i].AssetID = string(assetID)
			layers[i].TranslationKey = string(translationKey)
		} else {
			layers[i].Pattern = entries[holder-1]
		}
		color, err := reader.ReadVarInt()
		if err != nil || color < 0 || color >= 16 {
			return fmt.Errorf("unknown banner color id %d", color)
		}
		layers[i].Color = int32(color)
	}
	components.BannerPatterns = &layers
	return nil
}

func (bannerPatternsCodec) Clear(components *Components) { components.BannerPatterns = nil }
func (bannerPatternsCodec) Differs(components, defaults *Components) (bool, bool) {
	return !equalBannerPatterns(components.BannerPatterns, defaults.BannerPatterns), components.BannerPatterns != nil
}
func (bannerPatternsCodec) Encode(components *Components) ([]byte, error) {
	if components.BannerPatterns == nil {
		return nil, fmt.Errorf("banner patterns are absent")
	}
	writer := ns.NewWriter()
	if err := writer.WriteVarInt(ns.VarInt(len(*components.BannerPatterns))); err != nil {
		return nil, err
	}
	entries := registries.SynchronizedEntries["minecraft:banner_pattern"]
	for _, layer := range *components.BannerPatterns {
		pattern := -1
		for id, name := range entries {
			if name == layer.Pattern {
				pattern = id
				break
			}
		}
		if layer.Color < 0 || layer.Color >= 16 {
			return nil, fmt.Errorf("invalid banner color %d", layer.Color)
		}
		if pattern >= 0 {
			if err := writer.WriteVarInt(ns.VarInt(pattern + 1)); err != nil {
				return nil, err
			}
		} else {
			if layer.AssetID == "" || layer.TranslationKey == "" {
				return nil, fmt.Errorf("invalid direct banner pattern")
			}
			if err := writer.WriteVarInt(0); err != nil {
				return nil, err
			}
			if err := writer.WriteIdentifier(ns.Identifier(layer.AssetID)); err != nil {
				return nil, err
			}
			if err := writer.WriteString(ns.String(layer.TranslationKey)); err != nil {
				return nil, err
			}
		}
		if err := writer.WriteVarInt(ns.VarInt(layer.Color)); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}

func equalBannerPatterns(a, b *[]BannerPatternLayer) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(*a) != len(*b) {
		return false
	}
	for i := range *a {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}

type dyeCodec struct{}

func (dyeCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := writer.CopyVarInt(buf); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func (dyeCodec) Apply(components *Components, data []byte) error {
	id, err := ns.NewReader(data).ReadVarInt()
	if err != nil || id < 0 || id >= 16 {
		return fmt.Errorf("unknown dye id %d", id)
	}
	value := int32(id)
	components.Dye = &value
	return nil
}

func (dyeCodec) Clear(components *Components) { components.Dye = nil }
func (dyeCodec) Differs(components, defaults *Components) (bool, bool) {
	return !equalInt32Pointers(components.Dye, defaults.Dye), components.Dye != nil
}
func (dyeCodec) Encode(components *Components) ([]byte, error) {
	if components.Dye == nil || *components.Dye < 0 || *components.Dye >= 16 {
		return nil, fmt.Errorf("unknown dye id %v", components.Dye)
	}
	writer := ns.NewWriter()
	if err := writer.WriteVarInt(ns.VarInt(*components.Dye)); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

type dyedColorCodec struct{}

func (dyedColorCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := writer.CopyInt32(buf); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}
func (dyedColorCodec) Apply(components *Components, data []byte) error {
	color, err := ns.NewReader(data).ReadInt32()
	value := int32(color)
	components.DyedColor = &value
	return err
}
func (dyedColorCodec) Clear(components *Components) { components.DyedColor = nil }
func (dyedColorCodec) Differs(components, defaults *Components) (bool, bool) {
	return !equalInt32Pointers(components.DyedColor, defaults.DyedColor), components.DyedColor != nil
}
func (dyedColorCodec) Encode(components *Components) ([]byte, error) {
	if components.DyedColor == nil {
		return nil, fmt.Errorf("dyed color is absent")
	}
	writer := ns.NewWriter()
	if err := writer.WriteInt32(ns.Int32(*components.DyedColor)); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func equalInt32Pointers(a, b *int32) bool {
	return a == nil && b == nil || a != nil && b != nil && *a == *b
}

type paintingVariantCodec struct{}

func (paintingVariantCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	writer := ns.NewWriter()
	if err := writer.CopyVarInt(buf); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func (paintingVariantCodec) Apply(components *Components, data []byte) error {
	id, err := ns.NewReader(data).ReadVarInt()
	entries := registries.SynchronizedEntries["minecraft:painting_variant"]
	if err != nil || id < 0 || int(id) >= len(entries) {
		return fmt.Errorf("unknown painting variant id %d", id)
	}
	components.PaintingVariant = entries[id]
	return nil
}

func (paintingVariantCodec) Clear(components *Components) { components.PaintingVariant = "" }

func (paintingVariantCodec) Differs(components, defaults *Components) (bool, bool) {
	return components.PaintingVariant != defaults.PaintingVariant, components.PaintingVariant != ""
}

func (paintingVariantCodec) Encode(components *Components) ([]byte, error) {
	for id, name := range registries.SynchronizedEntries["minecraft:painting_variant"] {
		if name == components.PaintingVariant {
			writer := ns.NewWriter()
			if err := writer.WriteVarInt(ns.VarInt(id)); err != nil {
				return nil, err
			}
			return writer.Bytes(), nil
		}
	}
	return nil, fmt.Errorf("unknown painting variant %q", components.PaintingVariant)
}

func registerVarIntPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeVarIntWire})
}

func registerBoolPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeBoolWire})
}

func registerStringPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeStringWire})
}

func registerEmptyPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeEmptyWire})
}

func registerInt32Passthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeInt32Wire})
}

func registerNBTPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeNBTWire})
}

func registerHolderSetPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeHolderSetWire})
}

func registerSlotListPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeSlotListWire})
}

func registerSlotPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeSlotWire})
}

func decodeVarIntWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return w.CopyVarInt(buf)
}

func decodeInt32Wire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return w.CopyInt32(buf)
}

func decodeBoolWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return w.CopyBool(buf)
}

func decodeStringWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return w.CopyString(buf, maxStringLen)
}

func decodeEmptyWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return nil
}

func decodeNBTWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyNBT(buf, w)
}

func decodeBlockPredicatesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyVarIntPrefixedList(buf, w, copyBlockPredicate); err != nil {
		return err
	}
	return w.CopyBool(buf)
}

func decodeCustomModelDataWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	copyFloat32Fn := func(buf, w *ns.PacketBuffer) error { return w.CopyFloat32(buf) }
	copyBoolFn := func(buf, w *ns.PacketBuffer) error { return w.CopyBool(buf) }
	copyStringFn := func(buf, w *ns.PacketBuffer) error { return w.CopyString(buf, maxStringLen) }
	copyInt32Fn := func(buf, w *ns.PacketBuffer) error { return w.CopyInt32(buf) }

	if err := copyVarIntPrefixedList(buf, w, copyFloat32Fn); err != nil {
		return err
	}
	if err := copyVarIntPrefixedList(buf, w, copyBoolFn); err != nil {
		return err
	}
	if err := copyVarIntPrefixedList(buf, w, copyStringFn); err != nil {
		return err
	}
	return copyVarIntPrefixedList(buf, w, copyInt32Fn)
}

func decodeConsumableWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyFloat32(buf); err != nil { // consume seconds
		return err
	}
	if err := w.CopyVarInt(buf); err != nil { // animation
		return err
	}
	if err := copySoundEvent(buf, w); err != nil {
		return err
	}
	if err := w.CopyBool(buf); err != nil { // has particles
		return err
	}
	return copyVarIntPrefixedList(buf, w, copyConsumeEffect)
}

func decodeSlotWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copySlot(buf, w)
}

func decodeSlotListWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, copySlot)
}

func decodeUseEffectsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyBool(buf); err != nil { // can sprint
		return err
	}
	if err := w.CopyBool(buf); err != nil { // interact vibrations
		return err
	}
	return w.CopyFloat32(buf) // speed multiplier
}

func decodeHolderSetWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyHolderSet(buf, w)
}

func decodeEquippableWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyVarInt(buf); err != nil { // slot
		return err
	}
	if err := copySoundEvent(buf, w); err != nil {
		return err
	}
	if err := copyOptionalString(buf, w); err != nil { // optional asset
		return err
	}
	if err := copyOptionalIdentifier(buf, w); err != nil { // optional camera overlay
		return err
	}
	if err := copyOptionalHolderSet(buf, w); err != nil {
		return err
	}
	// dispensable, swappable, damages on hurt
	for range 3 {
		if err := w.CopyBool(buf); err != nil {
			return err
		}
	}
	// can be sheared (conditional)
	canBeSheared, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(canBeSheared)
	if canBeSheared {
		return copySoundEvent(buf, w)
	}
	return nil
}

func decodeDeathProtectionWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, copyConsumeEffect)
}

func decodeKineticWeaponWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyFloat32(buf); err != nil { // damage multiplier
		return err
	}
	// damage, dismount, knockback conditions
	for range 3 {
		if err := copyOptionalKineticConditions(buf, w); err != nil {
			return err
		}
	}
	if err := w.CopyFloat32(buf); err != nil { // forward movement
		return err
	}
	if err := w.CopyVarInt(buf); err != nil { // delay ticks
		return err
	}
	if err := copyOptionalSoundEvent(buf, w); err != nil { // sound
		return err
	}
	return copyOptionalSoundEvent(buf, w) // hit sound
}

func decodePiercingWeaponWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	// optional hit sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	return nil
}

func decodePotionContentsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyOptionalVarInt(buf, w); err != nil { // potion
		return err
	}
	if err := copyOptionalInt(buf, w); err != nil { // custom color
		return err
	}
	if err := copyVarIntPrefixedList(buf, w, copyStatusEffect); err != nil { // custom effects
		return err
	}
	return copyOptionalString(buf, w) // custom name
}

func copyItemStackTemplate(buf, writer *ns.PacketBuffer) error {
	stack, err := ReadItemStackTemplate(buf)
	if err != nil {
		return err
	}
	return WriteItemStackTemplate(writer, stack)
}

func decodeTemplateListWire(buf, writer *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, writer, copyItemStackTemplate)
}

func decodeOptionalTemplateListWire(buf, writer *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, writer, func(buf, writer *ns.PacketBuffer) error {
		hasStack, err := buf.ReadBool()
		if err != nil {
			return err
		}
		writer.WriteBool(hasStack)
		if !hasStack {
			return nil
		}
		return copyItemStackTemplate(buf, writer)
	})
}

func decodeTypedEntityDataWire(buf, writer *ns.PacketBuffer) error {
	if err := writer.CopyVarInt(buf); err != nil {
		return err
	}
	return copyNBT(buf, writer)
}

func decodeSuspiciousStewWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
		if err := w.CopyVarInt(buf); err != nil { // effect ID
			return err
		}
		return w.CopyVarInt(buf) // duration
	})
}

func decodeWritableBookWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
		if err := w.CopyString(buf, maxStringLen); err != nil { // raw content
			return err
		}
		return copyOptionalString(buf, w) // optional filtered content
	})
}

func decodeWrittenBookWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyString(buf, maxStringLen); err != nil { // raw title
		return err
	}
	if err := copyOptionalString(buf, w); err != nil { // filtered title
		return err
	}
	if err := w.CopyString(buf, maxStringLen); err != nil { // author
		return err
	}
	if err := w.CopyVarInt(buf); err != nil { // generation
		return err
	}
	if err := copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
		if err := copyNBT(buf, w); err != nil { // raw NBT
			return err
		}
		return copyOptionalNBT(buf, w) // optional filtered NBT
	}); err != nil {
		return err
	}
	return w.CopyBool(buf) // resolved
}

func decodeTrimWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// material
	if err := copyTrimMaterial(buf, w); err != nil {
		return err
	}
	// pattern
	if err := copyTrimPattern(buf, w); err != nil {
		return err
	}
	return nil
}

func decodeRecipesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
		return w.CopyString(buf, maxStringLen)
	})
}

func decodeLodestoneWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyOptionalGlobalPos(buf, w); err != nil {
		return err
	}
	return w.CopyBool(buf) // tracked
}

func decodeFireworkExplosionWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyFireworkExplosion(buf, w)
}

func decodeProfileWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyOptionalString(buf, w); err != nil { // name
		return err
	}
	if err := copyOptionalUUID(buf, w); err != nil {
		return err
	}
	return copyVarIntPrefixedList(buf, w, copyGameProfileProperty)
}

func decodeBannerPatternsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, copyBannerPattern)
}

func decodePotDecorationsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	for range 4 {
		present, err := buf.ReadBool()
		if err != nil {
			return err
		}
		w.WriteBool(present)
		if present {
			if err := copyItemStackTemplate(buf, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func decodeBlockStateWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
		if err := w.CopyString(buf, maxStringLen); err != nil { // key
			return err
		}
		return w.CopyString(buf, maxStringLen) // value
	})
}

func decodeBeesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
		if err := copyNBT(buf, w); err != nil { // entity data
			return err
		}
		if err := w.CopyVarInt(buf); err != nil { // ticks in hive
			return err
		}
		return w.CopyVarInt(buf) // min ticks in hive
	})
}

// copyVarIntPrefixedList copies a VarInt count followed by that many elements using the provided copy function.
func copyVarIntPrefixedList(buf *ns.PacketBuffer, w *ns.PacketBuffer, copyFn func(*ns.PacketBuffer, *ns.PacketBuffer) error) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copyFn(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copySlot(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)

	if count <= 0 {
		return nil
	}

	itemID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(itemID)

	addCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(addCount)

	removeCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(removeCount)

	for range int(addCount) {
		compID, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(compID)

		codec := componentCodecs[int32(compID)]
		if codec == nil {
			return fmt.Errorf("unknown component %d in slot", compID)
		}
		data, err := codec.DecodeWire(buf)
		if err != nil {
			return err
		}
		if _, err := w.Write(data); err != nil {
			return err
		}
	}

	for range int(removeCount) {
		compID, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(compID)
	}

	return nil
}

func copyOptionalVarInt(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return w.CopyVarInt(buf)
	}
	return nil
}

func copyOptionalInt(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return w.CopyInt32(buf)
	}
	return nil
}

func copyOptionalString(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return w.CopyString(buf, maxStringLen)
	}
	return nil
}

func copyOptionalIdentifier(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyOptionalString(buf, w)
}

func copyOptionalNBT(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return copyNBT(buf, w)
	}
	return nil
}

func copyOptionalUUID(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return w.CopyUUID(buf)
	}
	return nil
}

func copyHolderSet(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		return w.CopyString(buf, maxStringLen)
	}
	for range int(typeID) - 1 {
		if err := w.CopyVarInt(buf); err != nil {
			return err
		}
	}
	return nil
}

func copyOptionalHolderSet(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return copyHolderSet(buf, w)
	}
	return nil
}

func copySoundEvent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copySoundEventHolder(buf, w)
}

func copyOptionalSoundEvent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return copySoundEvent(buf, w)
	}
	return nil
}

func copyConsumeEffect(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	switch typeID {
	case 0: // apply_effects
		if err := copyVarIntPrefixedList(buf, w, copyStatusEffect); err != nil {
			return err
		}
		return w.CopyFloat32(buf) // probability
	case 1: // remove_effects
		return copyHolderSet(buf, w)
	case 2: // clear_all_effects
		return nil
	case 3: // teleport_randomly
		return w.CopyFloat32(buf) // diameter
	case 4: // play_sound
		return copySoundEvent(buf, w)
	}
	return nil
}

func copyStatusEffect(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyVarInt(buf); err != nil {
		return err
	}
	return copyStatusEffectDetails(buf, w)
}

func copyStatusEffectDetails(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// amplifier, duration
	if err := w.CopyVarInt(buf); err != nil {
		return err
	}
	if err := w.CopyVarInt(buf); err != nil {
		return err
	}
	// ambient, show particles, show icon
	for range 3 {
		if err := w.CopyBool(buf); err != nil {
			return err
		}
	}
	// hidden effect (recursive optional)
	hasHidden, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasHidden)
	if hasHidden {
		return copyStatusEffectDetails(buf, w)
	}
	return nil
}

func copyToolRule(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyHolderSet(buf, w); err != nil {
		return err
	}
	// optional speed
	hasSpeed, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasSpeed)
	if hasSpeed {
		if err := w.CopyFloat32(buf); err != nil {
			return err
		}
	}
	// optional correct for drops
	hasCorrect, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasCorrect)
	if hasCorrect {
		return w.CopyBool(buf)
	}
	return nil
}

func copyBlockPredicate(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional blocks
	if err := copyOptionalHolderSet(buf, w); err != nil {
		return err
	}
	// optional properties
	propCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(propCount)
	for range int(propCount) {
		if err := copyPropertyMatcher(buf, w); err != nil {
			return err
		}
	}
	// optional NBT
	if err := copyOptionalNBT(buf, w); err != nil {
		return err
	}
	return nil
}

func copyPropertyMatcher(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyString(buf, maxStringLen); err != nil { // property name
		return err
	}
	isExact, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(isExact)
	if isExact {
		return w.CopyString(buf, maxStringLen) // value
	}
	// range: min, max
	if err := copyOptionalString(buf, w); err != nil {
		return err
	}
	return copyOptionalString(buf, w)
}

func copyOptionalKineticConditions(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		if err := w.CopyVarInt(buf); err != nil { // max duration ticks
			return err
		}
		if err := w.CopyFloat32(buf); err != nil { // min speed
			return err
		}
		return w.CopyFloat32(buf) // min relative speed
	}
	return nil
}

func copyTrimMaterial(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		if err := w.CopyString(buf, maxStringLen); err != nil { // asset name
			return err
		}
		if err := w.CopyVarInt(buf); err != nil { // ingredient
			return err
		}
		if err := w.CopyFloat32(buf); err != nil { // item model index
			return err
		}
		if err := copyVarIntPrefixedList(buf, w, func(buf, w *ns.PacketBuffer) error {
			if err := w.CopyVarInt(buf); err != nil { // armor type
				return err
			}
			return w.CopyString(buf, maxStringLen) // asset name
		}); err != nil {
			return err
		}
		return copyNBT(buf, w)
	}
	return nil
}

func copyTrimPattern(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		if err := w.CopyString(buf, maxStringLen); err != nil { // asset ID
			return err
		}
		if err := w.CopyVarInt(buf); err != nil { // template item
			return err
		}
		if err := copyNBT(buf, w); err != nil {
			return err
		}
		return w.CopyBool(buf) // decal
	}
	return nil
}

func copyBannerPattern(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		if err := w.CopyString(buf, maxStringLen); err != nil { // asset ID
			return err
		}
		if err := w.CopyString(buf, maxStringLen); err != nil { // translation key
			return err
		}
	}

	return w.CopyVarInt(buf) // color
}

func copyGameProfileProperty(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyString(buf, maxStringLen); err != nil { // name
		return err
	}
	if err := w.CopyString(buf, maxStringLen); err != nil { // value
		return err
	}
	return copyOptionalString(buf, w) // signature
}

func copyOptionalGlobalPos(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		if err := w.CopyString(buf, maxStringLen); err != nil { // dimension
			return err
		}
		return w.CopyPosition(buf) // position
	}
	return nil
}

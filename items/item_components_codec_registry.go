package items

// Component codec registration - registers complex component codecs at init time.
// Simple codecs (varint, float32, string, empty) are auto-generated in item_components_codec_gen.go.

import (
	"fmt"
	"slices"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func init() {
	// struct components with dedicated codecs
	registerCodec(ComponentCustomName, textComponentCodec{
		get: func(c *Components) *ItemNameComponent { return c.CustomName },
		set: func(c *Components, v *ItemNameComponent) { c.CustomName = v },
	})
	registerCodec(ComponentItemName, textComponentCodec{
		get: func(c *Components) *ItemNameComponent { return c.ItemName },
		set: func(c *Components, v *ItemNameComponent) { c.ItemName = v },
	})
	registerCodec(ComponentAttributeModifiers, &attributeModifiersCodec{})
	registerCodec(ComponentRarity, &rarityCodec{})
	registerCodec(ComponentBreakSound, newHolderStringCodec(ComponentBreakSound, "minecraft:sound_event", func(c *Components) string { return c.BreakSound }, func(c *Components, value string) { c.BreakSound = value }, copySoundEventDirect))
	registerCodec(ComponentInstrument, newHolderStringCodec(ComponentInstrument, "minecraft:instrument", func(c *Components) string { return c.Instrument }, func(c *Components, value string) { c.Instrument = value }, copyInstrumentDirect))
	registerCodec(ComponentJukeboxPlayable, newHolderStringCodec(ComponentJukeboxPlayable, "minecraft:jukebox_song", func(c *Components) string { return c.JukeboxPlayable }, func(c *Components, value string) { c.JukeboxPlayable = value }, copyJukeboxSongDirect))
	registerCodec(ComponentProvidesTrimMaterial, newHolderStringCodec(ComponentProvidesTrimMaterial, "minecraft:trim_material", func(c *Components) string { return c.ProvidesTrimMaterial }, func(c *Components, value string) { c.ProvidesTrimMaterial = value }, copyTrimMaterialDirect))

	// lore, enchantments, and tool codecs (full implementations)
	registerCodec(ComponentLore, &loreCodec{})
	registerCodec(ComponentEnchantments, &enchantmentsCodec{
		get: func(c *Components) map[string]int32 { return c.Enchantments },
		set: func(c *Components, v map[string]int32) { c.Enchantments = v },
	})
	registerCodec(ComponentStoredEnchantments, &enchantmentsCodec{
		get: func(c *Components) map[string]int32 { return c.StoredEnchantments },
		set: func(c *Components, v map[string]int32) { c.StoredEnchantments = v },
	})
	registerCodec(ComponentTool, &toolCodec{})

	// typed optional components
	registerCodec(ComponentUseEffects, pointerComponentCodec[UseEffects]{
		get:  func(c *Components) *UseEffects { return c.UseEffects },
		set:  func(c *Components, v *UseEffects) { c.UseEffects = v },
		read: readUseEffects, write: writeUseEffects, equal: equalUseEffects,
	})
	registerCodec(ComponentAttackRange, pointerComponentCodec[AttackRange]{
		get:  func(c *Components) *AttackRange { return c.AttackRange },
		set:  func(c *Components, v *AttackRange) { c.AttackRange = v },
		read: readAttackRange, write: writeAttackRange, equal: equalValue[AttackRange],
	})
	registerCodec(ComponentEquippable, pointerComponentCodec[Equippable]{
		get:  func(c *Components) *Equippable { return c.Equippable },
		set:  func(c *Components, v *Equippable) { c.Equippable = v },
		read: readEquippable, write: writeEquippable, equal: deepEqual[Equippable],
	})
	registerCodec(ComponentBlocksAttacks, pointerComponentCodec[BlocksAttacks]{
		get:  func(c *Components) *BlocksAttacks { return c.BlocksAttacks },
		set:  func(c *Components, v *BlocksAttacks) { c.BlocksAttacks = v },
		read: readBlocksAttacks, write: writeBlocksAttacks, equal: deepEqual[BlocksAttacks],
	})
	registerCodec(ComponentDye, pointerComponentCodec[int32]{
		get:  func(c *Components) *int32 { return c.Dye },
		set:  func(c *Components, v *int32) { c.Dye = v },
		read: readDyeColor, write: writeDyeColor, equal: equalValue[int32],
	})
	registerCodec(ComponentBaseColor, pointerComponentCodec[int32]{
		get:  func(c *Components) *int32 { return c.BaseColor },
		set:  func(c *Components, v *int32) { c.BaseColor = v },
		read: readDyeColor, write: writeDyeColor, equal: equalValue[int32],
	})
	registerCodec(ComponentDyedColor, pointerComponentCodec[int32]{
		get:  func(c *Components) *int32 { return c.DyedColor },
		set:  func(c *Components, v *int32) { c.DyedColor = v },
		read: readInt32Pointer, write: writeInt32Pointer, equal: equalValue[int32],
	})

	// complex passthrough codecs - these have custom decoders
	// simple passthroughs (varint, bool, string, empty, int32, nbt, holderSet, slot, slotList)
	// are registered in item_components_codec_gen.go
	registerPassthrough(ComponentCanBreak, decodeBlockPredicatesWire)
	registerPassthrough(ComponentCanPlaceOn, decodeBlockPredicatesWire)
	registerPassthrough(ComponentCustomModelData, decodeCustomModelDataWire)
	registerPassthrough(ComponentBundleContents, decodeTemplateListWire)
	registerPassthrough(ComponentChargedProjectiles, decodeTemplateListWire)
	registerPassthrough(ComponentContainer, decodeOptionalTemplateListWire)
	registerPassthrough(ComponentSulfurCubeContent, copyItemStackTemplate)
	registerCodec(ComponentEntityData, entityDataCodec{})
	registerPassthrough(ComponentBlockEntityData, decodeTypedEntityDataWire)
	registerPassthrough(ComponentUseRemainder, copyItemStackTemplate)
	registerCodec(ComponentConsumable, consumableCodec{})
	registerCodec(ComponentDeathProtection, deathProtectionCodec{})
	registerPassthrough(ComponentKineticWeapon, decodeKineticWeaponWire)
	registerPassthrough(ComponentPiercingWeapon, decodePiercingWeaponWire)
	registerPassthrough(ComponentPotionContents, decodePotionContentsWire)
	registerPassthrough(ComponentSuspiciousStewEffects, decodeSuspiciousStewWire)
	registerPassthrough(ComponentWritableBookContent, decodeWritableBookWire)
	registerPassthrough(ComponentWrittenBookContent, decodeWrittenBookWire)
	registerCodec(ComponentTrim, trimCodec{})
	registerCodec(ComponentPaintingVariant, paintingVariantCodec{})
	registerCodec(ComponentBannerPatterns, bannerPatternsCodec{})
	registerCodec(ComponentPotDecorations, potDecorationsCodec{})
	registerCodec(ComponentFireworkExplosion, fireworkExplosionCodec{})
	registerCodec(ComponentFireworks, fireworksCodec{})
	registerPassthrough(ComponentRecipes, decodeNBTWire)
	registerPassthrough(ComponentLodestoneTracker, decodeLodestoneWire)
	registerPassthrough(ComponentProfile, CopyResolvableProfile)
	registerPassthrough(ComponentBlockState, decodeBlockStateWire)
	registerPassthrough(ComponentBees, decodeBeesWire)
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

// entries returns the registry codec.registry names: the built-in sound event
// registry or a synchronized one.
func (codec *holderStringCodec) entries() *registries.Registry {
	if codec.registry == "minecraft:sound_event" {
		return registries.SoundEvent
	}
	return registries.SynchronizedEntries[codec.registry]
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
		components.rawPassthrough[codec.id] = slices.Clone(data)
		codec.set(components, "")
		return nil
	}
	var name string
	if codec.plain {
		name = codec.entries().ByID(int32(holderID))
	} else {
		name = codec.entries().HolderName(int32(holderID))
	}
	if name == "" {
		return fmt.Errorf("unknown %s holder %d", codec.registry, holderID)
	}
	codec.set(components, name)
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
		return slices.Clone(direct), nil
	}
	value := codec.get(components)
	var entryID int32
	if codec.plain {
		entryID = codec.entries().Get(value)
	} else {
		entryID = codec.entries().HolderID(value)
	}
	if entryID < 0 {
		return nil, fmt.Errorf("unknown %s entry %q", codec.registry, value)
	}
	writer := ns.NewWriter()
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

// copySoundHolder validates and copies a vanilla Holder<SoundEvent>.
func copySoundHolder(reader, writer *ns.PacketBuffer) error {
	var sound ns.SoundEvent
	if err := sound.Decode(reader); err != nil {
		return err
	}
	return sound.Encode(writer)
}

func copyInstrumentDirect(reader, writer *ns.PacketBuffer) error {
	if err := copySoundHolder(reader, writer); err != nil {
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
	if err := copySoundHolder(reader, writer); err != nil {
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
	return components.Fireworks.FlightDuration != defaults.Fireworks.FlightDuration ||
		!slices.EqualFunc(components.FireworkExplosions, defaults.FireworkExplosions, equalFireworkExplosion), true
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
	return a.Shape == b.Shape && a.HasTrail == b.HasTrail && a.HasTwinkle == b.HasTwinkle &&
		slices.Equal(a.Colors, b.Colors) && slices.Equal(a.FadeColors, b.FadeColors)
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
	for _, side := range []**Stack{&decorations.Back, &decorations.Left, &decorations.Right, &decorations.Front} {
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
	for _, side := range []*Stack{
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
	patterns := registries.SynchronizedEntries["minecraft:banner_pattern"]
	layers := make([]BannerPatternLayer, count)
	for i := range layers {
		holder, err := reader.ReadVarInt()
		if err != nil || holder < 0 || holder > ns.VarInt(patterns.Size()) {
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
			layers[i].Pattern = patterns.HolderName(int32(holder))
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
	for _, layer := range *components.BannerPatterns {
		pattern := registries.SynchronizedEntries["minecraft:banner_pattern"].HolderID(layer.Pattern)
		if layer.Color < 0 || layer.Color >= 16 {
			return nil, fmt.Errorf("invalid banner color %d", layer.Color)
		}
		if pattern >= 0 {
			if err := writer.WriteVarInt(ns.VarInt(pattern)); err != nil {
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
		return a == b
	}
	return slices.Equal(*a, *b)
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
	name := registries.SynchronizedEntries["minecraft:painting_variant"].ByID(int32(id))
	if err != nil || name == "" {
		return fmt.Errorf("unknown painting variant id %d", id)
	}
	components.PaintingVariant = name
	return nil
}

func (paintingVariantCodec) Clear(components *Components) { components.PaintingVariant = "" }

func (paintingVariantCodec) Differs(components, defaults *Components) (bool, bool) {
	return components.PaintingVariant != defaults.PaintingVariant, components.PaintingVariant != ""
}

func (paintingVariantCodec) Encode(components *Components) ([]byte, error) {
	id := registries.SynchronizedEntries["minecraft:painting_variant"].Get(components.PaintingVariant)
	if id < 0 {
		return nil, fmt.Errorf("unknown painting variant %q", components.PaintingVariant)
	}
	writer := ns.NewWriter()
	if err := writer.WriteVarInt(ns.VarInt(id)); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
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
	if err := copySoundHolder(buf, w); err != nil {
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

func decodeHolderSetWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyHolderSet(buf, w)
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

func decodeLodestoneWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyOptionalGlobalPos(buf, w); err != nil {
		return err
	}
	return w.CopyBool(buf) // tracked
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
	if count < 0 {
		return fmt.Errorf("negative list count %d", count)
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
	if addCount < 0 || removeCount < 0 {
		return fmt.Errorf("invalid slot patch size %d/%d", addCount, removeCount)
	}

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
	if typeID < 0 {
		return fmt.Errorf("invalid holder set marker %d", typeID)
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

func copyOptionalSoundEvent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return copySoundHolder(buf, w)
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
		return copySoundHolder(buf, w)
	}
	return fmt.Errorf("unknown consume effect type %d", typeID)
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

// CopyResolvableProfile validates and copies a vanilla ResolvableProfile
// (the minecraft:profile component and the resolvable_profile entity
// metadata serializer): either a complete GameProfile or a partial profile,
// followed by the PlayerSkin.Patch overrides.
func CopyResolvableProfile(buf, w *ns.PacketBuffer) error {
	complete, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(complete)
	if complete {
		if err := w.CopyUUID(buf); err != nil {
			return err
		}
		if err := w.CopyString(buf, maxPlayerNameLen); err != nil {
			return err
		}
	} else {
		if err := copyOptional(buf, w, func(buf, w *ns.PacketBuffer) error { return w.CopyString(buf, maxPlayerNameLen) }); err != nil {
			return err
		}
		if err := copyOptionalUUID(buf, w); err != nil {
			return err
		}
	}
	if err := copyGameProfileProperties(buf, w); err != nil {
		return err
	}
	// skin patch: body, cape and elytra textures, then the player model type
	for range 3 {
		if err := copyOptional(buf, w, func(buf, w *ns.PacketBuffer) error { return w.CopyString(buf, maxStringLen) }); err != nil {
			return err
		}
	}
	return copyOptional(buf, w, func(buf, w *ns.PacketBuffer) error { return w.CopyBool(buf) })
}

// Vanilla GameProfile limits (ByteBufCodecs.PLAYER_NAME, GAME_PROFILE_PROPERTIES).
const (
	maxPlayerNameLen          = 16
	maxProfileProperties      = 16
	maxProfilePropertyNameLen = 64
	maxProfileSignatureLen    = 1024
)

func copyGameProfileProperties(buf, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if count < 0 || count > maxProfileProperties {
		return fmt.Errorf("profile property count %d outside 0..%d", count, maxProfileProperties)
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := w.CopyString(buf, maxProfilePropertyNameLen); err != nil {
			return err
		}
		if err := w.CopyString(buf, maxStringLen); err != nil {
			return err
		}
		if err := copyOptional(buf, w, func(buf, w *ns.PacketBuffer) error { return w.CopyString(buf, maxProfileSignatureLen) }); err != nil {
			return err
		}
	}
	return nil
}

// copyOptional copies a boolean presence flag and, when set, the value.
func copyOptional(buf, w *ns.PacketBuffer, copyValue func(buf, w *ns.PacketBuffer) error) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if !present {
		return nil
	}
	return copyValue(buf, w)
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

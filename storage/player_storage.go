package storage

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
)

// PlayerData holds the data persisted for a player.
type PlayerData struct {
	X, Y, Z          float64
	Yaw, Pitch       float32
	Dimension        string
	Gamemode         int32
	Inventory        []InventorySlot
	Equipment        []EquipmentSlot
	HeldSlot         int32
	Health           float32
	Food             int32
	Saturation       float32
	FoodExhaustion   float32
	FoodTickTimer    int32
	XpLevel          int32
	XpProgress       float32
	XpTotal          int32
	Tags             []string
	FireTicks        int32
	AirSupply        int32
	TicksFrozen      int32
	FallDistance     float64
	AbsorptionAmount float32
	FallFlying       bool
	Effects          []EffectData
	RecipeBook       RecipeBookData
	HasRespawn       bool
	RespawnDimension string
	RespawnX         int32
	RespawnY         int32
	RespawnZ         int32
	RespawnYaw       float32
	RespawnPitch     float32
	RespawnForced    bool
}

type EffectData struct {
	Name          string
	Duration      int32
	Amplifier     int32
	Ambient       bool
	ShowParticles bool
	ShowIcon      bool
	Hidden        *EffectData
}

type RecipeBookData struct {
	CraftingOpen          bool
	CraftingFiltering     bool
	FurnaceOpen           bool
	FurnaceFiltering      bool
	BlastFurnaceOpen      bool
	BlastFurnaceFiltering bool
	SmokerOpen            bool
	SmokerFiltering       bool
	Recipes               []string
	ToBeDisplayed         []string
}

// InventorySlot represents one item in a player's inventory.
type InventorySlot struct {
	Slot       int8
	ID         string // e.g. "minecraft:stone"
	Count      int32
	Components nbt.Compound
}

// EquipmentSlot represents one entry in the 26.2 equipment compound.
type EquipmentSlot struct {
	Slot       string
	ID         string
	Count      int32
	Components nbt.Compound
}

// SavePlayer writes player data to <dir>/<uuid>.dat (gzip-compressed NBT).
func SavePlayer(dir, uuid string, pd *PlayerData) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	root := nbt.Compound{
		"DataVersion":         nbt.Int(DataVersion),
		"playerGameType":      nbt.Int(pd.Gamemode),
		"Dimension":           nbt.String(pd.Dimension),
		"SelectedItemSlot":    nbt.Int(pd.HeldSlot),
		"Health":              nbt.Float(pd.Health),
		"foodLevel":           nbt.Int(pd.Food),
		"foodSaturationLevel": nbt.Float(pd.Saturation),
		"foodExhaustionLevel": nbt.Float(pd.FoodExhaustion),
		"foodTickTimer":       nbt.Int(pd.FoodTickTimer),
		"XpLevel":             nbt.Int(pd.XpLevel),
		"XpP":                 nbt.Float(pd.XpProgress),
		"XpTotal":             nbt.Int(pd.XpTotal),
		"Fire":                nbt.Short(pd.FireTicks),
		"Air":                 nbt.Short(pd.AirSupply),
		"TicksFrozen":         nbt.Int(pd.TicksFrozen),
		"fall_distance":       nbt.Double(pd.FallDistance),
		"AbsorptionAmount":    nbt.Float(pd.AbsorptionAmount),
		"FallFlying":          nbt.Byte(boolByte(pd.FallFlying)),
		"Pos": nbt.List{
			ElementType: nbt.TagDouble,
			Elements:    []nbt.Tag{nbt.Double(pd.X), nbt.Double(pd.Y), nbt.Double(pd.Z)},
		},
		"Rotation": nbt.List{
			ElementType: nbt.TagFloat,
			Elements:    []nbt.Tag{nbt.Float(pd.Yaw), nbt.Float(pd.Pitch)},
		},
	}

	root["Inventory"] = nbt.List{ElementType: nbt.TagCompound, Elements: encodeInventory(pd.Inventory)}
	if equipment := encodeEquipment(pd.Equipment); len(equipment) > 0 {
		root["equipment"] = equipment
	}
	tags := nbt.List{ElementType: nbt.TagString}
	for _, tag := range pd.Tags {
		tags.Elements = append(tags.Elements, nbt.String(tag))
	}
	root["Tags"] = tags
	if len(pd.Effects) > 0 {
		effects := nbt.List{ElementType: nbt.TagCompound}
		for _, effect := range pd.Effects {
			effects.Elements = append(effects.Elements, encodeEffect(effect))
		}
		root["active_effects"] = effects
	}
	root["recipeBook"] = encodeRecipeBook(pd.RecipeBook)
	if pd.HasRespawn {
		respawn := nbt.Compound{
			"dimension": nbt.String(pd.RespawnDimension),
			"pos":       nbt.IntArray{pd.RespawnX, pd.RespawnY, pd.RespawnZ},
			"yaw":       nbt.Float(pd.RespawnYaw),
			"pitch":     nbt.Float(pd.RespawnPitch),
		}
		if pd.RespawnForced {
			respawn["forced"] = nbt.Byte(1)
		}
		root["respawn"] = respawn
	}

	data, err := nbt.EncodeFile(root, "")
	if err != nil {
		return fmt.Errorf("encode NBT: %w", err)
	}

	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(data); err != nil {
		return fmt.Errorf("gzip write: %w", err)
	}
	if err := gw.Close(); err != nil {
		return fmt.Errorf("gzip close: %w", err)
	}

	path := filepath.Join(dir, uuid+".dat")
	return os.WriteFile(path, buf.Bytes(), 0644)
}

// LoadPlayer reads player data from <dir>/<uuid>.dat. Returns nil, nil if file doesn't exist.
func LoadPlayer(dir, uuid string) (*PlayerData, error) {
	path := filepath.Join(dir, uuid+".dat")
	compressed, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	gr, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, fmt.Errorf("gzip reader: %w", err)
	}
	defer gr.Close()

	data, err := io.ReadAll(gr)
	if err != nil {
		return nil, fmt.Errorf("gzip read: %w", err)
	}

	tag, _, err := nbt.DecodeFile(data, nbt.WithMaxBytes(0))
	if err != nil {
		return nil, fmt.Errorf("decode NBT: %w", err)
	}

	root, ok := tag.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("root is not compound")
	}

	pd := &PlayerData{
		Gamemode:         root.GetInt("playerGameType"),
		Dimension:        root.GetString("Dimension"),
		HeldSlot:         root.GetInt("SelectedItemSlot"),
		Health:           root.GetFloat("Health"),
		Food:             root.GetInt("foodLevel"),
		Saturation:       root.GetFloat("foodSaturationLevel"),
		FoodExhaustion:   root.GetFloat("foodExhaustionLevel"),
		FoodTickTimer:    root.GetInt("foodTickTimer"),
		XpLevel:          root.GetInt("XpLevel"),
		XpProgress:       root.GetFloat("XpP"),
		XpTotal:          root.GetInt("XpTotal"),
		FireTicks:        int32(root.GetShort("Fire")),
		AirSupply:        int32(root.GetShort("Air")),
		TicksFrozen:      root.GetInt("TicksFrozen"),
		FallDistance:     root.GetDouble("fall_distance"),
		AbsorptionAmount: root.GetFloat("AbsorptionAmount"),
		FallFlying:       root.GetByte("FallFlying") != 0,
	}

	// position
	posList := root.GetList("Pos")
	if posList.Len() >= 3 {
		if d, ok := posList.Get(0).(nbt.Double); ok {
			pd.X = float64(d)
		}
		if d, ok := posList.Get(1).(nbt.Double); ok {
			pd.Y = float64(d)
		}
		if d, ok := posList.Get(2).(nbt.Double); ok {
			pd.Z = float64(d)
		}
	}

	// rotation
	rotList := root.GetList("Rotation")
	if rotList.Len() >= 2 {
		if f, ok := rotList.Get(0).(nbt.Float); ok {
			pd.Yaw = float32(f)
		}
		if f, ok := rotList.Get(1).(nbt.Float); ok {
			pd.Pitch = float32(f)
		}
	}
	if respawn := root.GetCompound("respawn"); respawn != nil {
		if position := respawn.GetIntArray("pos"); len(position) == 3 {
			pd.HasRespawn = true
			pd.RespawnDimension = respawn.GetString("dimension")
			pd.RespawnX, pd.RespawnY, pd.RespawnZ = position[0], position[1], position[2]
			pd.RespawnYaw, pd.RespawnPitch = respawn.GetFloat("yaw"), respawn.GetFloat("pitch")
			pd.RespawnForced = respawn.GetByte("forced") != 0
		}
	}

	pd.Inventory = decodeInventory(root.GetList("Inventory"))
	pd.Equipment = decodeEquipment(root.GetCompound("equipment"))
	for _, tag := range root.GetList("Tags").Elements {
		if name, ok := tag.(nbt.String); ok {
			pd.Tags = append(pd.Tags, string(name))
		}
	}
	for _, tag := range root.GetList("active_effects").Elements {
		if effect, ok := tag.(nbt.Compound); ok {
			pd.Effects = append(pd.Effects, decodeEffect(effect))
		}
	}
	pd.RecipeBook = decodeRecipeBook(root.GetCompound("recipeBook"))

	return pd, nil
}

func encodeEffect(effect EffectData) nbt.Compound {
	return encodeEffectDetails(effect, true)
}

func encodeEffectDetails(effect EffectData, includeID bool) nbt.Compound {
	result := nbt.Compound{
		"duration": nbt.Int(effect.Duration), "amplifier": nbt.Byte(effect.Amplifier),
		"ambient": nbt.Byte(boolByte(effect.Ambient)), "show_particles": nbt.Byte(boolByte(effect.ShowParticles)), "show_icon": nbt.Byte(boolByte(effect.ShowIcon)),
	}
	if includeID {
		result["id"] = nbt.String(effect.Name)
	}
	if effect.Hidden != nil {
		result["hidden_effect"] = encodeEffectDetails(*effect.Hidden, false)
	}
	return result
}

func decodeEffect(effect nbt.Compound) EffectData {
	return decodeEffectDetails(effect, effect.GetString("id"))
}

func decodeEffectDetails(effect nbt.Compound, name string) EffectData {
	showParticles := true
	if effect.Get("show_particles") != nil {
		showParticles = effect.GetByte("show_particles") != 0
	}
	showIcon := showParticles
	if effect.Get("show_icon") != nil {
		showIcon = effect.GetByte("show_icon") != 0
	}
	result := EffectData{
		Name: name, Duration: effect.GetInt("duration"), Amplifier: int32(effect.GetByte("amplifier")),
		Ambient: effect.GetByte("ambient") != 0, ShowParticles: showParticles, ShowIcon: showIcon,
	}
	if hidden := effect.GetCompound("hidden_effect"); len(hidden) > 0 {
		value := decodeEffectDetails(hidden, name)
		result.Hidden = &value
	}
	return result
}

// encodeInventory converts inventory slots to nbt compounds, skipping invalid entries.
func encodeInventory(inv []InventorySlot) []nbt.Tag {
	var elements []nbt.Tag
	for _, slot := range inv {
		if slot.Slot < 0 || slot.Slot >= 36 {
			continue
		}
		comp := encodeItem(slot.ID, slot.Count, slot.Components)
		if comp == nil {
			continue
		}
		comp["Slot"] = nbt.Byte(slot.Slot)
		elements = append(elements, comp)
	}
	return elements
}

func encodeEquipment(equipment []EquipmentSlot) nbt.Compound {
	result := nbt.Compound{}
	for _, slot := range equipment {
		if !validEquipmentSlot(slot.Slot) {
			continue
		}
		if item := encodeItem(slot.ID, slot.Count, slot.Components); item != nil {
			result[slot.Slot] = item
		}
	}
	return result
}

func encodeItem(identifier string, count int32, components nbt.Compound) nbt.Compound {
	if registries.Item.Get(identifier) < 0 || count < 1 || count > 99 {
		return nil
	}
	result := nbt.Compound{"id": nbt.String(identifier), "count": nbt.Int(count)}
	if len(components) > 0 {
		result["components"] = components
	}
	return result
}

// decodeInventory converts an NBT inventory list back to inventory slots.
func decodeInventory(invList nbt.List) []InventorySlot {
	var inv []InventorySlot
	for _, elem := range invList.Elements {
		slotComp, ok := elem.(nbt.Compound)
		if !ok {
			continue
		}
		slot := slotComp.GetByte("Slot")
		if slot < 0 || slot >= 36 {
			continue
		}
		inv = append(inv, InventorySlot{Slot: slot, ID: slotComp.GetString("id"), Count: slotComp.GetInt("count"), Components: slotComp.GetCompound("components")})
	}
	return inv
}

func decodeEquipment(compound nbt.Compound) []EquipmentSlot {
	var equipment []EquipmentSlot
	for slot, tag := range compound {
		item, ok := tag.(nbt.Compound)
		if !ok || !validEquipmentSlot(slot) {
			continue
		}
		equipment = append(equipment, EquipmentSlot{Slot: slot, ID: item.GetString("id"), Count: item.GetInt("count"), Components: item.GetCompound("components")})
	}
	return equipment
}

func encodeRecipeBook(book RecipeBookData) nbt.Compound {
	result := nbt.Compound{
		"isGuiOpen":                           nbt.Byte(boolByte(book.CraftingOpen)),
		"isFilteringCraftable":                nbt.Byte(boolByte(book.CraftingFiltering)),
		"isFurnaceGuiOpen":                    nbt.Byte(boolByte(book.FurnaceOpen)),
		"isFurnaceFilteringCraftable":         nbt.Byte(boolByte(book.FurnaceFiltering)),
		"isBlastingFurnaceGuiOpen":            nbt.Byte(boolByte(book.BlastFurnaceOpen)),
		"isBlastingFurnaceFilteringCraftable": nbt.Byte(boolByte(book.BlastFurnaceFiltering)),
		"isSmokerGuiOpen":                     nbt.Byte(boolByte(book.SmokerOpen)),
		"isSmokerFilteringCraftable":          nbt.Byte(boolByte(book.SmokerFiltering)),
	}
	result["recipes"] = stringList(book.Recipes)
	result["toBeDisplayed"] = stringList(book.ToBeDisplayed)
	return result
}

func decodeRecipeBook(compound nbt.Compound) RecipeBookData {
	return RecipeBookData{
		CraftingOpen:          compound.GetByte("isGuiOpen") != 0,
		CraftingFiltering:     compound.GetByte("isFilteringCraftable") != 0,
		FurnaceOpen:           compound.GetByte("isFurnaceGuiOpen") != 0,
		FurnaceFiltering:      compound.GetByte("isFurnaceFilteringCraftable") != 0,
		BlastFurnaceOpen:      compound.GetByte("isBlastingFurnaceGuiOpen") != 0,
		BlastFurnaceFiltering: compound.GetByte("isBlastingFurnaceFilteringCraftable") != 0,
		SmokerOpen:            compound.GetByte("isSmokerGuiOpen") != 0,
		SmokerFiltering:       compound.GetByte("isSmokerFilteringCraftable") != 0,
		Recipes:               readStringList(compound.GetList("recipes")),
		ToBeDisplayed:         readStringList(compound.GetList("toBeDisplayed")),
	}
}

func stringList(values []string) nbt.List {
	result := nbt.List{ElementType: nbt.TagString}
	for _, value := range values {
		result.Elements = append(result.Elements, nbt.String(value))
	}
	return result
}

func readStringList(list nbt.List) []string {
	result := make([]string, 0, len(list.Elements))
	for _, value := range list.Elements {
		if text, ok := value.(nbt.String); ok {
			result = append(result, string(text))
		}
	}
	return result
}

func boolByte(value bool) int8 {
	if value {
		return 1
	}
	return 0
}

func validEquipmentSlot(slot string) bool {
	switch slot {
	case "mainhand", "offhand", "feet", "legs", "chest", "head", "body", "saddle":
		return true
	default:
		return false
	}
}

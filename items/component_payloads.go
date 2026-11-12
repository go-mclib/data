package items

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// SwingAnimation is shared by attack_animation and interact_animation.
// Type is none (0), whack (1), or stab (2); Duration may be zero.
type SwingAnimation struct {
	Type     int32
	Duration int32
}

type VillagerFood struct{ Nutrition int32 }

// ResolvableInt is a constant unless Reference names a context-int provider.
type ResolvableInt struct {
	Value     int32
	Reference string
}

// ResolvableFloat is a constant unless Reference names a context-float provider.
type ResolvableFloat struct {
	Value     float32
	Reference string
}

type Compostable struct{ Layers ResolvableInt }
type CookingFuel struct {
	BurnTime        ResolvableInt
	SpeedMultiplier ResolvableFloat
}
type BrewingFuel struct {
	Uses            ResolvableInt
	SpeedMultiplier ResolvableFloat
}
type MobVisibility struct {
	TargetingEntityTypes []string
	Visibility           float32
}

// SignText holds exactly four lines. A nil FilteredMessages means absent.
// Text components retain their complete NBT representation in Raw.
type SignText struct {
	Messages         [4]ItemNameComponent
	FilteredMessages *[4]ItemNameComponent
	Color            int32
	HasGlowingText   bool
}

func cloneValue[T any](value *T) *T {
	if value == nil {
		return nil
	}
	clone := *value
	return &clone
}

func cloneSignText(value *SignText) *SignText {
	clone := cloneValue(value)
	if clone == nil {
		return nil
	}
	for i := range clone.Messages {
		clone.Messages[i].Raw = slices.Clone(clone.Messages[i].Raw)
	}
	clone.FilteredMessages = cloneValue(value.FilteredMessages)
	if clone.FilteredMessages != nil {
		for i := range clone.FilteredMessages {
			clone.FilteredMessages[i].Raw = slices.Clone(clone.FilteredMessages[i].Raw)
		}
	}
	return clone
}

func cloneDecoration(value *Stack) *Stack {
	clone := cloneValue(value)
	if clone != nil {
		clone.Components = value.Components.Clone()
	}
	return clone
}

func equalDecoration(a, b *Stack) bool {
	if a == nil || b == nil {
		return a == b
	}
	left, err := a.ToSlot()
	if err != nil {
		return false
	}
	right, err := b.ToSlot()
	return err == nil && SameStack(left, right)
}

// pointerComponentCodec shares the optional-value lifecycle; payload functions
// own the wire format and equality, including nested text and holder sets.
type pointerComponentCodec[T any] struct {
	get   func(*Components) *T
	set   func(*Components, *T)
	read  func(*ns.PacketBuffer) (*T, error)
	write func(*ns.PacketBuffer, *T) error
	equal func(*T, *T) bool
}

func (codec pointerComponentCodec[T]) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	value, err := codec.read(buf)
	if err != nil {
		return nil, err
	}
	writer := ns.NewWriter()
	if err := codec.write(writer, value); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}
func (codec pointerComponentCodec[T]) Apply(c *Components, data []byte) error {
	value, err := codec.read(ns.NewReader(data))
	if err == nil {
		codec.set(c, value)
	}
	return err
}
func (codec pointerComponentCodec[T]) Clear(c *Components) { codec.set(c, nil) }
func (codec pointerComponentCodec[T]) Differs(c, defaults *Components) (bool, bool) {
	value, base := codec.get(c), codec.get(defaults)
	if value == nil || base == nil {
		return value != nil || base != nil, value != nil
	}
	return !codec.equal(value, base), true
}
func (codec pointerComponentCodec[T]) Encode(c *Components) ([]byte, error) {
	value := codec.get(c)
	if value == nil {
		return nil, fmt.Errorf("component is absent")
	}
	writer := ns.NewWriter()
	if err := codec.write(writer, value); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

// equalValue reports whether two optional values are both absent or equal.
func equalValue[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

func deepEqual[T any](a, b *T) bool { return reflect.DeepEqual(a, b) }

func init() {
	registerCodec(ComponentAttackAnimation, pointerComponentCodec[SwingAnimation]{
		get:  func(c *Components) *SwingAnimation { return c.AttackAnimation },
		set:  func(c *Components, v *SwingAnimation) { c.AttackAnimation = v },
		read: readSwingAnimation, write: writeSwingAnimation, equal: equalValue[SwingAnimation],
	})
	registerCodec(ComponentInteractAnimation, pointerComponentCodec[SwingAnimation]{
		get:  func(c *Components) *SwingAnimation { return c.InteractAnimation },
		set:  func(c *Components, v *SwingAnimation) { c.InteractAnimation = v },
		read: readSwingAnimation, write: writeSwingAnimation, equal: equalValue[SwingAnimation],
	})
	registerCodec(ComponentCompostable, pointerComponentCodec[Compostable]{
		get:  func(c *Components) *Compostable { return c.Compostable },
		set:  func(c *Components, v *Compostable) { c.Compostable = v },
		read: readCompostable, write: writeCompostable, equal: equalValue[Compostable],
	})
	registerCodec(ComponentCookingFuel, pointerComponentCodec[CookingFuel]{
		get:  func(c *Components) *CookingFuel { return c.CookingFuel },
		set:  func(c *Components, v *CookingFuel) { c.CookingFuel = v },
		read: readCookingFuel, write: writeCookingFuel, equal: equalValue[CookingFuel],
	})
	registerCodec(ComponentBrewingFuel, pointerComponentCodec[BrewingFuel]{
		get:  func(c *Components) *BrewingFuel { return c.BrewingFuel },
		set:  func(c *Components, v *BrewingFuel) { c.BrewingFuel = v },
		read: readBrewingFuel, write: writeBrewingFuel, equal: equalValue[BrewingFuel],
	})
	registerCodec(ComponentMobVisibility, pointerComponentCodec[MobVisibility]{
		get:  func(c *Components) *MobVisibility { return c.MobVisibility },
		set:  func(c *Components, v *MobVisibility) { c.MobVisibility = v },
		read: readMobVisibility, write: writeMobVisibility,
		equal: func(a, b *MobVisibility) bool {
			return a.Visibility == b.Visibility && slices.Equal(a.TargetingEntityTypes, b.TargetingEntityTypes)
		},
	})
	registerCodec(ComponentSignTextFront, pointerComponentCodec[SignText]{
		get:  func(c *Components) *SignText { return c.SignTextFront },
		set:  func(c *Components, v *SignText) { c.SignTextFront = v },
		read: readSignText, write: writeSignText, equal: equalSignText,
	})
	registerCodec(ComponentSignTextBack, pointerComponentCodec[SignText]{
		get:  func(c *Components) *SignText { return c.SignTextBack },
		set:  func(c *Components, v *SignText) { c.SignTextBack = v },
		read: readSignText, write: writeSignText, equal: equalSignText,
	})
	registerCodec(ComponentCushionColor, pointerComponentCodec[int32]{
		get:  func(c *Components) *int32 { return c.CushionColor },
		set:  func(c *Components, v *int32) { c.CushionColor = v },
		read: readDyeColor, write: writeDyeColor, equal: equalValue[int32],
	})
	registerCodec(ComponentBlockTransformer, &holderStringCodec{
		id: ComponentBlockTransformer, registry: "minecraft:block_transformer", plain: true,
		get: func(c *Components) string { return c.BlockTransformer },
		set: func(c *Components, v string) { c.BlockTransformer = v },
	})
	registerCodec(ComponentProvidesPotteryPattern, &holderStringCodec{
		id: ComponentProvidesPotteryPattern, registry: "minecraft:decorated_pot_pattern", plain: true,
		get: func(c *Components) string { return c.ProvidesPotteryPattern },
		set: func(c *Components, v string) { c.ProvidesPotteryPattern = v },
	})
}

func readSwingAnimation(reader *ns.PacketBuffer) (*SwingAnimation, error) {
	kind, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	duration, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if kind < 0 || kind > 2 || duration < 0 {
		return nil, fmt.Errorf("invalid swing animation %d/%d", kind, duration)
	}
	return &SwingAnimation{Type: int32(kind), Duration: int32(duration)}, nil
}
func writeSwingAnimation(writer *ns.PacketBuffer, value *SwingAnimation) error {
	if value.Type < 0 || value.Type > 2 || value.Duration < 0 {
		return fmt.Errorf("invalid swing animation %d/%d", value.Type, value.Duration)
	}
	writer.WriteVarInt(ns.VarInt(value.Type))
	return writer.WriteVarInt(ns.VarInt(value.Duration))
}
func readResolvableInt(reader *ns.PacketBuffer) (ResolvableInt, error) {
	constant, err := reader.ReadBool()
	if err != nil {
		return ResolvableInt{}, err
	}
	if constant {
		value, err := reader.ReadInt32()
		return ResolvableInt{Value: int32(value)}, err
	}
	reference, err := reader.ReadString(maxStringLen)
	return ResolvableInt{Reference: string(reference)}, err
}
func writeResolvableInt(writer *ns.PacketBuffer, value ResolvableInt) error {
	writer.WriteBool(ns.Boolean(value.Reference == ""))
	if value.Reference != "" {
		return writer.WriteString(ns.String(value.Reference))
	}
	return writer.WriteInt32(ns.Int32(value.Value))
}
func readResolvableFloat(reader *ns.PacketBuffer) (ResolvableFloat, error) {
	constant, err := reader.ReadBool()
	if err != nil {
		return ResolvableFloat{}, err
	}
	if constant {
		value, err := reader.ReadFloat32()
		return ResolvableFloat{Value: float32(value)}, err
	}
	reference, err := reader.ReadString(maxStringLen)
	return ResolvableFloat{Reference: string(reference)}, err
}
func writeResolvableFloat(writer *ns.PacketBuffer, value ResolvableFloat) error {
	writer.WriteBool(ns.Boolean(value.Reference == ""))
	if value.Reference != "" {
		return writer.WriteString(ns.String(value.Reference))
	}
	return writer.WriteFloat32(ns.Float32(value.Value))
}
func readCompostable(reader *ns.PacketBuffer) (*Compostable, error) {
	layers, err := readResolvableInt(reader)
	return &Compostable{Layers: layers}, err
}
func writeCompostable(writer *ns.PacketBuffer, value *Compostable) error {
	return writeResolvableInt(writer, value.Layers)
}
func readCookingFuel(reader *ns.PacketBuffer) (*CookingFuel, error) {
	burnTime, err := readResolvableInt(reader)
	if err != nil {
		return nil, err
	}
	speed, err := readResolvableFloat(reader)
	return &CookingFuel{BurnTime: burnTime, SpeedMultiplier: speed}, err
}
func writeCookingFuel(writer *ns.PacketBuffer, value *CookingFuel) error {
	if err := writeResolvableInt(writer, value.BurnTime); err != nil {
		return err
	}
	return writeResolvableFloat(writer, value.SpeedMultiplier)
}
func readBrewingFuel(reader *ns.PacketBuffer) (*BrewingFuel, error) {
	uses, err := readResolvableInt(reader)
	if err != nil {
		return nil, err
	}
	speed, err := readResolvableFloat(reader)
	return &BrewingFuel{Uses: uses, SpeedMultiplier: speed}, err
}
func writeBrewingFuel(writer *ns.PacketBuffer, value *BrewingFuel) error {
	if err := writeResolvableInt(writer, value.Uses); err != nil {
		return err
	}
	return writeResolvableFloat(writer, value.SpeedMultiplier)
}
func readMobVisibility(reader *ns.PacketBuffer) (*MobVisibility, error) {
	types, err := readHolderSet(reader, "entity type", registries.EntityType.ByID)
	if err != nil {
		return nil, err
	}
	visibility, err := reader.ReadFloat32()
	return &MobVisibility{TargetingEntityTypes: types, Visibility: float32(visibility)}, err
}
func writeMobVisibility(writer *ns.PacketBuffer, value *MobVisibility) error {
	if err := writeHolderSet(writer, "entity type", value.TargetingEntityTypes, registries.EntityType.Get); err != nil {
		return err
	}
	return writer.WriteFloat32(ns.Float32(value.Visibility))
}
func readDyeColor(reader *ns.PacketBuffer) (*int32, error) {
	value, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if value < 0 || value >= 16 {
		return nil, fmt.Errorf("invalid dye color %d", value)
	}
	return int32Pointer(int32(value)), nil
}
func writeDyeColor(writer *ns.PacketBuffer, value *int32) error {
	if *value < 0 || *value >= 16 {
		return fmt.Errorf("invalid dye color %d", *value)
	}
	return writer.WriteVarInt(ns.VarInt(*value))
}
func readInt32Pointer(reader *ns.PacketBuffer) (*int32, error) {
	value, err := reader.ReadInt32()
	if err != nil {
		return nil, err
	}
	return int32Pointer(int32(value)), nil
}
func writeInt32Pointer(writer *ns.PacketBuffer, value *int32) error {
	return writer.WriteInt32(ns.Int32(*value))
}
func readVarIntPointer(reader *ns.PacketBuffer) (*int32, error) {
	value, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return int32Pointer(int32(value)), nil
}
func writeVarIntPointer(writer *ns.PacketBuffer, value *int32) error {
	return writer.WriteVarInt(ns.VarInt(*value))
}
func readUseEffects(reader *ns.PacketBuffer) (*UseEffects, error) {
	canSprint, err := reader.ReadBool()
	if err != nil {
		return nil, err
	}
	vibrations, err := reader.ReadBool()
	if err != nil {
		return nil, err
	}
	speed, err := reader.ReadFloat32()
	if err != nil {
		return nil, err
	}
	return &UseEffects{CanSprint: bool(canSprint), InteractVibrations: bool(vibrations), SpeedMultiplier: float64(speed)}, nil
}

// equalUseEffects compares at wire precision: the speed multiplier is a float32 on the wire.
func equalUseEffects(a, b *UseEffects) bool {
	return a.CanSprint == b.CanSprint && a.InteractVibrations == b.InteractVibrations &&
		float32(a.SpeedMultiplier) == float32(b.SpeedMultiplier)
}
func writeUseEffects(writer *ns.PacketBuffer, value *UseEffects) error {
	writer.WriteBool(ns.Boolean(value.CanSprint))
	writer.WriteBool(ns.Boolean(value.InteractVibrations))
	return writer.WriteFloat32(ns.Float32(value.SpeedMultiplier))
}
func readSignText(reader *ns.PacketBuffer) (*SignText, error) {
	value := &SignText{}
	if err := readSignMessages(reader, &value.Messages); err != nil {
		return nil, err
	}
	filtered, err := reader.ReadBool()
	if err != nil {
		return nil, err
	}
	if filtered {
		value.FilteredMessages = new([4]ItemNameComponent)
		if err := readSignMessages(reader, value.FilteredMessages); err != nil {
			return nil, err
		}
	}
	color, err := readDyeColor(reader)
	if err != nil {
		return nil, err
	}
	value.Color = *color
	glowing, err := reader.ReadBool()
	value.HasGlowingText = bool(glowing)
	return value, err
}
func readSignMessages(reader *ns.PacketBuffer, messages *[4]ItemNameComponent) error {
	for i := range messages {
		tag, err := reader.ReadNBT()
		if err != nil {
			return err
		}
		message := textComponentFromNBT(tag)
		if message == nil {
			return fmt.Errorf("invalid sign message %d", i)
		}
		messages[i] = *message
	}
	return nil
}
func writeSignText(writer *ns.PacketBuffer, value *SignText) error {
	if err := writeSignMessages(writer, &value.Messages); err != nil {
		return err
	}
	writer.WriteBool(ns.Boolean(value.FilteredMessages != nil))
	if value.FilteredMessages != nil {
		if err := writeSignMessages(writer, value.FilteredMessages); err != nil {
			return err
		}
	}
	if err := writeDyeColor(writer, &value.Color); err != nil {
		return err
	}
	return writer.WriteBool(ns.Boolean(value.HasGlowingText))
}
func writeSignMessages(writer *ns.PacketBuffer, messages *[4]ItemNameComponent) error {
	for i := range messages {
		if messages[i].Raw != nil {
			if _, err := writer.Write(messages[i].Raw); err != nil {
				return err
			}
		} else if err := encodeItemName(writer, &messages[i]); err != nil {
			return err
		}
	}
	return nil
}
func equalSignText(a, b *SignText) bool {
	if a.Color != b.Color || a.HasGlowingText != b.HasGlowingText || (a.FilteredMessages == nil) != (b.FilteredMessages == nil) {
		return false
	}
	for i := range a.Messages {
		if !itemNameComponentsEqual(&a.Messages[i], &b.Messages[i]) {
			return false
		}
		if a.FilteredMessages != nil && !itemNameComponentsEqual(&a.FilteredMessages[i], &b.FilteredMessages[i]) {
			return false
		}
	}
	return true
}

package registries

// MobEffectColor returns the vanilla effect's RGB color, or 0 for an unknown effect.
func MobEffectColor(name string) int32 { return mobEffectColors[name] }

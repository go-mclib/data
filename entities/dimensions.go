package entities

// Dimensions returns the standing width, height, and eye height for an entity
// type. Returns zeros if the entity type is not found.
func Dimensions(name string) (width, height, eyeHeight float32) {
	d := dimensions[name]
	return d.Width, d.Height, d.EyeHeight
}

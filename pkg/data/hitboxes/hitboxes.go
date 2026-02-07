// Package hitboxes provides collision geometry types shared by block and entity hitbox packages.
package hitboxes

// AABB is an axis-aligned bounding box in block-local coordinates (typically 0.0–1.0).
type AABB struct {
	MinX, MinY, MinZ float64
	MaxX, MaxY, MaxZ float64
}

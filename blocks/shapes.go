package blocks

// AABB is an axis-aligned bounding box in block-local coordinates (typically 0.0–1.0).
type AABB struct {
	MinX, MinY, MinZ float64
	MaxX, MaxY, MaxZ float64
}

// CollisionShape returns the collision AABBs for a block state ID.
// Returns nil for invalid state IDs or blocks with no collision (air, etc.).
func CollisionShape(stateID int32) []AABB {
	if stateID < 0 || int(stateID) >= len(shapeByState) {
		return nil
	}
	return shapes[shapeByState[stateID]]
}

// HasCollision reports whether the given block state has any collision geometry.
func HasCollision(stateID int32) bool {
	if stateID < 0 || int(stateID) >= len(shapeByState) {
		return false
	}
	return shapeByState[stateID] != 0
}

// IsFullBlock reports whether the given block state is a full 1x1x1 collision block.
func IsFullBlock(stateID int32) bool {
	if stateID < 0 || int(stateID) >= len(shapeByState) {
		return false
	}
	return collisionFlags[stateID]&flagFullBlock != 0
}

// CollisionCoordinates returns the original voxel grid planes for axis 0/1/2
// (X/Y/Z), retaining internal planes relevant to clipping an initial overlap.
func CollisionCoordinates(stateID int32, axis int) []float64 {
	if stateID < 0 || int(stateID) >= len(shapeByState) || axis < 0 || axis > 2 {
		return nil
	}
	return shapeCoordinates[shapeByState[stateID]][axis]
}

// BlocksMotion reports membership in the vanilla blocks_motion tag.
// It is not interchangeable with collision geometry or fluid-flow blocking.
func BlocksMotion(stateID int32) bool {
	return stateID >= 0 && int(stateID) < len(collisionFlags) && collisionFlags[stateID]&flagBlocksMotion != 0
}

// BlocksFluidFlow reports membership in the vanilla blocks_fluid_flow tag.
func BlocksFluidFlow(stateID int32) bool {
	return stateID >= 0 && int(stateID) < len(collisionFlags) && collisionFlags[stateID]&flagBlocksFluidFlow != 0
}

// HasLargeCollisionShape reports whether collision geometry can leave its block cell.
func HasLargeCollisionShape(stateID int32) bool {
	return stateID >= 0 && int(stateID) < len(collisionFlags) && collisionFlags[stateID]&flagLargeCollisionShape != 0
}

// IsSuffocating is the vanilla suffocation predicate, not merely a solid shape.
func IsSuffocating(stateID int32) bool {
	return stateID >= 0 && int(stateID) < len(collisionFlags) && collisionFlags[stateID]&flagSuffocating != 0
}

// Vanilla direction IDs (Direction.get3DDataValue) for IsFaceSturdy.
const (
	DirectionDown = iota
	DirectionUp
	DirectionNorth
	DirectionSouth
	DirectionWest
	DirectionEast
)

// IsFaceSturdy reports whether the state's face in the given vanilla direction
// (DirectionDown … DirectionEast) is sturdy.
func IsFaceSturdy(stateID int32, direction int) bool {
	return stateID >= 0 && int(stateID) < len(sturdyFaces) && direction >= 0 && direction < 6 &&
		sturdyFaces[stateID]&(1<<direction) != 0
}

// FluidState returns the contained fluid: kind 0 empty, 1 water, 2 lava,
// amount in eighths, and the falling property. Waterlogged blocks are included.
func FluidState(stateID int32) (kind uint8, amount uint8, falling bool) {
	if stateID < 0 || int(stateID) >= len(fluidStates) {
		return 0, 0, false
	}
	state := fluidStates[stateID]
	return state & fluidKindMask, (state >> fluidAmountShift) & fluidAmountMask, state&fluidFalling != 0
}

package test

import "math"

type Vec2 struct {
	X, Y int
}

type Context struct{}

func ArrayContainsVec2(slice []Vec2, vec Vec2) bool {
	for _, v := range slice {
		if v.X == vec.X && v.Y == vec.Y {
			return true
		}
	}
	return false
}

func (ctx *Context) ComputeIntersectingCubes(origin, targetPos Vec2) []Vec2 {
	var result []Vec2

	dx := 1
	if targetPos.X < origin.X {
		dx = -1
	}
	dy := 1
	if targetPos.Y < origin.Y {
		dy = -1
	}

	var err int
	deltaX := int(math.Abs(float64(targetPos.X - origin.X)))
	deltaY := int(math.Abs(float64(targetPos.Y - origin.Y)))
	x, y := origin.X, origin.Y
	err = deltaX - deltaY

	for x != targetPos.X || y != targetPos.Y {
		result = append(result, Vec2{x, y})

		e2 := 2 * err
		if e2 > -deltaY {
			err -= deltaY
			x += dx
		}
		if e2 < deltaX {
			err += deltaX
			y += dy
		}
	}

	// Ensure targetPos is included if it's supposed to be part of the intersecting cubes
	if ArrayContainsVec2(result, targetPos) == false {
		result = append(result, targetPos)
	}

	return result
}

// Assume arrayContainsVec2 is implemented elsewhere as it's a test helper function

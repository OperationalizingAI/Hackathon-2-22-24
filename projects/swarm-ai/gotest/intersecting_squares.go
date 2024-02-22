package main

type Vec2 struct {
	X, Y int
}

// Context could potentially hold other relevant data for the computation,
// for the sake of this example, it's left empty.
type Context struct{}

// ComputeIntersectingCubes calculates the intersecting cubes (grid cells) between the origin and targetPos.
func (ctx *Context) ComputeIntersectingCubes(origin, targetPos Vec2) []Vec2 {
	var result []Vec2

	dx := abs(targetPos.X - origin.X)
	dy := -abs(targetPos.Y - origin.Y)
	err := dx + dy // error value e_xy
	sx, sy := -1, -1
	if origin.X < targetPos.X {
		sx = 1
	}
	if origin.Y < targetPos.Y {
		sy = 1
	}

	x, y := origin.X, origin.Y

	for {
		// Add the current cube to the result.
		result = append(result, Vec2{x, y})

		if x == targetPos.X && y == targetPos.Y {
			break
		}

		e2 := 2 * err
		if e2 >= dy {
			if x == targetPos.X {
				break
			}
			err += dy
			x += sx
		}
		if e2 <= dx {
			if y == targetPos.Y {
				break
			}
			err += dx
			y += sy
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

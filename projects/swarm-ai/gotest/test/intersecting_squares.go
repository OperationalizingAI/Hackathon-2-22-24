package test

type Vec2 struct {
	X, Y int
}

type Context struct{}

func (ctx *Context) ComputeIntersectingCubes(origin, target Vec2) []Vec2 {
	var result []Vec2
	x, y := origin.X, origin.Y
	dx := abs(target.X - origin.X)
	dy := abs(target.Y - origin.Y)
	stepX, stepY := 1, 1
	if origin.X > target.X {
		stepX = -1
	}
	if origin.Y > target.Y {
		stepY = -1
	}
	err := dx - dy
	dx *= 2
	dy *= 2

	for x != target.X || y != target.Y {
		e2 := err
		if e2 > -dy && x != target.X {
			err -= dy
			x += stepX
		}
		if e2 < dx && y != target.Y {
			err += dx
			y += stepY
		}
		// Only add the cube if it's not the origin or the final target cube
		if (x != origin.X || y != origin.Y) && (x != target.X || y != target.Y) {
			result = append(result, Vec2{x, y})
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

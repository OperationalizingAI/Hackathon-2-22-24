package solution

import "math"

type Vec2 struct {
	X, Y int
}

func GetMagAndSign(x int) (int, int) {
	if x < 0 {
		return -x, -1
	}
	return x, 1
}

func ComputeDistance(a Vec2, b Vec2) int {
	return int(math.Max(math.Abs(float64(a.X)-float64(b.X)), math.Abs(float64(a.Y)-float64(b.Y))))
}

type Context struct{}

func (c *Context) ComputeIntersectingCubes(casterPos Vec2, targetPos Vec2) []Vec2 {
	var intersectingCubes []Vec2
	dx, sign_dx := GetMagAndSign(targetPos.X - casterPos.X)
	dy, sign_dy := GetMagAndSign(targetPos.Y - casterPos.Y)

	var intersectingCube Vec2

	if dx == dy {
		for i := 1; i < dx; i++ {
			intersectingCube = Vec2{X: casterPos.X + i*sign_dx, Y: casterPos.Y + i*sign_dy}
			if ComputeDistance(casterPos, intersectingCube) == 0 {
				break
			}
			intersectingCubes = append(intersectingCubes, intersectingCube)
		}
	} else if dx > dy {
		j := 1
		for i := 1; i < dx; i++ {

			if i*dy > j*dx {
				j++
				i--
			} else if i*dy == j*dx {
				j++
			}

			intersectingCube = Vec2{X: casterPos.X + i*sign_dx, Y: casterPos.Y + (j-1)*sign_dy}
			intersectingCubes = append(intersectingCubes, intersectingCube)
		}
	} else {
		i := 1
		for j := 1; j < dy; j++ {

			if i*dy < j*dx {
				i++
				j--
			} else if i*dy == j*dx {
				i++
			}

			intersectingCube = Vec2{X: casterPos.X + (i-1)*sign_dx, Y: casterPos.Y + j*sign_dy}
			intersectingCubes = append(intersectingCubes, intersectingCube)
		}
	}
	return intersectingCubes
}

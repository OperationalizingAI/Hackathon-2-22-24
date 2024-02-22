package solution

import (
	"testing"
)

func arrayContainsVec2(slice []Vec2, vec Vec2) bool {
	for _, v := range slice {
		if v.X == vec.X && v.Y == vec.Y {
			return true
		}
	}
	return false
}

func TestComputeIntersectingCubes(t *testing.T) {
	t.Log("TestComputeIntersectingCubes")

	ctx := Context{}

	origin := Vec2{0, 0}
	var intersectingCubes []Vec2

	for _, testcase := range []struct {
		targetPos     Vec2
		expectedCubes []Vec2
	}{{
		targetPos: Vec2{5, 5},
		expectedCubes: []Vec2{
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4},
		},
	}, {
		targetPos: Vec2{-5, -5},
		expectedCubes: []Vec2{
			{-1, -1},
			{-2, -2},
			{-3, -3},
			{-4, -4},
		},
	}, {
		targetPos: Vec2{5, -5},
		expectedCubes: []Vec2{
			{1, -1},
			{2, -2},
			{3, -3},
			{4, -4},
		},
	}, {
		targetPos: Vec2{6, 3},
		expectedCubes: []Vec2{
			{1, 0},
			{2, 1},
			{3, 1},
			{4, 2},
			{5, 2},
		},
	}, {
		targetPos: Vec2{5, 3},
		expectedCubes: []Vec2{
			{1, 0},
			{1, 1},
			{2, 1},
			{3, 1},
			{3, 2},
			{4, 2},
		},
	}, {
		targetPos: Vec2{3, 7},
		expectedCubes: []Vec2{
			{0, 1},
			{0, 2},
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 4},
			{2, 5},
			{2, 6},
		},
	}} {

		t.Log("TestComputeIntersectingCubes for target pos", testcase.targetPos)
		intersectingCubes = ctx.ComputeIntersectingCubes(origin, testcase.targetPos)

		for _, expectedCube := range testcase.expectedCubes {
			if !arrayContainsVec2(intersectingCubes, expectedCube) {
				t.Errorf("Intersecting cubes missing %v", expectedCube)
			}
		}

		for _, intersectingCube := range intersectingCubes {
			if !arrayContainsVec2(testcase.expectedCubes, intersectingCube) {
				t.Errorf("Unexpected intersecting cube %v", intersectingCube)
			}
		}
	}
}

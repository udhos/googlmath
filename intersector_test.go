package math

import (
	. "gopkg.in/check.v1"
)

type IsPointInTriangleTestValue struct {
	Point      Vector3
	T1, T2, T3 Vector3
	Expected   bool
}

func (s *S) TestIsPointInTriangle(c *C) {
	containTestTable := []IsPointInTriangleTestValue{
		{Vec3(0.5, 0.5, 0), Vec3(0, 0, 0), Vec3(1, 1, 0), Vec3(0, 1, 0), true},
		{Vec3(2, 0.5, 0), Vec3(0, 0, 0), Vec3(1, 1, 0), Vec3(0, 1, 0), false},
		{Vec3(1, 1, 1), Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0), false},
	}
	for _, value := range containTestTable {
		c.Assert(IsPointInTriangle(value.Point, value.T1, value.T2, value.T3), Equals, value.Expected)
	}
}

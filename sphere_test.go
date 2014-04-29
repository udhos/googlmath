package math

import (
	. "gopkg.in/check.v1"
)

type SphereOverlapsTestValue struct {
	Sphere, Sphere2 Sphere
	Expected        bool
}

func (s *S) TestSphereOverlaps(c *C) {
	containTestTable := []SphereOverlapsTestValue{
		SphereOverlapsTestValue{Spe(Vec3(1, -2, 0), 12.0), Spe(Vec3(0, 2, 0), 30.0), true},
	}
	for i := range containTestTable {
		value := containTestTable[i]
		c.Assert(value.Sphere.Overlaps(value.Sphere2), Equals, value.Expected)
	}
}

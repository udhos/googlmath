package math

import (
	. "gopkg.in/check.v1"
)

type CircleContainTestValue struct {
	X, Y     float32
	Expected bool
}

func (s *S) TestCircleContains(c *C) {
	circle := Circ(0, 0, 1.0)
	containTestTable := []CircleContainTestValue{
		CircleContainTestValue{0, 0, true},
		CircleContainTestValue{1, 0, true},
		CircleContainTestValue{-2, -2, false},
		CircleContainTestValue{2, 0, false}}

	// Contains
	for i := range containTestTable {
		c.Assert(circle.Contains(containTestTable[i].X, containTestTable[i].Y), Equals, containTestTable[i].Expected)
	}
}

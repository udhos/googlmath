package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_2F32_B struct {
	X, Y     float32
	Expected bool
}

func (s *S) TestCircleContains(c *C) {
	circle := Circ(0, 0, 1.0)
	containTestTable := []TestValue_2F32_B{
		TestValue_2F32_B{0, 0, true},
		TestValue_2F32_B{1, 0, true},
		TestValue_2F32_B{-2, -2, false},
		TestValue_2F32_B{2, 0, false}}

	// Contains
	for i := range containTestTable {
		c.Assert(circle.Contains(Vec2(containTestTable[i].X, containTestTable[i].Y)), Equals, containTestTable[i].Expected)
	}
}

package math

import (
	. "gopkg.in/check.v1"
)

type EllipseContainTestValue struct {
	X, Y     float32
	Expected bool
}

func (s *S) TestEllipse(c *C) {
	e := Ell(0, 0, 1, 1)
	containTestTable := []EllipseContainTestValue{
		EllipseContainTestValue{0, 0, true},
		EllipseContainTestValue{1, 0, true},
		EllipseContainTestValue{-2, -2, false},
		EllipseContainTestValue{2, 0, false},
	}
	// Contains
	for i := range containTestTable {
		c.Assert(e.Contains(containTestTable[i].X, containTestTable[i].Y), Equals, containTestTable[i].Expected)
	}
}

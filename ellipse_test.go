package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestEllipse(c *C) {
	e := Ell(0, 0, 1, 1)
	containTestTable := []TestValue_2F32_B{
		TestValue_2F32_B{0, 0, true},
		TestValue_2F32_B{1, 0, true},
		TestValue_2F32_B{-2, -2, false},
		TestValue_2F32_B{2, 0, false},
	}
	// Contains
	for i := range containTestTable {
		c.Assert(e.Contains(containTestTable[i].X, containTestTable[i].Y), Equals, containTestTable[i].Expected)
	}
}

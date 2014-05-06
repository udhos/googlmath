package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_4F32_Ell struct {
	f0 float32
	f1 float32
	f2 float32
	f3 float32
	e  Ellipse
}

func (s *S) TestEllipseEll(c *C) {
	tests := []TestValue_4F32_Ell{
		{
			0, 0, 0, 0, Ellipse{0, 0, 0, 0},
		},
		{
			-1.2, -3.3, 1, 2, Ellipse{-1.2, -3.3, 1, 2},
		},
	}
	for _, t := range tests {
		obtained := Ell(t.f0, t.f1, t.f2, t.f3)
		c.Assert(obtained, DeepEquals, t.e)
	}
}

type TestValue_Ell2F32_B struct {
	ell Ellipse
	f0  float32
	f1  float32
	e   bool
}

func (s *S) TestEllipseContains(c *C) {
	tests := []TestValue_Ell2F32_B{
		{Ellipse{0, 0, 1, 1}, 0, 0, true},
		{Ellipse{0, 0, 1, 1}, 1, 0, true},
		{Ellipse{0, 0, 1, 1}, -2, -2, false},
		{Ellipse{0, 0, 1, 1}, 2, 0, false},
		{Ellipse{0, 0, 0, 1}, 2, 0, false},
		{Ellipse{0, 0, 1, 0}, 2, 0, false},
		{Ellipse{1, 2, 0, 0}, 2, 0, false},
	}
	// Contains
	for _, t := range tests {
		obtained := t.ell.Contains(t.f0, t.f1)
		c.Assert(obtained, Equals, t.e)
	}
}

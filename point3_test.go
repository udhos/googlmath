package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestPoint3Pt(c *C) {
	tests := []struct{ X, Y, Z int }{
		{0, 0, 0},
		{2, 0, 1},
		{0, 4, 0},
		{2, 5, 0},
		{-3, 0, 0},
		{0, -5, 2},
		{-1, -22, -3},
		{-3, 5, 1},
	}
	for _, t := range tests {
		actual := Pt3(t.X, t.Y, t.Z)
		expected := Point3{t.X, t.Y, t.Z}
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPoint3Add(c *C) {
	tests := []struct{ p0, p1, r Point3 }{
		{Point3{0, 0, 0}, Point3{0, 0, 0}, Point3{0, 0, 0}},
		{Point3{2, 0, 0}, Point3{0, 0, 0}, Point3{2, 0, 0}},
		{Point3{2, 0, 0}, Point3{0, 5, 0}, Point3{2, 5, 0}},
		{Point3{-3, 2, 4}, Point3{1, -10, 0}, Point3{-2, -8, 4}},
	}
	for _, t := range tests {
		actual := t.p0.Add(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPoint3Sub(c *C) {
	tests := []struct{ p0, p1, r Point3 }{
		{Point3{0, 0, 0}, Point3{0, 0, 0}, Point3{0, 0, 0}},
		{Point3{2, 0, 0}, Point3{0, 0, 1}, Point3{2, 0, -1}},
		{Point3{2, 0, 0}, Point3{0, 5, 2}, Point3{2, -5, -2}},
		{Point3{-3, 2, 0}, Point3{1, -10, 3}, Point3{-4, 12, -3}},
	}
	for _, t := range tests {
		actual := t.p0.Sub(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPoint3Mul(c *C) {
	tests := []struct{ p0, p1, r Point3 }{
		{Point3{0, 0, 0}, Point3{0, 0, 0}, Point3{0, 0, 0}},
		{Point3{2, 0, 0}, Point3{0, 0, 0}, Point3{0, 0, 0}},
		{Point3{2, 0, 0}, Point3{0, 5, 0}, Point3{0, 0, 0}},
		{Point3{-3, 2, 0}, Point3{1, -10, 0}, Point3{-3, -20, 0}},
	}
	for _, t := range tests {
		actual := t.p0.Mul(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPoint3Div(c *C) {
	tests := []struct{ p0, p1, r Point3 }{
		{Point3{2, 0, 0}, Point3{1, 5, 1}, Point3{2, 0, 0}},
		{Point3{-3, 2, 1}, Point3{1, -10, 1}, Point3{-3, 0, 1}},
	}
	for _, t := range tests {
		actual := t.p0.Div(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

type TestValue_2P3_B struct {
	p0, p1 Point3
	r      bool
}

func (s *S) TestPoint3Eq(c *C) {
	tests := []TestValue_2P3_B{
		{Point3{0, 0, 0}, Point3{0, 0, 0}, true},
		{Point3{-3, 2, 0}, Point3{-3, 2, 0}, true},
		{Point3{2, 0, 0}, Point3{0, 0, 0}, false},
		{Point3{2, 0, 0}, Point3{0, 0, 5}, false},
		{Point3{-3, 2, 0}, Point3{1, -10, 0}, false},
	}
	for _, t := range tests {
		actual := t.p0.Eq(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

type TestValue_P3_Str struct {
	p Point3
	r string
}

func (s *S) TestPoint3String(c *C) {
	tests := []TestValue_P3_Str{
		{Point3{0, 0, 0}, "(0,0,0)"},
		{Point3{-3, 2, 0}, "(-3,2,0)"},
		{Point3{200, 0, 0}, "(200,0,0)"},
		{Point3{22, 12, 2}, "(22,12,2)"},
		{Point3{-31, 22, 1}, "(-31,22,1)"},
	}
	for _, t := range tests {
		actual := t.p.String()
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestPointPt(c *C) {
	tests := []struct{ X, Y int }{
		{0, 0},
		{2, 0},
		{0, 4},
		{2, 5},
		{-3, 0},
		{0, -5},
		{-1, -22},
		{-3, 5},
	}
	for _, t := range tests {
		actual := Pt(t.X, t.Y)
		expected := Point{t.X, t.Y}
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPointAdd(c *C) {
	tests := []struct{ p0, p1, r Point }{
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},
		{Point{2, 0}, Point{0, 0}, Point{2, 0}},
		{Point{2, 0}, Point{0, 5}, Point{2, 5}},
		{Point{-3, 2}, Point{1, -10}, Point{-2, -8}},
	}
	for _, t := range tests {
		actual := t.p0.Add(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPointSub(c *C) {
	tests := []struct{ p0, p1, r Point }{
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},
		{Point{2, 0}, Point{0, 0}, Point{2, 0}},
		{Point{2, 0}, Point{0, 5}, Point{2, -5}},
		{Point{-3, 2}, Point{1, -10}, Point{-4, 12}},
	}
	for _, t := range tests {
		actual := t.p0.Sub(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPointMul(c *C) {
	tests := []struct{ p0, p1, r Point }{
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},
		{Point{2, 0}, Point{0, 0}, Point{0, 0}},
		{Point{2, 0}, Point{0, 5}, Point{0, 0}},
		{Point{-3, 2}, Point{1, -10}, Point{-3, -20}},
	}
	for _, t := range tests {
		actual := t.p0.Mul(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

func (s *S) TestPointDiv(c *C) {
	tests := []struct{ p0, p1, r Point }{
		{Point{2, 0}, Point{1, 5}, Point{2, 0}},
		{Point{-3, 2}, Point{1, -10}, Point{-3, 0}},
	}
	for _, t := range tests {
		actual := t.p0.Div(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

type PointEqTestValue struct {
	p0, p1 Point
	r      bool
}

func (s *S) TestPointEq(c *C) {
	tests := []PointEqTestValue{
		{Point{0, 0}, Point{0, 0}, true},
		{Point{-3, 2}, Point{-3, 2}, true},
		{Point{2, 0}, Point{0, 0}, false},
		{Point{2, 0}, Point{0, 5}, false},
		{Point{-3, 2}, Point{1, -10}, false},
	}
	for _, t := range tests {
		actual := t.p0.Eq(t.p1)
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

type PointStringTestValue struct {
	p Point
	r string
}

func (s *S) TestPointString(c *C) {
	tests := []PointStringTestValue{
		{Point{0, 0}, "(0,0)"},
		{Point{-3, 2}, "(-3,2)"},
		{Point{200, 0}, "(200,0)"},
		{Point{22, 12}, "(22,12)"},
		{Point{-31, 22}, "(-31,22)"},
	}
	for _, t := range tests {
		actual := t.p.String()
		expected := t.r
		c.Assert(actual, Equals, expected)
	}
}

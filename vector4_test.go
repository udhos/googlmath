package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestVector4Vec4(c *C) {
	tests := []struct{ X, Y, Z, W float32 }{
		{0, 0, 0, 0},
		{2, 0, 1, 0},
		{0, 4, 0, 0},
		{2, 5, 0, 0.1},
		{-3, 0, 0, 0.5},
		{0, -5, 0, 2.3},
		{-1, -22, -3.3, 0},
		{-3.3, 5.1, 1.2, 1.1},
	}
	for _, t := range tests {
		actual := Vec4(t.X, t.Y, t.Z, t.W)
		expected := Vector4{t.X, t.Y, t.Z, t.W}
		c.Assert(actual, Vector4Check, expected)
	}
}

func (s *S) TestVector4Add(c *C) {
	tests := []struct{ v0, v1, r Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{2, 0, 1, 0}, Vector4{2, 0, 1, 0}},
		{Vector4{2, 5, 0, 0.1}, Vector4{-3, 0, 0, 0.5}, Vector4{-1, 5, 0, 0.6}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.5}, Vector4{-4, 7, 1.2, 0.6}},
	}
	for _, t := range tests {
		actual := t.v0.Add(t.v1)
		expected := t.r
		c.Assert(actual, Vector4Check, expected)
	}
}

func (s *S) TestVector4Sub(c *C) {
	tests := []struct{ v0, v1, r Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{2, 0, 1, 0}, Vector4{-2, 0, -1, 0}},
		{Vector4{2, 5, 0, 0.1}, Vector4{-3, 0, 0, 0.5}, Vector4{5, 5, 0, -0.4}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.5}, Vector4{2, 3, -1.2, -0.4}},
	}
	for _, t := range tests {
		actual := t.v0.Sub(t.v1)
		expected := t.r
		c.Assert(actual, Vector4Check, expected)
	}
}

func (s *S) TestVector4Mul(c *C) {
	tests := []struct{ v0, v1, r Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{2, 0, 1, 0}, Vector4{0, 0, 0, 0}},
		{Vector4{2, 5, 0, 0.1}, Vector4{-3, 0, 0, 0.5}, Vector4{-6, 0, 0, 0.05}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.1}, Vector4{3, 10, 0, 0.01}},
	}
	for _, t := range tests {
		actual := t.v0.Mul(t.v1)
		expected := t.r
		c.Assert(actual, Vector4Check, expected)
	}
}

func (s *S) TestVector4Div(c *C) {
	tests := []struct{ v0, v1, r Vector4 }{
		{Vector4{2, 5, 0, 0.1}, Vector4{-2, 1, 1, 0.5}, Vector4{-1, 5, 0, 0.2}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.1}, Vector4{1.0 / 3, 2.5, 0, 1}},
	}
	for _, t := range tests {
		actual := t.v0.Div(t.v1)
		expected := t.r
		c.Assert(actual, Vector4Check, expected)
	}
}

type Vector4ScaleTestValue struct {
	v Vector4
	s float32
	r Vector4
}

func (s *S) TestVector4Scale(c *C) {
	tests := []Vector4ScaleTestValue{
		{Vector4{2, 5, 0, 0.1}, 1.0, Vector4{2, 5, 0, 0.1}},
		{Vector4{2, 5, 0, 0.1}, 0, Vector4{0, 0, 0, 0}},
		{Vector4{-1, 5, 0, 0.1}, -1.0, Vector4{1, -5, 0, -0.1}},
		{Vector4{2, 5, 0, 0.1}, 1.2, Vector4{2.4, 6, 0, 0.12}},
	}
	for _, t := range tests {
		actual := t.v.Scale(t.s)
		expected := t.r
		c.Assert(actual, Vector4Check, expected)
	}
}

package math

import (
	"fmt"
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
	tests := []struct{ v0, v1, e Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{2, 0, 1, 0}, Vector4{2, 0, 1, 0}},
		{Vector4{2, 5, 0, 0.1}, Vector4{-3, 0, 0, 0.5}, Vector4{-1, 5, 0, 0.6}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.5}, Vector4{-4, 7, 1.2, 0.6}},
	}
	for _, t := range tests {
		actual := t.v0.Add(t.v1)
		c.Assert(actual, Vector4Check, t.e)
	}
}

func (s *S) TestVector4Sub(c *C) {
	tests := []struct{ v0, v1, e Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{2, 0, 1, 0}, Vector4{-2, 0, -1, 0}},
		{Vector4{2, 5, 0, 0.1}, Vector4{-3, 0, 0, 0.5}, Vector4{5, 5, 0, -0.4}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.5}, Vector4{2, 3, -1.2, -0.4}},
	}
	for _, t := range tests {
		actual := t.v0.Sub(t.v1)
		c.Assert(actual, Vector4Check, t.e)
	}
}

func (s *S) TestVector4Mul(c *C) {
	tests := []struct{ v0, v1, e Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{2, 0, 1, 0}, Vector4{0, 0, 0, 0}},
		{Vector4{2, 5, 0, 0.1}, Vector4{-3, 0, 0, 0.5}, Vector4{-6, 0, 0, 0.05}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.1}, Vector4{3, 10, 0, 0.01}},
	}
	for _, t := range tests {
		actual := t.v0.Mul(t.v1)
		c.Assert(actual, Vector4Check, t.e)
	}
}

func (s *S) TestVector4Div(c *C) {
	tests := []struct{ v0, v1, e Vector4 }{
		{Vector4{2, 5, 0, 0.1}, Vector4{-2, 1, 1, 0.5}, Vector4{-1, 5, 0, 0.2}},
		{Vector4{-1, 5, 0, 0.1}, Vector4{-3, 2, 1.2, 0.1}, Vector4{1.0 / 3, 2.5, 0, 1}},
	}
	for _, t := range tests {
		actual := t.v0.Div(t.v1)
		c.Assert(actual, Vector4Check, t.e)
	}
}

type TestValue_V4F32_V4 struct {
	v Vector4
	f float32
	e Vector4
}

func (s *S) TestVector4Scale(c *C) {
	tests := []TestValue_V4F32_V4{
		{Vector4{2, 5, 0, 0.1}, 1.0, Vector4{2, 5, 0, 0.1}},
		{Vector4{2, 5, 0, 0.1}, 0, Vector4{0, 0, 0, 0}},
		{Vector4{-1, 5, 0, 0.1}, -1.0, Vector4{1, -5, 0, -0.1}},
		{Vector4{2, 5, 0, 0.1}, 1.2, Vector4{2.4, 6, 0, 0.12}},
	}
	for _, t := range tests {
		actual := t.v.Scale(t.f)
		c.Assert(actual, Vector4Check, t.e)
	}
}

type TestValue_V4_F32 struct {
	v Vector4
	e float32
}

func (s *S) TestVector4Len(c *C) {
	tests := []TestValue_V4_F32{
		{Vector4{2, 0, 0, 0}, 2.0},
		{Vector4{0, 1, 0, 0}, 1.0},
		{Vector4{0, 0, 0, 0}, 0},
		{Vector4{-1, 0, 0, 0}, 1.0},
		{Vector4{2, 2, 0, 0}, 2.828427},
	}
	for _, t := range tests {
		actual := t.v.Len()
		c.Assert(actual, Equals, t.e)
	}
}

func (s *S) TestVector4Len2(c *C) {
	tests := []TestValue_V4_F32{
		{Vector4{2, 0, 0, 0}, 4.0},
		{Vector4{0, 1, 0, 0}, 1.0},
		{Vector4{0, 0, 0, 0}, 0},
		{Vector4{-1, 0, 0, 0}, 1.0},
		{Vector4{2, 2, 0, 0}, 8.0},
	}
	for _, t := range tests {
		actual := t.v.Len2()
		c.Assert(actual, Equals, t.e)
	}
}

func (s *S) TestVector4Nor(c *C) {
	tests := []struct{ v, e Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{0, 0, 0, 0}},
		{Vector4{1, -1, 1, 1}, Vector4{0.5, -0.5, 0.5, 0.5}},
		{Vector4{2, 0, 1, 1}, Vector4{0.816497, 0.0, 0.408248, 0.408248}},
	}
	for _, t := range tests {
		actual := t.v.Nor()
		c.Assert(actual, Vector4Check, t.e)
	}
}

type TestValue_V4_B struct {
	v Vector4
	e bool
}

func (s *S) TestVector4IsZero(c *C) {
	tests := []TestValue_V4_B{
		{Vector4{0, 0, 0, 0}, true},
		{Vector4{1, -1, 1, 1}, false},
		{Vector4{2, 0, 1, 1}, false},
		{Vector4{0.816497, 0.0, 0.408248, 0.408248}, false},
	}
	for _, t := range tests {
		actual := t.v.IsZero()
		c.Assert(actual, Equals, t.e)
	}
}

func (s *S) TestVector4IsUnit(c *C) {
	tests := []TestValue_V4_B{
		{Vector4{1, 0, 0, 0}, true},
		{Vector4{0, 0.0, 0.408248, 0}, false},
		{Vector4{-1, 0, 0, 0}, true},
		{Vector4{1, -1, 1, 1}, false},
		{Vector4{2, 0, 1, 0}, false},
		{Vector4{0.816497, 0.0, 0.408248, 0.408248}, false},
	}
	for _, t := range tests {
		actual := t.v.IsUnit()
		c.Assert(actual, Equals, t.e, Commentf("%v.IsUnit(), actual: %t expected: %t", t.v, actual, t.e))
	}
}

func (s *S) TestVector4Invert(c *C) {
	tests := []struct{ v, e Vector4 }{
		{Vector4{0, 0, 0, 0}, Vector4{0, 0, 0, 0}},
		{Vector4{1, -1, 1, 1}, Vector4{-1, 1, -1, -1}},
		{Vector4{2, 0, 1, 1}, Vector4{-2, 0, -1, -1}},
		{Vector4{0.816497, 0.0, 0.408248, 0.408248}, Vector4{-0.816497, 0.0, -0.408248, -0.408248}},
	}
	for _, t := range tests {
		actual := t.v.Invert()
		c.Assert(actual, Vector4Check, t.e)
	}
}

func (s *S) TestVector4Eq(c *C) {
	tests := []Vector4{
		Vector4{0, 0, 0, 0}, Vector4{0, 0, 0, 0},
		Vector4{1, -1, 1, 1}, Vector4{-1, 1, -1, -1},
		Vector4{2, 0, 1, 1}, Vector4{-2, 0, -1, -1},
		Vector4{0.816497, 0.0, 0.408248, 0.408248}, Vector4{-0.816497, 0.0, -0.408248, -0.408248},
	}
	for i := range tests {
		v0 := tests[i]
		for j := range tests {
			v1 := tests[j]
			e := v0.Eq(v1)
			actual := v0.X == v1.X && v0.Y == v1.Y && v0.Z == v1.Z && v0.W == v1.W
			c.Assert(actual, Equals, e)
		}
	}
}

func (s *S) TestVector4String(c *C) {
	tests := []Vector4{
		Vector4{0, 0, 0, 0}, Vector4{0, 0, 0, 0},
		Vector4{1, -1, 1, 1}, Vector4{-1, 1, -1, -1},
		Vector4{2, 0, 1, 1}, Vector4{-2, 0, -1, -1},
		Vector4{0.816497, 0.0, 0.408248, 0.408248}, Vector4{-0.816497, 0.0, -0.408248, -0.408248},
	}
	for _, v := range tests {
		actual := v.String()
		c.Assert(actual, Equals, fmt.Sprintf("%v", v))
	}
}

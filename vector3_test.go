package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestVec3(c *C) {
	tests := []struct{ X, Y, Z float32 }{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 1},
		{-1, 0, 0},
		{-1, 0, -1},
		{-22.013, 0.125, 12.3},
		{-0.013, 1.025, 10.01},
	}
	var vec Vector3
	for _, t := range tests {
		vec = Vec3(t.X, t.Y, t.Z)
		c.Assert(vec, Vector3Check, Vector3{t.X, t.Y, t.Z})
	}
}

func (s *S) TestVector3Add(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(1, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(-1, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(4, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(-3.2, 20, -3)},
	}
	var r Vector3
	for _, t := range tests {
		r = t.v1.Add(t.v2)
		c.Assert(r, Vector3Check, t.v3)
	}
}

func (s *S) TestVector3Sub(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(1, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(-1, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(-2, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(-0.8, 24, -3.6)},
	}
	var r Vector3
	for _, t := range tests {
		r = t.v1.Sub(t.v2)
		c.Assert(r, Vector3Check, t.v3)
	}
}

func (s *S) TestVector3Mul(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(3, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(2.4, -44, -0.99)},
	}
	var r Vector3
	for _, t := range tests {
		r = t.v1.Mul(t.v2)
		c.Assert(r, Vector3Check, t.v3)
	}
}

func (s *S) TestVector3Div(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(-2/-1.2, -11, -11)},
	}
	var r Vector3
	for _, t := range tests {
		r = t.v1.Div(t.v2)
		c.Assert(r, Vector3Check, t.v3)
	}
}

type TestValue_V3_F32 struct {
	v Vector3
	e float32
}

func (s *S) TestVector3Len(c *C) {
	tests := []TestValue_V3_F32{
		{Vector3{0, 0, 0}, 0},
		{Vector3{1, 0, 0}, 1},
		{Vector3{0, 1, 0}, 1},
		{Vector3{0, 0, 1}, 1},
		{Vector3{-1, 0, 0}, 1},
		{Vector3{0, -1, 0}, 1},
		{Vector3{0, 0, -1}, 1},
		{Vector3{2, 1, 3}, 3.741657386},
	}
	for _, t := range tests {
		obtained := t.v.Len()
		c.Assert(obtained, Equals, t.e)
	}
}

func (s *S) TestVector3Len2(c *C) {
	tests := []TestValue_V3_F32{
		{Vector3{0, 0, 0}, 0},
		{Vector3{1, 0, 0}, 1},
		{Vector3{0, 1, 0}, 1},
		{Vector3{0, 0, 1}, 1},
		{Vector3{-1, 0, 0}, 1},
		{Vector3{0, -1, 0}, 1},
		{Vector3{0, 0, -1}, 1},
		{Vector3{2, 1, 3}, 14},
	}
	for _, t := range tests {
		obtained := t.v.Len2()
		c.Assert(obtained, Equals, t.e)
	}
}

type TestValue_2V3_F32 struct {
	v0 Vector3
	v1 Vector3
	e  float32
}

func (s *S) TestVector3Distance(c *C) {
	tests := []TestValue_2V3_F32{
		{Vector3{0, 0, 0}, Vector3{0, 0, 0}, 0},
		{Vector3{1, 0, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, 1, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, 0, 1}, Vector3{0, 0, 0}, 1},
		{Vector3{-1, 0, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, -1, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, 0, -1}, Vector3{0, 0, 0}, 1},
		{Vector3{2, 1, 3}, Vector3{0, 0, 0}, 3.741657386},
		{Vector3{2, 0, 0}, Vector3{4, 0, 0}, 2},
	}
	for _, t := range tests {
		obtained := t.v0.Distance(t.v1)
		c.Assert(obtained, Equals, t.e)
	}
}

func (s *S) TestVector3Distance2(c *C) {
	tests := []TestValue_2V3_F32{
		{Vector3{0, 0, 0}, Vector3{0, 0, 0}, 0},
		{Vector3{1, 0, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, 1, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, 0, 1}, Vector3{0, 0, 0}, 1},
		{Vector3{-1, 0, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, -1, 0}, Vector3{0, 0, 0}, 1},
		{Vector3{0, 0, -1}, Vector3{0, 0, 0}, 1},
		{Vector3{2, 1, 3}, Vector3{0, 0, 0}, 14},
		{Vector3{2, 0, 0}, Vector3{4, 0, 0}, 4},
	}
	for _, t := range tests {
		obtained := t.v0.Distance2(t.v1)
		c.Assert(obtained, Equals, t.e)
	}
}

func (s *S) TestVector3Nor(c *C) {
	tests := []struct{ v, e Vector3 }{
		{Vector3{0, 0, 0}, Vector3{0, 0, 0}},
		{Vector3{1, -1, 1}, Vector3{0.57735, -0.57735, 0.57735}},
		{Vector3{2, 0, 1}, Vector3{0.894427, 0, 0.447214}},
	}
	for _, t := range tests {
		obtained := t.v.Nor()
		c.Assert(obtained, Vector3Check, t.e)
	}
}

func (s *S) TestVector3Dot(c *C) {
	tests := []TestValue_2V3_F32{
		{Vector3{0, 0, 0}, Vector3{0, 0, 0}, 0},
		{Vector3{1, 2, 3}, Vector3{3, 4, 5}, 26},
		{Vector3{1.1, 0.2, 1.3}, Vector3{-3.3, 4.2, 5.5}, 4.36},
	}
	for _, t := range tests {
		obtained := t.v0.Dot(t.v1)
		c.Assert(obtained, EqualsFloat32, t.e)
	}
}

type TestValue_2V3_V3 struct {
	v0 Vector3
	v1 Vector3
	e  Vector3
}

func (s *S) TestVector3Cross(c *C) {
	tests := []TestValue_2V3_V3{
		{Vector3{0, 0, 0}, Vector3{0, 0, 0}, Vector3{0, 0, 0}},
		{Vector3{0.1, -0.2, 102.3}, Vector3{-3.3, -14.2, 0.5}, Vector3{1452.56, -337.64, -2.08}},
		{Vector3{1.1, 0.2, 1.3}, Vector3{-3.3, 4.2, 5.5}, Vector3{-4.36, -10.34, 5.28}},
	}
	for _, t := range tests {
		obtained := t.v0.Cross(t.v1)
		c.Assert(obtained, Vector3Check, t.e)
	}
}

// TODO MulMatrix
// TODO Prj
// TODO Rot
// TODO IsUnit
// TODO IsZero
// TODO Lerp
// TODO Slerp

type TestValue_V3F32_V3 struct {
	v Vector3
	f float32
	e Vector3
}

func (s *S) TestVector3Limit(c *C) {
	tests := []TestValue_V3F32_V3{
		{Vector3{0, 0, 0}, 0, Vector3{0, 0, 0}},
		{Vector3{2, -2, 1}, 1, Vector3{2.0 / 3, -2.0 / 3, 1.0 / 3}},
		{Vector3{2, 2, 2}, 1, Vector3{0.57735, 0.57735, 0.57735}},
	}
	for _, t := range tests {
		obtained := t.v.Limit(t.f)
		c.Assert(obtained, Vector3Check, t.e)
	}
}

func (s *S) TestVector3Scale(c *C) {
	tests := []TestValue_V3F32_V3{
		{Vector3{0, 0, 0}, 0, Vector3{0, 0, 0}},
		{Vector3{2, -2, 1}, 1, Vector3{2, -2, 1}},
		{Vector3{2, 1, 0}, -2.2, Vector3{-4.4, -2.2, 0}},
	}
	for _, t := range tests {
		obtained := t.v.Scale(t.f)
		c.Assert(obtained, Vector3Check, t.e)
	}
}

func (s *S) TestVector3Invert(c *C) {
	tests := []struct{ v, e Vector3 }{
		{Vector3{0, 0, 0}, Vector3{0, 0, 0}},
		{Vector3{2, -2, 1}, Vector3{-2, 2, -1}},
		{Vector3{3.3, 1, 0}, Vector3{-3.3, -1, 0}},
	}
	for _, t := range tests {
		obtained := t.v.Invert()
		c.Assert(obtained, Vector3Check, t.e)
	}
}

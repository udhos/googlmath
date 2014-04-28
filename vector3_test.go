package math

import (
	. "launchpad.net/gocheck"
)

type Vector3TestSuite struct {
	vec Vector3
}

var _ = Suite(&Vector3TestSuite{})

func (s *Vector3TestSuite) Vec3(c *C) {
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
		c.Assert(vec, Equals, Vector3{t.X, t.Y, t.Z})
	}
}

func (s *Vector3TestSuite) Vector3Set(c *C) {
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
		vec.Set(t.X, t.Y, t.Z)
		c.Assert(vec, Equals, Vector3{t.X, t.Y, t.Z})
	}
}

func (s *Vector3TestSuite) Vector3SetVec3(c *C) {
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
	var vec2 Vector3
	for _, t := range tests {
		vec2 = vec
		vec.SetVec3(Vec3(t.X, t.Y, t.Z))
		c.Assert(vec, Equals, Vector3{t.X, t.Y, t.Z})
		c.Assert(vec, Not(Equals), vec2)
	}
}

func (s *Vector3TestSuite) Vector3Cpy(c *C) {
	tests := []struct{ X, Y, Z float32 }{
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
	var vec2 Vector3
	for _, t := range tests {
		vec = Vec3(t.X, t.Y, t.Z)
		vec2 = Vec3(0, 0, 0)

		vec2 = vec.Cpy()

		c.Assert(vec, Equals, vec2)
		c.Assert(vec, Not(Equals), Vec3(0, 0, 0))
		c.Assert(vec2, Not(Equals), Vec3(0, 0, 0))
	}
}

func (s *Vector3TestSuite) Vector3Clr(c *C) {
	tests := []struct{ X, Y, Z float32 }{
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
	var vec2 Vector3
	for _, t := range tests {
		vec = Vec3(t.X, t.Y, t.Z)
		vec2 = vec

		vec.Clr()

		c.Assert(vec, Equals, Vec3(0, 0, 0))
		c.Assert(vec, Not(Equals), vec2)
	}
}

func (s *Vector3TestSuite) Vector3Add(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(1, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(-1, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(4, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(-3.2, 20, -3)},
	}
	var vec Vector3
	var vec2 Vector3
	for _, t := range tests {
		vec = Vec3(0, 0, 0)
		vec2 = Vec3(0, 0, 0)

		vec = t.v1
		vec2 = t.v1.Add(t.v2)

		c.Assert(vec, Not(Equals), vec2)
		c.Assert(vec, Equals, t.v1)
		c.Assert(vec2, Equals, t.v1)
		c.Assert(vec2, Equals, t.v3)
	}
}

func (s *Vector3TestSuite) Vector3Sub(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(1, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(-1, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(-2, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(-0.8, 24, -3.6)},
	}
	var vec Vector3
	var vec2 Vector3
	for _, t := range tests {
		vec = Vec3(0, 0, 0)
		vec2 = Vec3(0, 0, 0)

		vec = t.v1
		vec2 = t.v1.Sub(t.v2)

		c.Assert(vec, Not(Equals), vec2)
		c.Assert(vec, Equals, t.v1)
		c.Assert(vec2, Equals, t.v1)
		c.Assert(vec2, Equals, t.v3)
	}
}

func (s *Vector3TestSuite) Vector3Mul(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(3, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(2.4, -44, -1.1)},
	}
	var vec Vector3
	var vec2 Vector3
	for _, t := range tests {
		vec = Vec3(0, 0, 0)
		vec2 = Vec3(0, 0, 0)

		vec = t.v1
		vec2 = t.v1.Mul(t.v2)

		c.Assert(vec, Not(Equals), vec2)
		c.Assert(vec, Equals, t.v1)
		c.Assert(vec2, Equals, t.v1)
		c.Assert(vec2, Equals, t.v3)
	}
}

func (s *Vector3TestSuite) Vector3Div(c *C) {
	tests := []struct{ v1, v2, v3 Vector3 }{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(-1, 0, 0), Vec3(0, 0, 0), Vec3(0, 0, 0)},
		{Vec3(1, 0, 0), Vec3(3, 0, 0), Vec3(1/3, 0, 0)},
		{Vec3(-2, 22, -3.3), Vec3(-1.2, -2, 0.3), Vec3(-2/-1.2, -11, -9.9)},
	}
	var vec Vector3
	var vec2 Vector3
	for _, t := range tests {
		vec = Vec3(0, 0, 0)
		vec2 = Vec3(0, 0, 0)

		vec = t.v1
		vec2 = t.v1.Div(t.v2)

		c.Assert(vec, Not(Equals), vec2)
		c.Assert(vec, Equals, t.v1)
		c.Assert(vec2, Equals, t.v1)
		c.Assert(vec2, Equals, t.v3)
	}
}

func (s *Vector3TestSuite) Vector3Len(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Len(), Equals, 2)
}

func (s *Vector3TestSuite) Vector3Len2(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Len2(), Equals, 4)
}

func (s *Vector3TestSuite) Vector3Distance(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Distance(Vec3(0, 0, 0)), Equals, 2)
}

func (s *Vector3TestSuite) Vector3Distance2(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Distance2(Vec3(0, 0, 0)), Equals, 4)
}

func (s *Vector3TestSuite) Vector3Nor(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Nor(), Equals, Vec3(1, 0, 0))
}

func (s *Vector3TestSuite) Vector3Dot(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Dot(Vec3(2, 0, 0)), Equals, 4)
}

func (s *Vector3TestSuite) Vector3Cross(c *C) {
	s.vec.SetVec3(Vec3(2, 1, 4))
	c.Assert(s.vec.Cross(Vec3(2, -3, 0)), Equals, Vec3(12, 6, -9))
}

// TODO MulMatrix
// TODO Prj
// TODO Rot
// TODO IsUnit
// TODO IsZero
// TODO Lerp
// TODO Slerp

func (s *Vector3TestSuite) Vector3Limit(c *C) {
	s.vec.SetVec3(Vec3(4, 0, 0))
	c.Assert(s.vec.Limit(3.3), Equals, Vec3(3.3, 0, 0))
}

func (s *Vector3TestSuite) Vector3Scale(c *C) {
	s.vec.SetVec3(Vec3(4, 0, -2))
	c.Assert(s.vec.Scale(3), Equals, Vec3(12, 0, -6))
}

func (s *Vector3TestSuite) Vector3Invert(c *C) {
	s.vec.SetVec3(Vec3(4, 0, -2))
	c.Assert(s.vec.Invert(), Equals, Vec3(-4, 0, -2))
}

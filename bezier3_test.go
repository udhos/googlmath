package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_V3s_Pth struct {
	Points   []Vector3
	Expected Path3
}

func (s *S) TestBezier3Bzr3(c *C) {
	tests := []TestValue_V3s_Pth{
		TestValue_V3s_Pth{
			Expected: Bezier3(nil),
		},
		TestValue_V3s_Pth{
			Points:   []Vector3{Vec3(0, 0, 0)},
			Expected: Bezier3{Vec3(0, 0, 0)},
		},
		TestValue_V3s_Pth{
			Points:   []Vector3{Vec3(0, 0, 0), Vec3(-1, -1, -1), Vec3(2, 3, 4)},
			Expected: Bezier3{Vec3(0, 0, 0), Vec3(-1, -1, -1), Vec3(2, 3, 4)},
		},
	}
	for _, t := range tests {
		obtained := Bzr3(t.Points...)
		c.Check(obtained, DeepEquals, t.Expected)
	}
}

type TestValue_Pth3F32_V3 struct {
	p Path3
	f float32
	e Vector3
}

func (s *S) TestBezier3ValueAt(c *C) {
	tests := []TestValue_Pth3F32_V3{
		{Bezier3(nil), 0.5, Vec3(NaN(), NaN(), NaN())},
		{Bezier3{Vec3(0, 0, 0), Vec3(1, 1, 1)}, 0.5, Vec3(0.5, 0.5, 0.5)},
		{Bezier3{Vec3(0, 0, 0), Vec3(-1, -1, -1), Vec3(2, 2, 2)}, 0.5, Vec3(0, 0, 0)},
		{Bezier3{Vec3(0, 0, 0), Vec3(1, 1, 1), Vec3(2, 2, 2), Vec3(3, 3, 3)}, 0.5, Vec3(1.5, 1.5, 1.5)},
	}
	for _, t := range tests {
		obtained := t.p.ValueAt(t.f)
		c.Check(obtained, Vector3Check, t.e)
	}
}

type TestValue_Pth3V3_F32 struct {
	p Path3
	v Vector3
	e float32
}

func (s *S) TestBezier3Approximate(c *C) {
	tests := []TestValue_Pth3V3_F32{
		{Bezier3(nil), Vec3(0, 0, 0), NaN()},
		{Bezier3{Vec3(0, 0, 0), Vec3(1, 1, 1)}, Vec3(0.5, 0.5, 0.5), 0.5},
		{Bezier3{Vec3(0, 0, 0), Vec3(-1, -1, -1), Vec3(2, 2, 2)}, Vec3(0, 0, 0), 0},
	}
	for _, t := range tests {
		obtained := t.p.Approximate(t.v)
		c.Check(obtained, EqualsFloat32, t.e)
	}
}

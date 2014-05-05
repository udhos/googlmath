package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_V2s_Pth2 struct {
	pts []Vector2
	e   Path2
}

func (s *S) TestBezier2Bzr2(c *C) {
	tests := []TestValue_V2s_Pth2{
		{nil, Bezier2(nil)},
		{[]Vector2{Vec2(0, 0)}, Bezier2{Vec2(0, 0)}},
		{[]Vector2{Vec2(0, 0), Vec2(-1, -1), Vec2(2, 3)}, Bezier2{Vec2(0, 0), Vec2(-1, -1), Vec2(2, 3)}},
	}
	for _, t := range tests {
		obtained := Bzr2(t.pts...)
		c.Check(obtained, DeepEquals, t.e)
	}
}

type TestValue_PthF32_V2 struct {
	p Path2
	f float32
	e Vector2
}

func (s *S) TestBezier2ValueAt(c *C) {
	tests := []TestValue_PthF32_V2{
		{Bezier2(nil), 0.5, Vec2(0, 0)},
		{Bezier2{Vec2(0, 0), Vec2(1, 1)}, 0.5, Vec2(0.5, 0.5)},
		{Bezier2{Vec2(0, 0), Vec2(-1, -1), Vec2(2, 2)}, 0.5, Vec2(0, 0)},
	}
	for _, t := range tests {
		obtained := t.p.ValueAt(t.f)
		c.Check(obtained, DeepEquals, t.e)
	}
}

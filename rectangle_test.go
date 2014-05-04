package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_4F32_Rect struct {
	f0, f1, f2, f3 float32
	r              Rectangle
}

func (s *S) TestRectangleRect(c *C) {
	tests := []TestValue_4F32_Rect{
		{1.23, -3.21, 0.2, 0, Rectangle{Min: Vector2{1.23, -3.21}, Max: Vector2{0.2, 0}}},
		{0.23, 3.21, 1.2, -0.2, Rectangle{Min: Vector2{0.23, 3.21}, Max: Vector2{1.2, -0.2}}},
	}
	for _, t := range tests {
		actual := Rect(t.f0, t.f1, t.f2, t.f3)
		c.Assert(actual, Equals, t.r)
	}
}

type TestValue_Rect_F32 struct {
	r Rectangle
	f float32
}

func (s *S) TestRectangleDx(c *C) {
	tests := []TestValue_Rect_F32{
		{Rectangle{Min: Vector2{1.23, -3.21}, Max: Vector2{0.2, 0}}, -1.03},
		{Rectangle{Min: Vector2{0.23, 3.21}, Max: Vector2{1.2, -0.2}}, 0.97},
	}
	for _, t := range tests {
		actual := t.r.Dx()
		c.Assert(actual, Equals, t.f)
	}
}

func (s *S) TestRectangleDy(c *C) {
	tests := []TestValue_Rect_F32{
		{Rectangle{Min: Vector2{1.23, -3.21}, Max: Vector2{0.2, 0}}, 3.21},
		{Rectangle{Min: Vector2{0.23, 3.21}, Max: Vector2{1.2, -0.2}}, -3.41},
	}
	for _, t := range tests {
		actual := t.r.Dy()
		c.Assert(actual, Equals, t.f)
	}
}

type TestValue_2Rect_B struct {
	r0 Rectangle
	r1 Rectangle
	e  bool
}

func (s *S) TestRectangleIn(c *C) {
	tests := []TestValue_2Rect_B{
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, true},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{-2.2, 0}, Max: Vector2{1, 3.3}}, true},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{-0.5, 0}, Max: Vector2{0.5, 1}}, false},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{0, 1}, Max: Vector2{2, 2}}, false},
		{Rectangle{Min: Vector2{-2.2, -3.3}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{0, 1}, Max: Vector2{2, 0.5}}, false},
	}
	for _, t := range tests {
		actual := t.r0.In(t.r1)
		c.Assert(actual, Equals, t.e, Commentf("%v.In(%v)", t.r0, t.r1))
	}
}

func (s *S) TestRectangleOverlaps(c *C) {
	tests := []TestValue_2Rect_B{
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, true},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{-2.2, 0}, Max: Vector2{1, 3.3}}, true},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{-0.5, 0}, Max: Vector2{0.5, 1}}, true},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{1.1, 1.01}, Max: Vector2{2, 2}}, false},
	}
	for _, t := range tests {
		actual := t.r0.Overlaps(t.r1)
		c.Assert(actual, Equals, t.e, Commentf("%v.Overlaps(%v)", t.r0, t.r1))
	}
}

type TestValue_RectV2_B struct {
	r Rectangle
	v Vector2
	e bool
}

func (s *S) TestRectangleContains(c *C) {
	tests := []TestValue_RectV2_B{
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Vector2{0, 0}, true},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Vector2{0.5, 0}, true},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}, Vector2{-0.2, 1.0}, true},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Vector2{-1, 0}, false},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Vector2{-1.2, -2}, false},
	}
	for _, t := range tests {
		actual := t.r.Contains(t.v)
		c.Assert(actual, Equals, t.e, Commentf("%v.Contains(%v)", t.r, t.v))
	}
}

type TestValue_Rect_B struct {
	r Rectangle
	e bool
}

func (s *S) TestRectangleEmpty(c *C) {
	tests := []TestValue_Rect_B{
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, false},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}, false},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{-1, -1}}, true},
		{Rectangle{Min: Vector2{1, 1}, Max: Vector2{-1, -1}}, true},
	}
	for _, t := range tests {
		actual := t.r.Empty()
		c.Assert(actual, Equals, t.e, Commentf("%v.Empty()", t.r))
	}
}

type TestValue_2Rect_Rect struct {
	r0 Rectangle
	r1 Rectangle
	e  Rectangle
}

func (s *S) TestRectangleMerge(c *C) {
	tests := []TestValue_2Rect_Rect{
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{-2.2, 0}, Max: Vector2{1, 3.3}}, Rectangle{Min: Vector2{-2.2, 0}, Max: Vector2{1, 3.3}}},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{-0.5, 0}, Max: Vector2{0.5, 1}}, Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}},
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Rectangle{Min: Vector2{1.1, 1.01}, Max: Vector2{22.1, 21.2}}, Rectangle{Min: Vector2{0, 0}, Max: Vector2{22.1, 21.2}}},
	}
	for _, t := range tests {
		actual := t.r0.Merge(t.r1)
		c.Assert(actual, Equals, t.e, Commentf("%v.Merge(%v)", t.r0, t.r1))
	}
}

type TestValue_Rect_V2 struct {
	r Rectangle
	e Vector2
}

func (s *S) TestRectangleCenter(c *C) {
	tests := []TestValue_Rect_V2{
		{Rectangle{Min: Vector2{0, 0}, Max: Vector2{1, 1}}, Vector2{0.5, 0.5}},
		{Rectangle{Min: Vector2{-1, -1}, Max: Vector2{1, 1}}, Vector2{0, 0}},
		{Rectangle{Min: Vector2{-1.2, 0}, Max: Vector2{-1, -1}}, Vector2{-1.1, -0.5}},
		{Rectangle{Min: Vector2{1, 1}, Max: Vector2{2.2, -1}}, Vector2{1.6, 0}},
	}
	for _, t := range tests {
		actual := t.r.Center()
		c.Assert(actual, Equals, t.e, Commentf("%v.Center()", t.r))
	}
}

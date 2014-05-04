package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_V2s_Pth struct {
	Points   []Vector2
	Expected Path2
}

func (s *S) TestNewBezier2(c *C) {
	tests := []TestValue_V2s_Pth{
		TestValue_V2s_Pth{
			Expected: NewBezier2(),
		},
		TestValue_V2s_Pth{
			Points:   []Vector2{Vec2(0, 0)},
			Expected: NewBezier2(Vec2(0, 0)),
		},
		TestValue_V2s_Pth{
			Points:   []Vector2{Vec2(0, 0), Vec2(-1, -1), Vec2(2, 3)},
			Expected: NewBezier2(Vec2(0, 0), Vec2(-1, -1), Vec2(2, 3)),
		},
	}
	for _, t := range tests {
		obtained := NewBezier2(t.Points...)
		c.Check(obtained, DeepEquals, t.Expected)
	}
}

package math

import (
	. "gopkg.in/check.v1"
)

type Bezier2Points struct {
	Points   []Vector2
	Expected Path2
}

type Bezier2TestSuite struct {
	setTestTable []*Bezier2Points
}

var _ = Suite(&Bezier2TestSuite{})

func (s *Bezier2TestSuite) SetUpTest(c *C) {
	s.setTestTable = []*Bezier2Points{
		&Bezier2Points{
			Expected: NewBezier2(),
		},
		&Bezier2Points{
			Points:   []Vector2{Vec2(0, 0)},
			Expected: NewBezier2(Vec2(0, 0)),
		},
		&Bezier2Points{
			Points:   []Vector2{Vec2(0, 0), Vec2(-1, -1), Vec2(2, 3)},
			Expected: NewBezier2(Vec2(0, 0), Vec2(-1, -1), Vec2(2, 3)),
		},
	}
}

func (s *Bezier2TestSuite) TestNew(c *C) {
	for _, t := range s.setTestTable {
		obtained := NewBezier2(t.Points...)
		c.Check(obtained, DeepEquals, t.Expected)
	}
}

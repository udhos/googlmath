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

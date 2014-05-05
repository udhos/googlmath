package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_2V3_Sgt struct {
	v0 Vector3
	v1 Vector3
	e  Segment
}

func (s *S) TestSegmentSgt(c *C) {
	tests := []TestValue_2V3_Sgt{
		{Vec3(0, 0, 0), Vec3(0, 0, 0), Segment{Vec3(0, 0, 0), Vec3(0, 0, 0)}},
		{Vec3(1, 2, 3), Vec3(3, 2, 1), Segment{Vec3(1, 2, 3), Vec3(3, 2, 1)}},
		{Vec3(-2.2, -3.3, -3.4), Vec3(-2, -1, -404), Segment{Vec3(-2.2, -3.3, -3.4), Vec3(-2, -1, -404)}},
		{Vec3(-2, 1, -3), Vec3(2, 1, -0.5), Segment{Vec3(-2, 1, -3), Vec3(2, 1, -0.5)}},
	}
	for _, t := range tests {
		actual := Sgt(t.v0, t.v1)
		c.Assert(actual, Equals, t.e)
	}
}

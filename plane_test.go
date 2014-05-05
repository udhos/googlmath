package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_V3F32_Ple struct {
	v Vector3
	f float32
	e Plane
}

func (s *S) TestPlanePle(c *C) {
	tests := []TestValue_V3F32_Ple{}
	for _, t := range tests {
		actual := Ple(t.v, t.f)
		c.Assert(actual, Equals, t.e)
	}
}

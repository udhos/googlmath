package math

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *S) TestQuaternionVec4(c *C) {
	tests := []struct{ X, Y, Z, W float32 }{
		{0, 0, 0, 0},
		{2, 0, 1, 0},
		{0, 4, 0, 0},
		{2, 5, 0, 0.1},
		{-3, 0, 0, 0.5},
		{0, -5, 0, 2.3},
		{-1, -22, -3.3, 0},
		{-3.3, 5.1, 1.2, 1.1},
	}
	for _, t := range tests {
		actual := Qtn(t.X, t.Y, t.Z, t.W)
		expected := Quaternion{t.X, t.Y, t.Z, t.W}
		c.Assert(actual, QuaternionCheck, expected)
	}
}

type TestValue_QF32_Q struct {
	v Quaternion
	f float32
	e Quaternion
}

func (s *S) TestQuaternionScale(c *C) {
	tests := []TestValue_QF32_Q{
		{Quaternion{2, 5, 0, 0.1}, 1.0, Quaternion{2, 5, 0, 0.1}},
		{Quaternion{2, 5, 0, 0.1}, 0, Quaternion{0, 0, 0, 0}},
		{Quaternion{-1, 5, 0, 0.1}, -1.0, Quaternion{1, -5, 0, -0.1}},
		{Quaternion{2, 5, 0, 0.1}, 1.2, Quaternion{2.4, 6, 0, 0.12}},
	}
	for _, t := range tests {
		actual := t.v.Scale(t.f)
		c.Assert(actual, QuaternionCheck, t.e)
	}
}

// TODO Dot

type TestValue_Q_F32 struct {
	v Quaternion
	e float32
}

func (s *S) TestQuaternionLen(c *C) {
	tests := []TestValue_Q_F32{
		{Quaternion{2, 0, 0, 0}, 2.0},
		{Quaternion{0, 1, 0, 0}, 1.0},
		{Quaternion{0, 0, 0, 0}, 0},
		{Quaternion{-1, 0, 0, 0}, 1.0},
		{Quaternion{2, 2, 0, 0}, 2.828427},
	}
	for _, t := range tests {
		actual := t.v.Len()
		c.Assert(actual, Equals, t.e)
	}
}

func (s *S) TestQuaternionLen2(c *C) {
	tests := []TestValue_Q_F32{
		{Quaternion{2, 0, 0, 0}, 4.0},
		{Quaternion{0, 1, 0, 0}, 1.0},
		{Quaternion{0, 0, 0, 0}, 0},
		{Quaternion{-1, 0, 0, 0}, 1.0},
		{Quaternion{2, 2, 0, 0}, 8.0},
	}
	for _, t := range tests {
		actual := t.v.Len2()
		c.Assert(actual, Equals, t.e)
	}
}

func (s *S) TestQuaternionNor(c *C) {
	tests := []struct{ v, e Quaternion }{
		{Quaternion{0, 0, 0, 0}, Quaternion{0, 0, 0, 0}},
		{Quaternion{1, -1, 1, 1}, Quaternion{0.5, -0.5, 0.5, 0.5}},
		{Quaternion{2, 0, 1, 1}, Quaternion{0.816497, 0.0, 0.408248, 0.408248}},
	}
	for _, t := range tests {
		actual := t.v.Nor()
		c.Assert(actual, QuaternionCheck, t.e)
	}
}

func (s *S) TestQuaternionEq(c *C) {
	tests := []Quaternion{
		Quaternion{0, 0, 0, 0}, Quaternion{0, 0, 0, 0},
		Quaternion{1, -1, 1, 1}, Quaternion{-1, 1, -1, -1},
		Quaternion{2, 0, 1, 1}, Quaternion{-2, 0, -1, -1},
		Quaternion{0.816497, 0.0, 0.408248, 0.408248}, Quaternion{-0.816497, 0.0, -0.408248, -0.408248},
	}
	for i := range tests {
		v0 := tests[i]
		for j := range tests {
			v1 := tests[j]
			e := v0.Eq(v1)
			actual := v0.X == v1.X && v0.Y == v1.Y && v0.Z == v1.Z && v0.W == v1.W
			c.Assert(actual, Equals, e)
		}
	}
}

// TODO EulerAngles
// TODO Conjugate
// TODO FromAxis
// TODO FromMatrix
// TODO FromAxes
// TODO FromCross
// TODO Slerp
// TODO Matrix

func (s *S) TestQuaternionString(c *C) {
	tests := []Quaternion{
		Quaternion{0, 0, 0, 0}, Quaternion{0, 0, 0, 0},
		Quaternion{1, -1, 1, 1}, Quaternion{-1, 1, -1, -1},
		Quaternion{2, 0, 1, 1}, Quaternion{-2, 0, -1, -1},
		Quaternion{0.816497, 0.0, 0.408248, 0.408248}, Quaternion{-0.816497, 0.0, -0.408248, -0.408248},
	}
	for _, v := range tests {
		actual := v.String()
		c.Assert(actual, Equals, fmt.Sprintf("%v", v))
	}
}

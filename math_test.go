package math

import (
	. "gopkg.in/check.v1"
	"math/rand"
)

func (s *S) TestNextPowerOfTwo(c *C) {
	for _, t := range []struct{ i, e int }{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 8},
		{6, 8},
		{7, 8},
		{8, 8},
		{9, 16},
		{10, 16},
	} {
		obtained := NextPowerOfTwo(t.i)
		c.Assert(obtained, Equals, t.e, Commentf("NextPowerOfTwo(%v)", t.i))
	}
}

func (s *S) TestIsPowerOfTwo(c *C) {
	for i := 0; i < 100; i++ {
		n := rand.Int()
		expected := n&(n-1) == 0 && n != 0
		obtained := IsPowerOfTwo(n)
		c.Assert(obtained, Equals, expected, Commentf("IsPowerOfTwo(%v)", n))
	}
	for _, t := range []int{
		1, 2, 4, 8, 16, 2147483648,
	} {
		c.Assert(IsPowerOfTwo(t), Equals, true, Commentf("IsPowerOfTwo(%v)", t))
	}
}

func (s *S) TestDegreeRadiansConversion(c *C) {
	for i := 0; i < 100; i++ {
		degree := rand.Float32()
		radian := ToRadians(degree)
		degree2 := ToDegrees(radian)
		c.Assert(degree2, EqualsFloat32, degree)
	}
	for _, t := range []struct{ f, e float32 }{
		{12.523, 0.2186},
		{-2452.2354235, -42.79958217},
	} {
		obtained := ToRadians(t.f)
		c.Assert(obtained, EqualsFloat32, t.e)
	}
}

type TestValue_3I_I struct {
	i0 int
	i1 int
	i2 int
	e  int
}

func (s *S) TestClampi(c *C) {
	tests := []TestValue_3I_I{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 1, 0, 1}, // Clampf does no check against min<max == true
		{0, 0, 1, 0},
		{-1, 1, -1, 1},
		{-2, 32, -2, 32},
		{1, 322, 1, 322},
	}
	for _, t := range tests {
		obtained := Clampi(t.i0, t.i1, t.i2)
		c.Assert(obtained, Equals, t.e, Commentf("Clampf(%v,%v,%v)", t.i0, t.i1, t.i2))
	}
}

type TestValue_3F32_F32 struct {
	f0 float32
	f1 float32
	f2 float32
	e  float32
}

func (s *S) TestClampf(c *C) {
	tests := []TestValue_3F32_F32{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 1, 0, 1}, // Clampf does no check against min<max == true
		{0, 0, 1, 0},
		{-1, 1, -1, 1},
		{-2.4, 32.3, -2.4, 32.3},
		{1, 322.2, 1, 322.2},
	}
	for _, t := range tests {
		obtained := Clampf(t.f0, t.f1, t.f2)
		c.Assert(obtained, Equals, t.e, Commentf("Clampf(%v,%v,%v)", t.f0, t.f1, t.f2))
	}
}

type TestValue_2F32_F32 struct {
	f0 float32
	f1 float32
	e  float32
}

func (s *S) TestMin(c *C) {
	tests := []TestValue_2F32_F32{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{-1, 1, -1},
		{-2.4, 32.3, -2.4},
		{1, 322.2, 1},
	}
	for _, t := range tests {
		obtained := Min(t.f0, t.f1)
		c.Assert(obtained, Equals, t.e)
	}
}

func (s *S) TestMax(c *C) {
	tests := []TestValue_2F32_F32{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{-1, 1, 1},
		{-2.4, 32.3, 32.3},
		{1, 322.2, 322.2},
	}
	for _, t := range tests {
		obtained := Max(t.f0, t.f1)
		c.Assert(obtained, Equals, t.e)
	}
}

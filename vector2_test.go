package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestVec2(c *C) {
	vec := Vec2(1.23, -3.21)
	c.Assert(vec, Equals, Vector2{1.23, -3.21})
}

func (s *S) TestVector2Len(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Len(), Equals, float32(2))
}

func (s *S) TestVector2Len2(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Len2(), Equals, float32(4))
}

func (s *S) TestVector2Sub(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Sub(Vec2(-2, -1)), Equals, Vec2(0, 1))
}

func (s *S) TestVector2Nor(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Nor(), Equals, Vec2(-1, 0))
}

func (s *S) TestVector2Add(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Add(Vec2(2, 2)), Equals, Vec2(0, 2))
}

func (s *S) TestVector2Dot(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Dot(Vec2(-3, 1)), Equals, float32(6))
}

func (s *S) TestVector2Mul(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Mul(Vec2(-2, 2)), Equals, Vec2(4, 0))
}

func (s *S) TestVector2Div(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Div(Vec2(-2, 2)), Equals, Vec2(1, 0))
}

func (s *S) TestVector2Scale(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Scale(2), Equals, Vec2(-4, 0))
}

func (s *S) TestVector2Distance(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Distance(Vec2(0, 0)), Equals, float32(2))
}

func (s *S) TestVector2Distance2(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Distance2(Vec2(0, 0)), Equals, float32(4))
}

func (s *S) TestVector2Limit(c *C) {
	vec := Vec2(-2, 0)
	c.Assert(vec.Limit(1), Equals, Vec2(-1, 0))
}

// TODO MulMatrix

func (s *S) TestVector2Cross(c *C) {
	vec := Vec2(5, 1)
	c.Assert(vec.Cross(Vec2(-1, 0)), Equals, float32(1.0))
}

func (s *S) TestVector2Angle(c *C) {
	vec := Vec2(1, 1)
	c.Assert(vec.Angle(), Equals, float32(45.0))
}

func (s *S) TestVector2Rotate(c *C) {
	var angle float32 = 45
	var x float32 = 5
	var y float32 = -2
	xResult := Cos(angle*DegreeToRadians)*x - Sin(angle*DegreeToRadians)*y
	yResult := Sin(angle*DegreeToRadians)*x + Cos(angle*DegreeToRadians)*y
	vec := Vec2(x, y)

	vec = vec.Rotate(angle)
	c.Assert(vec.X, Equals, xResult)
	c.Assert(vec.Y, Equals, yResult)
}

func (s *S) TestVector2Lerp(c *C) {
	var alpha float32 = 0.5
	v1 := Vec2(1, 1)
	v2 := Vec2(-2, 0)
	v3 := v1.Lerp(v2, alpha)

	xResult := v1.X*(1-alpha) + v2.X*alpha
	yResult := v1.Y*(1-alpha) + v2.Y*alpha

	c.Assert(v3.X, Equals, xResult)
	c.Assert(v3.Y, Equals, yResult)
}

func (s *S) TestVector2Faceforward(c *C) {
	v := Vec2(1.0, -2.0)
	n := Vec2(0.0, 0.0)
	i := Vec2(2.2, 0.3)

	expected := Vec2(-1.0, 2.0)
	result := v.Faceforward(i, n)
	c.Check(result, Equals, expected)
}

// ### Benchmarks ###

func (s *S) TestBenchmarkVector2Add(c *C) {
	vec1 := Vec2(0, 0)
	vec2 := Vec2(1, 1)
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		vec1.Add(vec2)
	}
}

// Benchmarking pointer vs no pointer for Vector2
// NOTE: A pointer uses 64bit on a 64bit system.
// NoPointer is faster since less code is copied into the function (only the vector2 value)
// but Pointer requires less memory usage.

type Vector2NoPointer struct {
	X, Y float32
}

func (vec Vector2NoPointer) Add(vec2 Vector2NoPointer) Vector2NoPointer {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func (s *S) TestBenchmarkVector2NoPointerAdd(c *C) {
	vec1 := Vector2NoPointer{0, 0}
	vec2 := Vector2NoPointer{1, 1}
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		vec1.Add(vec2)
	}
}

type Vector2Pointer struct {
	X, Y float32
}

func (vec *Vector2Pointer) Add(vec2 *Vector2Pointer) *Vector2Pointer {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func (s *S) TestBenchmarkVector2PointerAdd(c *C) {
	vec1 := &Vector2Pointer{0, 0}
	vec2 := &Vector2Pointer{1, 1}
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		vec1.Add(vec2)
	}
}

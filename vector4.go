package math

import "fmt"

var ZV4 Vector4

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func Vec4(x, y, z, w float32) Vector4 {
	return Vector4{x, y, z, w}
}

func (vec Vector4) Add(vec2 Vector4) Vector4 {
	return Vector4{vec.X + vec2.X, vec.Y + vec2.Y, vec.Z + vec2.Z, vec.W + vec2.W}
}

func (vec Vector4) Sub(vec2 Vector4) Vector4 {
	return Vector4{vec.X - vec2.X, vec.Y - vec2.Y, vec.Z - vec2.Z, vec.W - vec2.W}
}

func (vec Vector4) Mul(vec2 Vector4) Vector4 {
	return Vector4{vec.X * vec2.X, vec.Y * vec2.Y, vec.Z * vec2.Z, vec.W * vec2.W}
}

func (vec Vector4) Div(vec2 Vector4) Vector4 {
	return Vector4{vec.X / vec2.X, vec.Y / vec2.Y, vec.Z / vec2.Z, vec.W / vec2.W}
}

func (vec Vector4) Scale(scalar float32) Vector4 {
	return Vector4{vec.X * scalar, vec.Y * scalar, vec.Z * scalar, vec.W * scalar}
}

// The euclidian length
func (vec Vector4) Len() float32 {
	return Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z + vec.W*vec.W)
}

// The squared euclidian length
func (vec Vector4) Len2() float32 {
	return vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z + vec.W*vec.W
}

func (vec Vector4) Nor() Vector4 {
	l := vec.Len()
	if l == 0 {
		return vec
	}
	return vec.Scale(1 / l)
}

func (vec Vector4) IsZero() bool {
	return vec.X == 0 && vec.Y == 0 && vec.Z == 0 && vec.W == 0
}

// Whether this vector is a unit length vector
func (vec Vector4) IsUnit() bool {
	return vec.Len() == 1
}

func (vec Vector4) Invert() Vector4 {
	return Vector4{-vec.X, -vec.Y, -vec.Z, -vec.W}
}

func (vec Vector4) Eq(v2 Vector4) bool {
	return vec.X == v2.X && vec.Y == v2.Y && vec.Z == v2.Z && vec.W == v2.W
}

func (vec Vector4) String() string {
	return fmt.Sprintf("(%g,%g,%g,%g)", vec.X, vec.Y, vec.Z, vec.W)
}

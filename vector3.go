package math

import "fmt"

var ZV3 Vector3

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func Vec3(x, y, z float32) Vector3 {
	return Vector3{x, y, z}
}

func (vec Vector3) Add(vec2 Vector3) Vector3 {
	return Vec3(vec.X+vec2.X, vec.Y+vec2.Y, vec.Z+vec2.Z)
}

func (vec Vector3) Sub(vec2 Vector3) Vector3 {
	return Vec3(vec.X-vec2.X, vec.Y-vec2.Y, vec.Z-vec2.Z)
}

func (vec Vector3) Mul(vec2 Vector3) Vector3 {
	return Vec3(vec.X*vec2.X, vec.Y*vec2.Y, vec.Z*vec2.Z)
}

func (vec Vector3) Div(vec2 Vector3) Vector3 {
	return Vec3(vec.X/vec2.X, vec.Y/vec2.Y, vec.Z/vec2.Z)
}

// The euclidian length
func (vec Vector3) Len() float32 {
	return Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
}

// The squared euclidian length
func (vec Vector3) Len2() float32 {
	return vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z
}

func (vec Vector3) Distance(vec2 Vector3) float32 {
	return Sqrt(vec.Distance2(vec2))
}

// Returns the squared distance between this point and the given point
func (vec Vector3) Distance2(vec2 Vector3) float32 {
	a := vec2.X - vec.X
	b := vec2.Y - vec.Y
	c := vec2.Z - vec.Z
	a *= a
	b *= b
	c *= c
	return a + b + c
}

func (vec Vector3) Nor() Vector3 {
	l := vec.Len()
	if l == 0 {
		return vec
	}
	return vec.Scale(1 / l)
}

func (vec Vector3) Dot(vec2 Vector3) float32 {
	return vec.X*vec2.X + vec.Y*vec2.Y + vec.Z*vec2.Z
}

// Returns the cross product between this vector and the other vector
func (vec Vector3) Cross(vec2 Vector3) Vector3 {
	x := vec.Y*vec2.Z - vec.Z*vec2.Y
	y := vec.Z*vec2.X - vec.X*vec2.Z
	z := vec.X*vec2.Y - vec.Y*vec2.X
	return Vec3(x, y, z)
}

// Whether this vector is a unit length vector
func (vec Vector3) IsUnit() bool {
	return vec.Len() == 1
}

func (vec Vector3) IsZero() bool {
	return vec.X == 0 && vec.Y == 0 && vec.Z == 0
}

// Linearly interpolates between this vector and the target vector by alpha which is in the range [0:1].
func (vec Vector3) Lerp(target Vector3, alpha float32) Vector3 {
	if alpha > 1.0 || alpha < 0 {
		panic("alpha is not in the range [0:1]")
	}
	k := alpha
	if vec.Dot(target) < 0 {
		k = -alpha
	}
	v := vec.Scale(1.0 - alpha)
	return v.Add(target.Scale(k))
}

// Spherically interpolates between this vector and the target vector by alpha which is in the range [0,1].
func (vec Vector3) Slerp(target Vector3, alpha float32) Vector3 {
	dot := vec.Dot(target)
	if dot > 0.99995 || dot < 0.9995 {
		vec = vec.Add(target.Sub(vec).Scale(alpha))
		vec = vec.Nor()
		return vec
	}

	if dot > 1 {
		dot = 1
	}
	if dot < -1 {
		dot = -1
	}

	theta0 := Acos(dot)
	theta := theta0 * alpha
	v2 := target.Sub(Vec3(vec.X*dot, vec.Y*dot, vec.Z*dot))
	v2 = v2.Nor()
	return vec.Scale(Cos(theta)).Add(v2.Scale(Sin(theta))).Nor()
}

// Returns this vector, it's length limited to given value.
func (vec Vector3) Limit(limit float32) Vector3 {
	if vec.Len2() > limit*limit {
		vec = vec.Nor()
		vec = vec.Scale(limit)
	}
	return vec
}

func (vec Vector3) Scale(scalar float32) Vector3 {
	return Vec3(vec.X*scalar, vec.Y*scalar, vec.Z*scalar)
}

func (vec Vector3) Invert() Vector3 {
	return Vec3(-vec.X, -vec.Y, -vec.Z)
}

func (vec Vector3) Eq(v2 Vector3) bool {
	return vec.X == v2.X && vec.Y == v2.Y && vec.Z == v2.Z
}

func (vec Vector3) String() string {
	return fmt.Sprintf("(%g,%g,%g)", vec.X, vec.Y, vec.Z)
}

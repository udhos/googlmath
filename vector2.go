package math

import "fmt"

var ZV2 Vector2

type Vector2 struct {
	X float32
	Y float32
}

func Vec2(x, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func (vec Vector2) Add(vec2 Vector2) Vector2 {
	return Vector2{vec.X + vec2.X, vec.Y + vec2.Y}
}

func (vec Vector2) Sub(vec2 Vector2) Vector2 {
	return Vector2{vec.X - vec2.X, vec.Y - vec2.Y}
}

func (vec Vector2) Mul(vec2 Vector2) Vector2 {
	return Vector2{vec.X * vec2.X, vec.Y * vec2.Y}
}

func (vec Vector2) Div(vec2 Vector2) Vector2 {
	return Vector2{vec.X / vec2.X, vec.Y / vec2.Y}
}

func (vec Vector2) Scale(scale float32) Vector2 {
	return Vector2{vec.X * scale, vec.Y * scale}
}

// Returns the normalized vector
func (vec Vector2) Nor() Vector2 {
	l := vec.Len()
	if l != 0 {
		return Vector2{vec.X / l, vec.Y / l}
	}
	return Vector2{vec.X, vec.Y}
}

// The euclidian length
func (vec Vector2) Len() float32 {
	return Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

// The squared euclidian length
func (vec Vector2) Len2() float32 {
	return vec.X*vec.X + vec.Y*vec.Y
}

// Dot returns the dot product of this vector and the given vector.
func (vec Vector2) Dot(vec2 Vector2) float32 {
	return vec.X*vec2.X + vec.Y*vec2.Y
}

// Distance returns the distance between this and the given vector.
func (vec Vector2) Distance(vec2 Vector2) float32 {
	xd := vec2.X - vec.X
	yd := vec2.Y - vec.Y
	return Sqrt(xd*xd + yd*yd)
}

// Distance2 returns the squared distance between this and the given vector.
func (vec Vector2) Distance2(vec2 Vector2) float32 {
	xd := vec2.X - vec.X
	yd := vec2.Y - vec.Y
	return xd*xd + yd*yd
}

// Returns a vector limited to given value based on this vector
func (vec Vector2) Limit(limit float32) Vector2 {
	if vec.Len2() > limit*limit {
		return vec.Nor().Scale(limit)
	}
	return Vector2{vec.X, vec.Y}
}

func (vec Vector2) MulMatrix(m *Matrix3) Vector2 {
	x := vec.X*m.M11 + vec.Y*m.M21 + m.M31
	y := vec.X*m.M12 + vec.Y*m.M22 + m.M32
	return Vector2{x, y}
}

// Cross returns the cross product of this vector ang the given vector.
func (vec Vector2) Cross(vec2 Vector2) float32 {
	return vec.X*vec2.Y - vec.Y*vec2.X
}

func (vec Vector2) Angle() float32 {
	angle := Atan2(vec.Y, vec.X) * RadiansToDegrees
	if angle < 0 {
		angle += 360
	}
	return angle
}

func (vec *Vector2) SetAngle(angle float32) Vector2 {
	return (Vector2{vec.Len(), 0}).Rotate(angle)
}

// Returns the rotated Vector2 by the given angle, counter-clockwise.
func (vec Vector2) Rotate(degrees float32) Vector2 {
	rad := degrees * DegreeToRadians
	cos := Cos(rad)
	sin := Sin(rad)
	return Vector2{vec.X*cos - vec.Y*sin, vec.X*sin + vec.Y*cos}
}

// Lerp returns the linearly interpolates between this vector and the target vector by alpha which is in the range [0,1].
func (vec Vector2) Lerp(target Vector2, alpha float32) Vector2 {
	invAlpha := 1.0 - alpha
	return Vector2{vec.X*invAlpha + target.X*alpha, vec.Y*invAlpha + target.Y*alpha}
}

// Faceforward returns this vector if n,Dot(i) < 0, otherwise, returns the negative of this vector.
func (vec Vector2) Faceforward(i, n Vector2) Vector2 {
	if n.Dot(i) < 0 {
		return Vector2{vec.X, vec.Y}
	}
	return Vector2{-vec.X, -vec.Y}
}

// Whether this vector is a unit length vector
func (vec Vector2) IsUnit() bool {
	return vec.Len() == 1
}

func (vec Vector2) IsZero() bool {
	return vec.X == 0 && vec.Y == 0
}

func (vec Vector2) String() string {
	return fmt.Sprintf("(%g,%g)", vec.X, vec.Y)
}

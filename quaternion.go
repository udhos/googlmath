package math

import (
	"math"
)

var IdtQ Quaternion = Quaternion(ZV4)

type Quaternion Vector4

func Qtn(x, y, z, w float32) Quaternion {
	return Quaternion{x, y, z, w}
}

func (q Quaternion) Scale(scalar float32) Quaternion {
	return Quaternion{q.X * scalar, q.Y * scalar, q.Z * scalar, q.W * scalar}
}

func (q Quaternion) Dot(q2 Quaternion) float32 {
	return q.X*q2.X + q.Y*q2.Y + q.Z*q2.Z + q.W*q2.W
}

// The euclidian length
func (q Quaternion) Len() float32 {
	return Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

// The squared euclidian length
func (q Quaternion) Len2() float32 {
	return q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W
}

func (q Quaternion) Nor() Quaternion {
	l := q.Len()
	if l == 0 {
		return q
	}
	return q.Scale(1 / l)
}

func (q Quaternion) Eq(q2 Quaternion) bool {
	return q.X == q2.X && q.Y == q2.Y && q.Z == q2.Z && q.W == q2.W
}

// Returns the quaternion to the given euler angles.
// Values in radians
func (q Quaternion) EulerAngles(yaw, pitch, roll float32) Quaternion {
	num9 := roll * 0.5
	num6 := Sin(num9)
	num5 := Cos(num9)
	num8 := pitch * 0.5
	num4 := Sin(num8)
	num3 := Cos(num8)
	num7 := yaw * 0.5
	num2 := Sin(num7)
	num := Cos(num7)
	f1 := num * num4
	f2 := num2 * num3
	f3 := num * num3
	f4 := num2 * num4

	q.X = (f1 * num5) + (f2 * num6)
	q.Y = (f2 * num5) - (f1 * num6)
	q.Z = (f3 * num6) - (f4 * num5)
	q.W = (f3 * num5) + (f4 * num6)
	return q
}

// Conjugate the quaternion.
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{-q.X, -q.Y, -q.Z, q.W}
}

// Sets the quaternion components from the given axis and angle around that axis.
// Angle in radians
func (q Quaternion) FromAxis(x, y, z, angle float32) Quaternion {
	lSin := Sin(angle / 2)
	lCos := Cos(angle / 2)
	return Quaternion{q.X * lSin, q.Y * lSin, q.Z * lSin, lCos}.Nor()
}

func (q Quaternion) FromMatrix(m *Matrix4) Quaternion {
	return q.FromAxes(m.M11, m.M12, m.M13, m.M21, m.M22, m.M23, m.M31, m.M32, m.M33)
}

// Sets the Quaternion from the given x-, y- and z-axis which have to be orthonormal.
func (q Quaternion) FromAxes(xx, xy, xz, yx, yy, yz, zx, zy, zz float32) Quaternion {
	m00 := float64(xx)
	m01 := float64(xy)
	m02 := float64(xz)

	m10 := float64(yx)
	m11 := float64(yy)
	m12 := float64(yz)

	m20 := float64(zx)
	m21 := float64(zy)
	m22 := float64(zz)

	t := m00 + m11 + m22

	var x, y, z, w float64
	if t >= 0 {
		s := math.Sqrt(t + 1)
		w = 0.5 * s
		s = 0.5 / s
		x = (m21 - m12) * s
		y = (m02 - m20) * s
		z = (m10 - m01) * s
	} else if m00 > m11 && m00 > m22 {
		s := math.Sqrt(1.0 + m00 + m11 - m22)
		x = s * 0.5
		s = 0.5 / s
		y = (m10 + m01) * s
		z = (m02 + m20) * s
		w = (m21 - m12) * s
	} else if m11 > m22 {
		s := math.Sqrt(1.0 + m22 - m00 - m11)
		z = s * 0.5
		s = 0.5 / s
		x = (m02 + m20) * s
		y = (m21 + m12) * s
		w = (m10 - m01) * s
	}

	return Quaternion{float32(x), float32(y), float32(z), float32(w)}
}

// Set this quaternion to the rotation between two vectors.
func (q Quaternion) FromCross(v1, v2 Vector3) Quaternion {
	dot := Clampf(v1.Dot(v2), -1.0, 1.0)
	angle := ToDegrees(Acos(dot))
	return q.FromAxis(v1.Y*v2.Z-v1.Z*v2.Y, v1.Z*v2.X-v1.X*v2.Z, v1.X*v2.Y-v1.Y*v2.X, angle)
}

// Spherical linear interpolation between this quaternion and the other quaternion, based on the alpha value in the range [0,1].
func (q Quaternion) Slerp(end Quaternion, alpha float32) Quaternion {
	if q.Eq(end) {
		return q
	}

	result := q.Dot(end)

	if result < 0 {
		end.Scale(-1)
		result = -result
	}

	scale0 := 1 - alpha
	scale1 := alpha

	if (1 - result) > 0.1 {
		theta := Acos(result)
		invSinTheta := 1 / Sin(theta)

		scale0 = Sin((1-alpha)*theta) * invSinTheta
		scale1 = Sin(alpha*theta) * invSinTheta
	}

	q.X = (scale0 * q.X) + (scale1 * end.X)
	q.Y = (scale0 * q.Y) + (scale1 * end.Y)
	q.Z = (scale0 * q.Z) + (scale1 * end.Z)
	q.W = (scale0 * q.W) + (scale1 * end.W)
	return q
}

// Fills a 4x4 matrix with the rotation matrix represented by this quaternion.
func (q Quaternion) Matrix() *Matrix4 {
	xx := q.X * q.X
	xy := q.X * q.Y
	xz := q.X * q.Z
	xw := q.X * q.W
	yy := q.Y * q.W
	yz := q.Y * q.Z
	yw := q.Y * q.W
	zz := q.Z * q.Z
	zw := q.Z * q.W
	// Set matrix from quaternion
	matrix := NewIdentityMatrix4()
	matrix.M11 = 1 - 2*(yy+zz)
	matrix.M21 = 2 * (xy - zw)
	matrix.M31 = 2 * (xz + yw)
	matrix.M41 = 0
	matrix.M12 = 2 * (xy + zw)
	matrix.M22 = 1 - 2*(xx+zz)
	matrix.M32 = 2 * (yz - xw)
	matrix.M42 = 0
	matrix.M13 = 2 * (xz - yw)
	matrix.M23 = 2 * (yz + xw)
	matrix.M33 = 1 - 2*(xx+yy)
	matrix.M43 = 0
	matrix.M14 = 0
	matrix.M24 = 0
	matrix.M34 = 0
	matrix.M44 = 1
	return matrix
}

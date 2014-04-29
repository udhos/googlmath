package math

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func Ry(origin, direction Vector3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r Ray) GetEndPoint(distance float32) Vector3 {
	return r.Origin.Add(r.Direction.Scale(distance))
}

// Multiplies the ray by the given matrix. Use this to transform a ray into another coordinate system.
func (r Ray) Mul(matrix Matrix4) Ray {
	tmp := r.Origin.Add(r.Direction)
	tmp = matrix.MulVec3(tmp)
	r.Origin = matrix.MulVec3(r.Origin)
	tmp = tmp.Sub(r.Origin)
	r.Direction = Vec3(tmp.X, tmp.Y, tmp.Z)
	return r
}

package math

import (
	"math"
)

var ZBB BoundingBox

type BoundingBox struct {
	Min Vector3
	Max Vector3
}

func BBox(Minimum, Maximum Vector3) BoundingBox {
	return BoundingBox{Minimum, Maximum} // TODO Should this perform Minimum<->Maximum correction? See func Set
}

func (box BoundingBox) Dx() float32 {
	return box.Max.X - box.Min.X
}

func (box BoundingBox) Dy() float32 {
	return box.Max.Y - box.Min.Y
}

func (box BoundingBox) Dz() float32 {
	return box.Max.Z - box.Min.Z
}

func (box BoundingBox) Center() Vector3 {
	dimension := box.Max.Sub(box.Min)
	dimension = dimension.Scale(0.5)
	return box.Min.Add(dimension)
}

// Order returns a new BoundingBox with ordered values.
// Min has the lower values and Maximum the higher values.
func (box BoundingBox) Order() BoundingBox {
	min := Vec3(Min(box.Min.X, box.Max.X), Min(box.Min.Y, box.Max.Y), Min(box.Min.Z, box.Max.Z))
	max := Vec3(Max(box.Min.X, box.Max.X), Max(box.Min.Y, box.Max.Y), Max(box.Min.Z, box.Max.Z))
	return BoundingBox{min, max}
}

func (box BoundingBox) IsValid() bool {
	return box.Min.X < box.Max.X && box.Min.Y < box.Max.Y && box.Min.Z < box.Max.Z
}

func (box BoundingBox) Corners() []Vector3 {
	corners := make([]Vector3, 8)
	corners[0] = Vec3(box.Min.X, box.Min.Y, box.Min.Z)
	corners[1] = Vec3(box.Max.X, box.Min.Y, box.Min.Z)
	corners[2] = Vec3(box.Max.X, box.Max.Y, box.Min.Z)
	corners[3] = Vec3(box.Min.X, box.Max.Y, box.Min.Z)
	corners[4] = Vec3(box.Min.X, box.Min.Y, box.Max.Z)
	corners[5] = Vec3(box.Max.X, box.Min.Y, box.Max.Z)
	corners[6] = Vec3(box.Max.X, box.Max.Y, box.Max.Z)
	corners[7] = Vec3(box.Min.X, box.Max.Y, box.Max.Z)
	return corners
}

func (box BoundingBox) Dimension() Vector3 {
	return box.Max.Sub(box.Min) // TODO Should this always return positive dimension values?
}

func (box BoundingBox) Extend(bounds BoundingBox) BoundingBox {
	box.Min.X = Min(box.Min.X, bounds.Min.X)
	box.Min.Y = Min(box.Min.Y, bounds.Min.Y)
	box.Min.Z = Min(box.Min.Z, bounds.Min.Z)
	box.Max.X = Max(box.Max.X, bounds.Max.X)
	box.Max.Y = Max(box.Max.Y, bounds.Max.Y)
	box.Max.Z = Max(box.Max.Z, bounds.Max.Z)
	return box
}

func (box BoundingBox) ExtendByVec(v Vector3) BoundingBox {
	box.Min.X = Min(box.Min.X, v.X)
	box.Min.Y = Min(box.Min.Y, v.Y)
	box.Min.Z = Min(box.Min.Z, v.Z)
	box.Max.X = Max(box.Max.X, v.X)
	box.Max.Y = Max(box.Max.Y, v.Y)
	box.Max.Z = Max(box.Max.Z, v.Z)
	return box
}

func (box BoundingBox) Contains(bounds BoundingBox) bool {
	return !box.IsValid() || (box.Min.X <= bounds.Min.X && box.Min.Y <= bounds.Min.Y && box.Min.Z <= bounds.Min.Z && box.Max.X >= bounds.Max.X && box.Max.Y >= bounds.Max.Y && box.Max.Z >= bounds.Max.Z)
}

func (box BoundingBox) Overlaps(bounds BoundingBox) bool {
	if bounds.ContainsVec(box.Min) || bounds.ContainsVec(box.Max) {
		return true
	}
	if bounds.Max.X < box.Max.X && bounds.Max.Y < box.Max.Y && bounds.Max.Z < box.Max.Z {
		if bounds.Min.X > box.Min.X && bounds.Min.Y > box.Min.Y && bounds.Min.Z > box.Min.Z {
			return true
		}
	}
	return false
}

func (box BoundingBox) ContainsVec(v Vector3) bool {
	return box.Min.X <= v.X && box.Max.X >= v.X && box.Min.Y <= v.Y && box.Max.Y >= v.Y && box.Min.Z <= v.Z && box.Max.Z >= v.Z
}

func (box BoundingBox) Inf() BoundingBox {
	box.Min = Vec3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)
	box.Max = Vec3(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32)
	return box
}

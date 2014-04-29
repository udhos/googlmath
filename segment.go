package math

// A Segment is a line in 3D-space having a staring and an ending position.
type Segment struct {
	A, B Vector3
}

func Sgt(a, b Vector3) Segment {
	return Segment{a, b}
}

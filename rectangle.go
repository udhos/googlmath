package math

var ZR Rectangle

type Rectangle struct {
	Min, Max Vector2
}

func Rect(x0, y0, x1, y1 float32) Rectangle {
	return Rectangle{Vector2{x0, y0}, Vector2{x1, y1}}
}

// Dx returns r's width.
func (r Rectangle) Dx() float32 {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rectangle) Dy() float32 {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle) In(s Rectangle) bool {
	if r.Empty() {
		return true
	}
	// Note that r.Max is an exclusive bound for r, so that r.In(s)
	// does not require that r.Max.In(s).
	return s.Min.X <= r.Min.X && r.Max.X <= s.Max.X && s.Min.Y <= r.Min.Y && r.Max.Y <= s.Max.Y
}

// Overlaps reports whether r and s have a non-empty intersection.
func (r Rectangle) Overlaps(s Rectangle) bool {
	return r.Min.X < s.Max.X && s.Min.X < r.Max.X && r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}

// Contains reports whether r contains v.
// Note that r is an inclusive bound.
func (r Rectangle) Contains(v Vector2) bool {
	return r.Min.X <= v.X && r.Max.X >= v.X && r.Min.Y <= v.Y && r.Max.Y >= v.Y
}

// Empty reports whether the rectangle contains no points.
func (r Rectangle) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Merge merges this rectangle with the other rectangle.
func (r Rectangle) Merge(s Rectangle) Rectangle {
	return Rectangle{
		Min: Vector2{Min(r.Min.X, s.Min.X), Min(r.Min.Y, s.Min.Y)},
		Max: Vector2{Max(r.Max.X, s.Max.X), Max(r.Max.Y, s.Max.Y)},
	}
}

func (r Rectangle) Center() Vector2 {
	return r.Max.Sub(r.Max.Sub(r.Min).Scale(0.5))
}

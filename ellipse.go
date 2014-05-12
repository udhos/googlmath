package math

// A two dimension ellipse.
type Ellipse struct {
	Center Vector2
	Width  float32
	Height float32
}

func Ell(x, y, width, height float32) Ellipse {
	return Ellipse{Vec2(x, y), width, height}
}

func (e Ellipse) Contains(v Vector2) bool {
	if e.Width <= 0.0 {
		return false
	}
	if e.Height <= 0.0 {
		return false
	}
	v.Sub(e.Center)
	r := Vec2(e.Width/2, e.Height/2)

	return v.X*v.X/r.X*r.X+v.Y*v.Y/r.Y*r.Y <= 1
}

package math

// A two dimension ellipse.
type Ellipse struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func Ell(x, y, width, height float32) Ellipse {
	return Ellipse{x, y, width, height}
}

func (e Ellipse) Contains(x, y float32) bool {
	if e.Width <= 0.0 {
		return false
	}
	if e.Height <= 0.0 {
		return false
	}
	x = x - e.X
	y = y - e.Y

	xr := (e.Width / 2)
	yr := (e.Height / 2)

	return x*x/xr*xr+y*y/yr*yr <= 1
}

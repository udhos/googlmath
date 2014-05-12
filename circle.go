package math

// A two dimension circle.
type Circle struct {
	Center Vector2
	Radius float32
}

func Circ(x, y, radius float32) Circle {
	return Circle{Vec2(x, y), radius}
}

func (c Circle) Contains(v Vector2) bool {
	s := c.Center.Sub(v)
	return s.Len2() <= c.Radius*c.Radius
}

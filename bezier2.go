package math

// Implementation of the Bezier curve in two dimensions.
type Bezier2 struct {
	Points []Vector2
}

func NewBezier2(points ...Vector2) Path2 {
	return &Bezier2{points}
}

func (b *Bezier2) Set(points ...Vector2) Path2 {
	b.Points = points
	return b
}

// The value of the path at t where 0<=t<=1
func (b *Bezier2) ValueAt(t float32) Vector2 {
	out := Vec2(0, 0)
	n := len(b.Points)

	if n == 2 {
		out = Linear2(t, b.Points[0], b.Points[1])
	} else if n == 3 {
		out = Quadratic2(t, b.Points[0], b.Points[1], b.Points[2])
	} else if n == 4 {
		out = Cubic2(t, b.Points[0], b.Points[1], b.Points[2], b.Points[3])
	}
	return out
}

// The approximated value (between 0 and 1) on the path which is closest to the specified value.
func (b *Bezier2) Approximate(p3 Vector2) float32 {
	p1 := b.Points[0]
	p2 := b.Points[len(b.Points)-1]

	l1 := p1.Dst(p2)
	l2 := p3.Dst(p2)
	l3 := p3.Dst(p1)
	s := (l2*l2 + l1*l1 - l3*l3) / (2 * l1)

	return Clampf((l1-s)/l1, 0.0, 1.0)
}

// Simple linear interpolation
func Linear2(t float32, p0, p1 Vector2) Vector2 {
	tmp := p1.Cpy()
	out := p0.Cpy()
	return out.Scale(1.0 - t).Add(tmp.Scale(t))
}

// Quadratic Bezier curve
func Quadratic2(t float32, p0, p1, p2 Vector2) Vector2 {
	dt := 1.0 - t
	tmp := p1.Cpy()
	tmp2 := p2.Cpy()
	out := p0.Cpy()
	return out.Scale(dt * dt).Add(tmp.Scale(2 * dt * t)).Add(tmp2.Scale(t * t))
}

// Cubic Bezier curve
func Cubic2(t float32, p0, p1, p2, p3 Vector2) Vector2 {
	dt := 1 - t
	dt2 := dt * dt
	t2 := t * t
	out := p0.Cpy()
	tmp := p1.Cpy()
	tmp2 := p2.Cpy()
	tmp3 := p3.Cpy()
	return out.Scale(dt2 * dt).Add(tmp.Scale(3 * dt2 * t)).Add(tmp2.Scale(3 * dt * t2)).Add(tmp3.Scale(t2 * t))
}

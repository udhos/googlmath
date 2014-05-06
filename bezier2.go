package math

// Implementation of the Bezier curve in two dimensions.
type Bezier2 []Vector2

func Bzr2(points ...Vector2) Path2 {
	return Bezier2(points)
}

// The value of the path at t where 0<=t<=1
func (b Bezier2) ValueAt(t float32) Vector2 {
	if len(b) == 2 {
		return Linear2(t, b[0], b[1])
	} else if len(b) == 3 {
		return Quadratic2(t, b[0], b[1], b[2])
	} else if len(b) == 4 {
		return Cubic2(t, b[0], b[1], b[2], b[3])
	}
	return Vector2{NaN(), NaN()}
}

// The approximated value (between 0 and 1) on the path which is closest to the specified value.
func (b Bezier2) Approximate(p3 Vector2) float32 {
	if len(b) < 2 {
		return NaN()
	}
	p1 := b[0]
	p2 := b[len(b)-1]

	l1 := p1.Distance(p2)
	l2 := p3.Distance(p2)
	l3 := p3.Distance(p1)
	s := (l2*l2 + l1*l1 - l3*l3) / (2 * l1)

	return Clampf((l1-s)/l1, 0.0, 1.0)
}

// Simple linear interpolation
func Linear2(t float32, p0, p1 Vector2) Vector2 {
	return p0.Scale(1.0 - t).Add(p1.Scale(t))
}

// Quadratic Bezier curve
func Quadratic2(t float32, p0, p1, p2 Vector2) Vector2 {
	dt := 1.0 - t
	return p0.Scale(dt * dt).Add(p1.Scale(2 * dt * t)).Add(p2.Scale(t * t))
}

// Cubic Bezier curve
func Cubic2(t float32, p0, p1, p2, p3 Vector2) Vector2 {
	dt := 1 - t
	dt2 := dt * dt
	t2 := t * t
	return p0.Scale(dt2 * dt).Add(p1.Scale(3 * dt2 * t)).Add(p2.Scale(3 * dt * t2)).Add(p3.Scale(t2 * t))
}

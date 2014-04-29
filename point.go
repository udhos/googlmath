package math

import "strconv"

type Point struct {
	X, Y int
}

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Mul(q Point) Point {
	return Point{p.X * q.X, p.Y * q.Y}
}

func (p Point) Div(q Point) Point {
	return Point{p.X / q.X, p.Y / q.Y}
}

func (p Point) Eq(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

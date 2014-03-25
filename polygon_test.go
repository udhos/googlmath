package math

import (
	. "launchpad.net/gocheck"
)

type PolygonTestSuite struct{}

var _ = Suite(&PolygonTestSuite{})

func (s *PolygonTestSuite) TestPolygonTransformedVertices(c *C) {
	poly, _ := NewPolygon([]float32{
		-1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
	})
	expected := []float32{
		-1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
	}
	actual := poly.Vertices()

	c.Assert(actual, DeepEquals, expected)

	poly.Translate(Vector2{1, 1})
	expected = []float32{
		0.0, 1.0,
		1.0, 2.0,
		2.0, 1.0,
	}
	actual = poly.TransformedVertices()
	c.Assert(actual, DeepEquals, expected)
}

package math

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestPolygonTransformedVertices(c *C) {
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

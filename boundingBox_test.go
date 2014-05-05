package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_2V3_BB struct {
	Min, Max Vector3
	Expected BoundingBox
}

type TestValue_BB_B struct {
	Value    BoundingBox
	Expected bool
}

type TestValue_2BB_B struct {
	Box      BoundingBox
	Bounds   BoundingBox
	Expected bool
}

type TestValue_BB_V3s struct {
	Value    BoundingBox
	Expected []Vector3
}

type TestValue_BB_V3 struct {
	Value    BoundingBox
	Expected Vector3
}

type TestValue_BBV3_B struct {
	Value    BoundingBox
	Vec      Vector3
	Expected bool
}

func (s *S) TestBBox(c *C) {
	tests := []TestValue_2V3_BB{
		TestValue_2V3_BB{
			Min:      Vec3(0, 0, 0),
			Max:      Vec3(1, 2, 3),
			Expected: BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		TestValue_2V3_BB{
			Min:      Vec3(-1, -2.2, 0),
			Max:      Vec3(2, 3, 3),
			Expected: BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)},
		},
	}
	for _, value := range tests {
		obtained := BBox(value.Min, value.Max)
		c.Check(obtained, BoundingBoxCheck, value.Expected)
	}
}

type TestValue_BB_F32 struct {
	b BoundingBox
	e float32
}

func (s *S) TestBBoxDx(c *C) {
	tests := []TestValue_BB_F32{
		{BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)}, 1.0},
		{BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)}, 3.0},
	}
	for _, t := range tests {
		obtained := t.b.Dx()
		c.Check(obtained, Equals, t.e)
	}
}

func (s *S) TestBBoxDy(c *C) {
	tests := []TestValue_BB_F32{
		{BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)}, 2.0},
		{BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)}, 5.2},
	}
	for _, t := range tests {
		obtained := t.b.Dy()
		c.Check(obtained, Equals, t.e)
	}
}

func (s *S) TestBBoxDz(c *C) {
	tests := []TestValue_BB_F32{
		{BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)}, 3.0},
		{BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)}, 3.0},
	}
	for _, t := range tests {
		obtained := t.b.Dz()
		c.Check(obtained, Equals, t.e)
	}
}

func (s *S) TestBBoxCenter(c *C) {
	tests := []TestValue_BB_V3{
		{BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)}, Vec3(0.5, 1, 1.5)},
		{BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)}, Vec3(0.5, 0.4, 1.5)},
	}
	for _, t := range tests {
		obtained := t.Value.Center()
		c.Check(obtained, Vector3Check, t.Expected)
	}
}

func (s *S) TestBBoxOrder(c *C) {
	tests := []TestValue_2V3_BB{
		TestValue_2V3_BB{
			Min:      Vec3(0, 0, 0),
			Max:      Vec3(1, 2, 3),
			Expected: BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		TestValue_2V3_BB{
			Min:      Vec3(-1, -2.2, 0),
			Max:      Vec3(2, 3, 3),
			Expected: BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)},
		},
	}
	for _, value := range tests {
		bb := BBox(value.Min, value.Max)
		obtained := bb.Order()
		c.Check(bb, BoundingBoxCheck, obtained)
		c.Check(obtained, BoundingBoxCheck, value.Expected)
	}
}

func (s *S) TestBBoxIsValid(c *C) {
	tests := []TestValue_BB_B{
		TestValue_BB_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			true,
		},
		TestValue_BB_B{
			BBox(Vec3(0, 0, 0), Vec3(1.2, 2.3, 4.3)),
			true,
		},
		TestValue_BB_B{
			BBox(Vec3(-2, -1, 1), Vec3(1, 0, 3)),
			true,
		},
		TestValue_BB_B{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			false,
		},
		TestValue_BB_B{
			BBox(Vec3(2, 2, 2), Vec3(1, 1, 1)),
			false,
		},
		TestValue_BB_B{
			BBox(Vec3(0, 2, 1), Vec3(0, 2, 1)),
			false,
		},
	}
	for _, value := range tests {
		obtained := value.Value.IsValid()
		c.Check(obtained, Equals, value.Expected, Commentf("%v.IsValid()==%t epected %t", value.Value, obtained, value.Expected))
	}
}

func (s *S) TestBBoxCorners(c *C) {
	tests := []TestValue_BB_V3s{
		TestValue_BB_V3s{
			BBox(Vec3(-1, -2, -3), Vec3(1, 2, 3)),
			[]Vector3{
				Vec3(-1, -2, -3),
				Vec3(1, -2, -3),
				Vec3(1, 2, -3),
				Vec3(-1, 2, -3),
				Vec3(-1, -2, 3),
				Vec3(1, -2, 3),
				Vec3(1, 2, 3),
				Vec3(-1, 2, 3),
			},
		},
	}
	for _, value := range tests {
		obtained := value.Value.Corners()
		c.Check(obtained, DeepEquals, value.Expected)
	}
}

func (s *S) TestBBoxDimension(c *C) {
	tests := []TestValue_BB_V3{
		TestValue_BB_V3{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			Vec3(1, 1, 1),
		},
		TestValue_BB_V3{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			Vec3(0, 0, 0),
		},
		TestValue_BB_V3{
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
			Vec3(2, 2, 1),
		},
		TestValue_BB_V3{
			BBox(Vec3(-1, -1, -1), Vec3(-2, -2, -2)),
			Vec3(-1, -1, -1),
		},
	}
	for _, value := range tests {
		obtained := value.Value.Dimension()
		c.Check(obtained, DeepEquals, value.Expected)
	}
}

func (s *S) TestBBoxContainsVec(c *C) {
	tests := []TestValue_BBV3_B{
		TestValue_BBV3_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			Vec3(0.5, 0.5, 0),
			true,
		},
		TestValue_BBV3_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			Vec3(0.5, 0.5, 0),
			true,
		},
		TestValue_BBV3_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			Vec3(0.5, 0.5, -1),
			false,
		},
		TestValue_BBV3_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			Vec3(0, 0, 0),
			true,
		},
	}
	for _, value := range tests {
		obtained := value.Value.ContainsVec(value.Vec)
		c.Check(obtained, Equals, value.Expected)
	}
}

func (s *S) TestBBoxOverlaps(c *C) {
	tests := []TestValue_2BB_B{
		TestValue_2BB_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			BBox(Vec3(-1, -1, 0), Vec3(0, 0, 0)),
			true,
		},
		TestValue_2BB_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, -1, -1), Vec3(-0.1, -0.1, -0.1)),
			false,
		},
		TestValue_2BB_B{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, 0, 0), Vec3(0, 0, 0.5)),
			true,
		},
	}
	for _, value := range tests {
		obtained := value.Box.Overlaps(value.Bounds)
		c.Check(obtained, Equals, value.Expected, Commentf("Ovleraps(%v,%v) %t Expected:%t", value.Box, value.Bounds, obtained, value.Expected))
	}
}

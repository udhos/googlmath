package math

import (
	. "gopkg.in/check.v1"
)

type TestValue_2V3_BB struct {
	v0 Vector3
	v1 Vector3
	e  BoundingBox
}

type TestValue_BB_B struct {
	b BoundingBox
	e bool
}

type TestValue_2BB_B struct {
	b0 BoundingBox
	b1 BoundingBox
	e  bool
}

type TestValue_BB_V3s struct {
	b BoundingBox
	e []Vector3
}

type TestValue_BB_V3 struct {
	b BoundingBox
	e Vector3
}

type TestValue_BBV3_B struct {
	b BoundingBox
	v Vector3
	e bool
}

func (s *S) TestBBox(c *C) {
	tests := []TestValue_2V3_BB{
		{
			Vec3(0, 0, 0),
			Vec3(1, 2, 3),
			BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		{
			Vec3(-1, -2.2, 0),
			Vec3(2, 3, 3),
			BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)},
		},
	}
	for _, t := range tests {
		obtained := BBox(t.v0, t.v1)
		c.Check(obtained, BoundingBoxCheck, t.e)
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
		obtained := t.b.Center()
		c.Check(obtained, Vector3Check, t.e)
	}
}

func (s *S) TestBBoxOrder(c *C) {
	tests := []TestValue_2V3_BB{
		{
			Vec3(0, 0, 0),
			Vec3(1, 2, 3),
			BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		{
			Vec3(-1, -2.2, 0),
			Vec3(2, 3, 3),
			BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)},
		},
	}
	for _, t := range tests {
		bb := BBox(t.v0, t.v1)
		obtained := bb.Order()
		c.Check(bb, BoundingBoxCheck, obtained)
		c.Check(obtained, BoundingBoxCheck, t.e)
	}
}

func (s *S) TestBBoxIsValid(c *C) {
	tests := []TestValue_BB_B{
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			true,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(1.2, 2.3, 4.3)),
			true,
		},
		{
			BBox(Vec3(-2, -1, 1), Vec3(1, 0, 3)),
			true,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			false,
		},
		{
			BBox(Vec3(2, 2, 2), Vec3(1, 1, 1)),
			false,
		},
		{
			BBox(Vec3(0, 2, 1), Vec3(0, 2, 1)),
			false,
		},
	}
	for _, t := range tests {
		obtained := t.b.IsValid()
		c.Check(obtained, Equals, t.e, Commentf("%v.IsValid()==%t epected %t", t.b, obtained, t.e))
	}
}

func (s *S) TestBBoxCorners(c *C) {
	tests := []TestValue_BB_V3s{
		{
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
	for _, t := range tests {
		obtained := t.b.Corners()
		c.Check(obtained, DeepEquals, t.e)
	}
}

func (s *S) TestBBoxDimension(c *C) {
	tests := []TestValue_BB_V3{
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			Vec3(1, 1, 1),
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			Vec3(0, 0, 0),
		},
		{
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
			Vec3(2, 2, 1),
		},
		{
			BBox(Vec3(-1, -1, -1), Vec3(-2, -2, -2)),
			Vec3(-1, -1, -1),
		},
	}
	for _, t := range tests {
		obtained := t.b.Dimension()
		c.Check(obtained, DeepEquals, t.e)
	}
}

type TestValue_2BB_BB struct {
	b0 BoundingBox
	b1 BoundingBox
	e  BoundingBox
}

func (s *S) TestBBoxExtend(c *C) {
	tests := []TestValue_2BB_BB{
		{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
		},
		{
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
		},
		{
			BBox(Vec3(-1, -1, -1), Vec3(-2, -2, -2)),
			BBox(Vec3(-1, -2.2, -0.51), Vec3(-42, 33.3, 12.25)),
			BBox(Vec3(-1, -2.2, -1), Vec3(-2, 33.3, 12.25)),
		},
	}
	for _, t := range tests {
		obtained := t.b0.Extend(t.b1)
		c.Check(obtained, DeepEquals, t.e)
	}
}

type TestValue_BBV3_BB struct {
	b BoundingBox
	v Vector3
	e BoundingBox
}

func (s *S) TestBBoxExtendByVec(c *C) {
	tests := []TestValue_BBV3_BB{
		{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			Vec3(0, 0, 0),
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
		},
		{
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
			Vec3(-1, -1, -1),
			BBox(Vec3(-1, -1, -1), Vec3(1, 1, 1)),
		},
		{
			BBox(Vec3(-1, -1, -1), Vec3(-2, -2, -2)),
			Vec3(-1, -2.2, 12.25),
			BBox(Vec3(-1, -2.2, -1), Vec3(-1, -2, 12.25)),
		},
	}
	for _, t := range tests {
		obtained := t.b.ExtendByVec(t.v)
		c.Check(obtained, DeepEquals, t.e)
	}
}

func (s *S) TestBBoxContains(c *C) {
	tests := []TestValue_2BB_B{
		{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			true,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(0.1, 0, 0.25), Vec3(0.5, 0.25, 1)),
			true,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1.00001)),
			false,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-0.1, 0, 0), Vec3(1, 1, 0)),
			false,
		},
	}
	for _, t := range tests {
		obtained := t.b0.Contains(t.b1)
		c.Check(obtained, Equals, t.e, Commentf("%v.Contains(%v)==%t", t.b0, t.b1, obtained))
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
	for _, t := range tests {
		obtained := t.b.ContainsVec(t.v)
		c.Check(obtained, Equals, t.e)
	}
}

func (s *S) TestBBoxOverlaps(c *C) {
	tests := []TestValue_2BB_B{
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			BBox(Vec3(-1, -1, 0), Vec3(0, 0, 0)),
			true,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, -1, -1), Vec3(-0.1, -0.1, -0.1)),
			false,
		},
		{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, 0, 0), Vec3(0, 0, 0.5)),
			true,
		},
	}
	for _, t := range tests {
		obtained := t.b0.Overlaps(t.b1)
		c.Check(obtained, Equals, t.e, Commentf("Ovleraps(%v,%v) %t e:%t", t.b0, t.b1, obtained, t.e))
	}
}

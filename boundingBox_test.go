package math

import (
	. "gopkg.in/check.v1"
)

type BB2Vec3TestValue struct {
	Min, Max Vector3
	Expected BoundingBox
}

type BBBoolTestValue struct {
	Value    BoundingBox
	Expected bool
}

type BB2BoolTestValue struct {
	Box      BoundingBox
	Bounds   BoundingBox
	Expected bool
}

type BBVec3ArrayTestValue struct {
	Value    BoundingBox
	Expected []Vector3
}

type BBVec3TestValue struct {
	Value    BoundingBox
	Expected Vector3
}

type BBVec3BoolTestValue struct {
	Value    BoundingBox
	Vec      Vector3
	Expected bool
}

func (s *S) TestBBox(c *C) {
	tests := []BB2Vec3TestValue{
		BB2Vec3TestValue{
			Min:      Vec3(0, 0, 0),
			Max:      Vec3(1, 2, 3),
			Expected: BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		BB2Vec3TestValue{
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

func (s *S) TestBBoxOrder(c *C) {
	tests := []BB2Vec3TestValue{
		BB2Vec3TestValue{
			Min:      Vec3(0, 0, 0),
			Max:      Vec3(1, 2, 3),
			Expected: BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		BB2Vec3TestValue{
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
	tests := []BBBoolTestValue{
		BBBoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			true,
		},
		BBBoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1.2, 2.3, 4.3)),
			true,
		},
		BBBoolTestValue{
			BBox(Vec3(-2, -1, 1), Vec3(1, 0, 3)),
			true,
		},
		BBBoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			false,
		},
		BBBoolTestValue{
			BBox(Vec3(2, 2, 2), Vec3(1, 1, 1)),
			false,
		},
		BBBoolTestValue{
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
	tests := []BBVec3ArrayTestValue{
		BBVec3ArrayTestValue{
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
	tests := []BBVec3TestValue{
		BBVec3TestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			Vec3(1, 1, 1),
		},
		BBVec3TestValue{
			BBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			Vec3(0, 0, 0),
		},
		BBVec3TestValue{
			BBox(Vec3(-1, -1, 0), Vec3(1, 1, 1)),
			Vec3(2, 2, 1),
		},
		BBVec3TestValue{
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
	tests := []BBVec3BoolTestValue{
		BBVec3BoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			Vec3(0.5, 0.5, 0),
			true,
		},
		BBVec3BoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			Vec3(0.5, 0.5, 0),
			true,
		},
		BBVec3BoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			Vec3(0.5, 0.5, -1),
			false,
		},
		BBVec3BoolTestValue{
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
	tests := []BB2BoolTestValue{
		BB2BoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 0)),
			BBox(Vec3(-1, -1, 0), Vec3(0, 0, 0)),
			true,
		},
		BB2BoolTestValue{
			BBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			BBox(Vec3(-1, -1, -1), Vec3(-0.1, -0.1, -0.1)),
			false,
		},
		BB2BoolTestValue{
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

package math

import (
	. "gopkg.in/check.v1"
)

type EllipseContainTestValue struct {
	X, Y     float32
	Expected bool
}

type EllipseTestSuite struct {
	e                Ellipse
	containTestTable []EllipseContainTestValue
}

var _ = Suite(&EllipseTestSuite{})

func (s *EllipseTestSuite) SetUpTest(c *C) {
	s.e = Ell(0, 0, 1, 1)
	s.containTestTable = []EllipseContainTestValue{
		EllipseContainTestValue{0, 0, true},
		EllipseContainTestValue{1, 0, true},
		EllipseContainTestValue{-2, -2, false},
		EllipseContainTestValue{2, 0, false},
	}
}

func (s *EllipseTestSuite) TestEllipse(c *C) {
	// Contains
	for i := range s.containTestTable {
		c.Assert(s.e.Contains(s.containTestTable[i].X, s.containTestTable[i].Y), Equals, s.containTestTable[i].Expected)
	}
}

package math

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
}

var _ = Suite(&S{})

// TODO Use one suite to run all tests in this package
// TODO Test values are created/stored in each test function
// TODO A test function for each public function
// TODO Name func (s *S) Test$FUNCNAME$PARAMETER

package math

import (
	. "gopkg.in/check.v1"
)

var Vector2Check = &Vector2Checker{}

type Vector2Checker struct{}

func (checker *Vector2Checker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Vector2Checker", Params: []string{"obtained", "expected"}}
}

func (checker *Vector2Checker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var v1, v2 Vector2
	var ok bool

	v1, ok = (params[0]).(Vector2)
	if ok == false {
		return false, "Param[0] not a Vector2 type"
	}
	v2, ok = (params[1]).(Vector2)
	if ok == false {
		return false, "Param[1] not a Vector2 type"
	}

	return NearlyEqualFloat32(v1.X, v2.X) && NearlyEqualFloat32(v1.Y, v2.Y), ""
}

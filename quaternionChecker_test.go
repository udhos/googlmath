package math

import (
	. "gopkg.in/check.v1"
)

var QuaternionCheck = &QuaternionChecker{}

type QuaternionChecker struct{}

func (checker *QuaternionChecker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "QuaternionChecker", Params: []string{"obtained", "expected"}}
}

func (checker *QuaternionChecker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var q1, q2 Quaternion
	var ok bool

	q1, ok = (params[0]).(Quaternion)
	if ok == false {
		return false, "Param[0] not a Quaternion type"
	}
	q2, ok = (params[1]).(Quaternion)
	if ok == false {
		return false, "Param[1] not a Quaternion type"
	}

	return NearlyEqualFloat32(q1.X, q2.X) && NearlyEqualFloat32(q1.Y, q2.Y) && NearlyEqualFloat32(q1.Z, q2.Z) && NearlyEqualFloat32(q1.W, q2.W), ""
}

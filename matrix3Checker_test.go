package math

import (
	. "gopkg.in/check.v1"
)

var Matrix3Check = &Matrix3Checker{}

type Matrix3Checker struct{}

func (checker Matrix3Checker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Matrix3Checker", Params: []string{"obtained", "expected"}}
}

func (checker Matrix3Checker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var m1, m2 Matrix3
	var ok bool

	m1, ok = (params[0]).(Matrix3)
	if ok == false {
		return false, "Param[0] not a Matrix3 type"
	}
	m2, ok = (params[1]).(Matrix3)
	if ok == false {
		return false, "Param[1] not a Matrix3 type"
	}

	return NearlyEqualFloat32(m1.M11, m2.M11) && NearlyEqualFloat32(m1.M12, m2.M12) && NearlyEqualFloat32(m1.M13, m2.M13) &&
		NearlyEqualFloat32(m1.M21, m2.M21) && NearlyEqualFloat32(m1.M22, m2.M22) && NearlyEqualFloat32(m1.M23, m2.M23) &&
		NearlyEqualFloat32(m1.M31, m2.M31) && NearlyEqualFloat32(m1.M32, m2.M32) && NearlyEqualFloat32(m1.M33, m2.M33), ""
}

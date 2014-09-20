package math

import (
	"testing"
)

const (
	Matrix4M00 = 0  // 0
	Matrix4M01 = 4  // 1
	Matrix4M02 = 8  // 2
	Matrix4M03 = 12 // 3
	Matrix4M10 = 1  // 4
	Matrix4M11 = 5  // 5
	Matrix4M12 = 9  // 6
	Matrix4M13 = 13 // 7
	Matrix4M20 = 2  // 8
	Matrix4M21 = 6  // 9
	Matrix4M22 = 10 // 10
	Matrix4M23 = 14 // 11
	Matrix4M30 = 3  // 12
	Matrix4M31 = 7  // 13
	Matrix4M32 = 11 // 14
	Matrix4M33 = 15 // 15
)

type Matrix4A [16]float32

func (m Matrix4A) Mul(n Matrix4A) Matrix4A {
	return Matrix4A{
		m[Matrix4M00]*n[Matrix4M00] + m[Matrix4M01]*n[Matrix4M10] + m[Matrix4M02]*n[Matrix4M20] + m[Matrix4M03]*n[Matrix4M30],
		m[Matrix4M00]*n[Matrix4M01] + m[Matrix4M01]*n[Matrix4M11] + m[Matrix4M02]*n[Matrix4M21] + m[Matrix4M03]*n[Matrix4M31],
		m[Matrix4M00]*n[Matrix4M02] + m[Matrix4M01]*n[Matrix4M12] + m[Matrix4M02]*n[Matrix4M22] + m[Matrix4M03]*n[Matrix4M32],
		m[Matrix4M00]*n[Matrix4M03] + m[Matrix4M01]*n[Matrix4M13] + m[Matrix4M02]*n[Matrix4M23] + m[Matrix4M03]*n[Matrix4M33],
		m[Matrix4M10]*n[Matrix4M00] + m[Matrix4M11]*n[Matrix4M10] + m[Matrix4M12]*n[Matrix4M20] + m[Matrix4M13]*n[Matrix4M30],
		m[Matrix4M10]*n[Matrix4M01] + m[Matrix4M11]*n[Matrix4M11] + m[Matrix4M12]*n[Matrix4M21] + m[Matrix4M13]*n[Matrix4M31],
		m[Matrix4M10]*n[Matrix4M02] + m[Matrix4M11]*n[Matrix4M12] + m[Matrix4M12]*n[Matrix4M22] + m[Matrix4M13]*n[Matrix4M32],
		m[Matrix4M10]*n[Matrix4M03] + m[Matrix4M11]*n[Matrix4M13] + m[Matrix4M12]*n[Matrix4M23] + m[Matrix4M13]*n[Matrix4M33],
		m[Matrix4M20]*n[Matrix4M00] + m[Matrix4M21]*n[Matrix4M10] + m[Matrix4M22]*n[Matrix4M20] + m[Matrix4M23]*n[Matrix4M30],
		m[Matrix4M20]*n[Matrix4M01] + m[Matrix4M21]*n[Matrix4M11] + m[Matrix4M22]*n[Matrix4M21] + m[Matrix4M23]*n[Matrix4M31],
		m[Matrix4M20]*n[Matrix4M02] + m[Matrix4M21]*n[Matrix4M12] + m[Matrix4M22]*n[Matrix4M22] + m[Matrix4M23]*n[Matrix4M32],
		m[Matrix4M20]*n[Matrix4M03] + m[Matrix4M21]*n[Matrix4M13] + m[Matrix4M22]*n[Matrix4M23] + m[Matrix4M23]*n[Matrix4M33],
		m[Matrix4M30]*n[Matrix4M00] + m[Matrix4M31]*n[Matrix4M10] + m[Matrix4M32]*n[Matrix4M20] + m[Matrix4M33]*n[Matrix4M30],
		m[Matrix4M30]*n[Matrix4M01] + m[Matrix4M31]*n[Matrix4M11] + m[Matrix4M32]*n[Matrix4M21] + m[Matrix4M33]*n[Matrix4M31],
		m[Matrix4M30]*n[Matrix4M02] + m[Matrix4M31]*n[Matrix4M12] + m[Matrix4M32]*n[Matrix4M22] + m[Matrix4M33]*n[Matrix4M32],
		m[Matrix4M30]*n[Matrix4M03] + m[Matrix4M31]*n[Matrix4M13] + m[Matrix4M32]*n[Matrix4M23] + m[Matrix4M33]*n[Matrix4M33]}
}

func BenchmarkMatrix4AMul(b *testing.B) {
	var m1, m2 Matrix4A
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m3 := m1.Mul(m2)
		_ = m3
	}
}

type Matrix4S struct {
	M11, M12, M13, M14 float32
	M21, M22, M23, M24 float32
	M31, M32, M33, M34 float32
	M41, M42, M43, M44 float32
}

// Multiplicates this matrix with m2 matrix and returns the new matrix.
func (m1 *Matrix4S) Mul(m2 *Matrix4S) *Matrix4S {
	return &Matrix4S{
		m1.M11*m2.M11 + m1.M21*m2.M12 + m1.M31*m2.M13 + m1.M41*m2.M14,
		m1.M12*m2.M11 + m1.M22*m2.M12 + m1.M32*m2.M13 + m1.M42*m2.M14,
		m1.M13*m2.M11 + m1.M23*m2.M12 + m1.M33*m2.M13 + m1.M43*m2.M14,
		m1.M14*m2.M11 + m1.M24*m2.M12 + m1.M34*m2.M13 + m1.M44*m2.M14,
		m1.M11*m2.M21 + m1.M21*m2.M22 + m1.M31*m2.M23 + m1.M41*m2.M24,
		m1.M12*m2.M21 + m1.M22*m2.M22 + m1.M32*m2.M23 + m1.M42*m2.M24,
		m1.M13*m2.M21 + m1.M23*m2.M22 + m1.M33*m2.M23 + m1.M43*m2.M24,
		m1.M14*m2.M21 + m1.M24*m2.M22 + m1.M34*m2.M23 + m1.M44*m2.M24,
		m1.M11*m2.M31 + m1.M21*m2.M32 + m1.M31*m2.M33 + m1.M41*m2.M34,
		m1.M12*m2.M31 + m1.M22*m2.M32 + m1.M32*m2.M33 + m1.M42*m2.M34,
		m1.M13*m2.M31 + m1.M23*m2.M32 + m1.M33*m2.M33 + m1.M43*m2.M34,
		m1.M14*m2.M31 + m1.M24*m2.M32 + m1.M34*m2.M33 + m1.M44*m2.M34,
		m1.M11*m2.M41 + m1.M21*m2.M42 + m1.M31*m2.M43 + m1.M41*m2.M44,
		m1.M12*m2.M41 + m1.M22*m2.M42 + m1.M32*m2.M43 + m1.M42*m2.M44,
		m1.M13*m2.M41 + m1.M23*m2.M42 + m1.M33*m2.M43 + m1.M43*m2.M44,
		m1.M14*m2.M41 + m1.M24*m2.M42 + m1.M34*m2.M43 + m1.M44*m2.M44}
}

// Multiplicates this matrix with m2 matrix and returns the new matrix.
func (m1 Matrix4S) MulV(m2 Matrix4S) Matrix4S {
	return Matrix4S{
		m1.M11*m2.M11 + m1.M21*m2.M12 + m1.M31*m2.M13 + m1.M41*m2.M14,
		m1.M12*m2.M11 + m1.M22*m2.M12 + m1.M32*m2.M13 + m1.M42*m2.M14,
		m1.M13*m2.M11 + m1.M23*m2.M12 + m1.M33*m2.M13 + m1.M43*m2.M14,
		m1.M14*m2.M11 + m1.M24*m2.M12 + m1.M34*m2.M13 + m1.M44*m2.M14,
		m1.M11*m2.M21 + m1.M21*m2.M22 + m1.M31*m2.M23 + m1.M41*m2.M24,
		m1.M12*m2.M21 + m1.M22*m2.M22 + m1.M32*m2.M23 + m1.M42*m2.M24,
		m1.M13*m2.M21 + m1.M23*m2.M22 + m1.M33*m2.M23 + m1.M43*m2.M24,
		m1.M14*m2.M21 + m1.M24*m2.M22 + m1.M34*m2.M23 + m1.M44*m2.M24,
		m1.M11*m2.M31 + m1.M21*m2.M32 + m1.M31*m2.M33 + m1.M41*m2.M34,
		m1.M12*m2.M31 + m1.M22*m2.M32 + m1.M32*m2.M33 + m1.M42*m2.M34,
		m1.M13*m2.M31 + m1.M23*m2.M32 + m1.M33*m2.M33 + m1.M43*m2.M34,
		m1.M14*m2.M31 + m1.M24*m2.M32 + m1.M34*m2.M33 + m1.M44*m2.M34,
		m1.M11*m2.M41 + m1.M21*m2.M42 + m1.M31*m2.M43 + m1.M41*m2.M44,
		m1.M12*m2.M41 + m1.M22*m2.M42 + m1.M32*m2.M43 + m1.M42*m2.M44,
		m1.M13*m2.M41 + m1.M23*m2.M42 + m1.M33*m2.M43 + m1.M43*m2.M44,
		m1.M14*m2.M41 + m1.M24*m2.M42 + m1.M34*m2.M43 + m1.M44*m2.M44}
}

func BenchmarkMatrix4SPointerMul(b *testing.B) {
	m1 := &Matrix4S{}
	m2 := &Matrix4S{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m3 := m1.Mul(m2)
		_ = m3
	}
}

func BenchmarkMatrix4SMul(b *testing.B) {
	m1 := Matrix4S{}
	m2 := Matrix4S{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m3 := m1.MulV(m2)
		_ = m3
	}
}

func BenchmarkMatrix4Mul(b *testing.B) {
	m1 := Matrix4{}
	m2 := Matrix4{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m3 := m1.Mul(m2)
		_ = m3
	}
}

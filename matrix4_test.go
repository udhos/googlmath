package math

import (
	. "gopkg.in/check.v1"
)

type MatrixPerspectiveTestValue struct {
	Fov         float32
	AspectRatio float32
	Near        float32
	Far         float32
	Expected    *Matrix4
}

type MatrixLookAtTestValue struct {
	Eye, Center, Up Vector3
	Expected        *Matrix4
}

type MatrixTranslateTestValue struct {
	Translation Vector3
	Expected    *Matrix4
}

type MatrixRotationTestValue struct {
	Axis     Vector3
	Angle    float32
	Expected *Matrix4
}

type MatrixOrthoTestValue struct {
	Left, Right, Bottom, Top, Near, Far float32
	Expected                            *Matrix4
}

type MatrixMulTestValue struct {
	M1       *Matrix4
	M2       *Matrix4
	Expected *Matrix4
}

type MatrixSetTestValue struct {
	M1       *Matrix4
	M2       *Matrix4
	Expected *Matrix4
}

type MatrixScaleTestValue struct {
	Scalar   Vector3
	Matrix   *Matrix4
	Expected *Matrix4
}

type MatrixInvertTestValue struct {
	Matrix   *Matrix4
	Expected *Matrix4
}

type MatrixDeterminantTestValue struct {
	Matrix   *Matrix4
	Expected float32
}

type Vector4Matrix4TestValue struct {
	Matrix   *Matrix4
	Value    Vector4
	Expected Vector4
}

type ProjectMatrix4TestValue struct {
	Value    Vector3
	Model    *Matrix4
	Proj     *Matrix4
	Viewport Vector4
	Expected Vector3
}

func (s *S) TestMatrixPerspective(c *C) {
	tests := []MatrixPerspectiveTestValue{
		MatrixPerspectiveTestValue{45.0, 4.0 / 3.0, 0.1, 100.0, &Matrix4{1.810660, 0.0, 0.0, 0.0, 0.0, 2.4142134, 0.0, 0.0, 0.0, 0.0, -1.002002, -1.0, 0.0, 0.0, -0.2002002, 0.0}},
		MatrixPerspectiveTestValue{90.0, 16.0 / 9.0, -1.0, 1.0, &Matrix4{0.562500, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0}},
	}
	for _, value := range tests {
		matrix := NewPerspectiveMatrix4(value.Fov, value.AspectRatio, value.Near, value.Far)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixLookAt(c *C) {
	tests := []MatrixLookAtTestValue{
		MatrixLookAtTestValue{Vec3(4, 3, 3), Vec3(0, 0, 0), Vec3(0, 1, 0), &Matrix4{0.600000, -0.411597, 0.685994, 0.0, 0.0, 0.857493, 0.514496, 0.0, -0.800000, -0.308697, 0.514496, 0.0, 0.0, 0.0, -5.830953, 1.0}},
	}
	for _, value := range tests {
		matrix := NewLookAtMatrix4(value.Eye, value.Center, value.Up)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixTranslation(c *C) {
	tests := []MatrixTranslateTestValue{
		MatrixTranslateTestValue{Vec3(0, 0, 15), &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 15.0, 1.0}},
		MatrixTranslateTestValue{Vec3(-3.0, 2.2, 15), &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.200000, 15.0, 1.0}},
	}
	for _, value := range tests {
		translation := value.Translation
		matrix := NewTranslationMatrix4(translation.X, translation.Y, translation.Z)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixRotation(c *C) {
	tests := []MatrixRotationTestValue{
		MatrixRotationTestValue{Vec3(0, 2.5, 0), 25.0, &Matrix4{0.906308, 0.0, -0.422618, 0.0, 0.0, 1.0, 0.0, 0.0, 0.422618, 0.0, 0.906308, 0.0, 0.0, 0.0, 0.0, 1.0}},
		MatrixRotationTestValue{Vec3(2.0, 2.5, 0.0), -45.0, &Matrix4{0.821407, 0.142875, 0.552158, 0.0, 0.142875, 0.885700, -0.441726, 0.0, -0.552158, 0.441726, 0.707107, 0.0, 0.0, 0.0, 0.0, 1.0}},
	}
	for _, value := range tests {
		matrix := NewRotationMatrix4(value.Axis, value.Angle)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixOrtho(c *C) {
	tests := []MatrixOrthoTestValue{
		MatrixOrthoTestValue{-10.0, 10.0, -10.0, 10.0, 0.0, 100.0, &Matrix4{.1, 0.0, 0.0, 0.0, 0.0, 0.1, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -0.0, -0.0, -1.0, 1.0}},
		MatrixOrthoTestValue{0.0, 10.0, 0.0, 10.0, 0.0, 100.0, &Matrix4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0}},
	}
	for _, value := range tests {
		matrix := NewOrthoMatrix4(value.Left, value.Right, value.Bottom, value.Top, value.Near, value.Far)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixMul(c *C) {
	tests := []MatrixMulTestValue{
		MatrixMulTestValue{
			M1:       &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			M2:       &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.200000, 15.0, 1.0},
			Expected: &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.200000, 15.0, 1.0},
		},
	}
	for _, value := range tests {
		matrix := value.M1.Mul(value.M2)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixMulVec4(c *C) {
	tests := []Vector4Matrix4TestValue{
		Vector4Matrix4TestValue{
			&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.2, 15.0, 1.0},
			Vec4(1.0, 2.0, 3.0, 4.0),
			Vec4(-11.0, 10.8, 63.0, 4.0),
		},
		Vector4Matrix4TestValue{
			&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.3, 2.2, 1.3, 1.0},
			Vec4(-1.0, 2.0, 3.0, -4.0),
			Vec4(12.2, -6.8, -2.2, -4.0),
		},
		Vector4Matrix4TestValue{
			&Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Vec4(1.2, 0.2, -3.3, 4.0),
			Vec4(0.675, 0.2, 4.0, 3.3),
		},
		Vector4Matrix4TestValue{
			&Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Vec4(-2.2, 0.0, -3.3, 4.0),
			Vec4(-1.2375, 0.0, 4.0, 3.3),
		},
	}
	for _, value := range tests {
		v := value.Matrix.MulVec4(value.Value)
		c.Check(v, Vector4Check, value.Expected)
	}
}

func (s *S) TestMatrixSet(c *C) {
	tests := []MatrixSetTestValue{
		MatrixSetTestValue{
			M1:       &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			M2:       &Matrix4{-1.0, 2.2, 0.33, 0.22, 0.2, 1.5, 0.1, 0.3, 4.0, 0.2, 1.1, -2.0, 0.1, 2.2, 32.0, 0.0},
			Expected: &Matrix4{-1.0, 2.2, 0.33, 0.22, 0.2, 1.5, 0.1, 0.3, 4.0, 0.2, 1.1, -2.0, 0.1, 2.2, 32.0, 0.0},
		},
	}
	for _, value := range tests {
		matrix := value.M1.Set(value.M2)
		c.Check(matrix, Matrix4Check, value.Expected)
		c.Check(matrix, Matrix4Check, value.M1)
		c.Check(matrix, Matrix4Check, value.M2)
		c.Check(value.M1, Matrix4Check, value.M2)
	}
}

func (s *S) TestMatrixScale(c *C) {
	tests := []MatrixScaleTestValue{
		MatrixScaleTestValue{Vec3(2.0, 3.3, -2.2), NewIdentityMatrix4(), &Matrix4{2.0, 0.0, 0.0, 0.0, 0.0, 3.3, 0.0, 0.0, -0.0, -0.0, -2.2, -0.0, 0.0, 0.0, 0.0, 1.0}},
		MatrixScaleTestValue{Vec3(2.0, 3.3, -2.2), &Matrix4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0}, &Matrix4{0.4, 0.0, 0.0, 0.0, 0.0, 0.66, 0.0, 0.0, -0.0, -0.0, 0.044, -0.0, -1.0, -1.0, -1.0, 1.0}},
	}
	for _, value := range tests {
		matrix := value.Matrix.Scale(value.Scalar)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixInvert(c *C) {
	tests := []MatrixInvertTestValue{
		MatrixInvertTestValue{NewIdentityMatrix4(), NewIdentityMatrix4()},
		MatrixInvertTestValue{&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.2, 15.0, 1.0}, &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 3.0, -2.2, -15.0, 1.0}},
	}
	for _, value := range tests {
		matrix, err := value.Matrix.Invert()
		c.Check(err, IsNil)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (s *S) TestMatrixDeterminant(c *C) {
	tests := []MatrixDeterminantTestValue{
		MatrixDeterminantTestValue{&Matrix4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0}, -0.0008},
		MatrixDeterminantTestValue{NewIdentityMatrix4(), 1.0},
		MatrixDeterminantTestValue{&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.2, 15.0, 1.0}, 1.0},
	}
	for _, value := range tests {
		det := value.Matrix.Determinant()
		c.Check(det, EqualsFloat32, value.Expected)
	}
}

func (s *S) TestProject(c *C) {
	tests := []ProjectMatrix4TestValue{
		ProjectMatrix4TestValue{
			Value:    Vec3(-1.0, 2.0, 3.0),
			Model:    &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.3, 2.2, 1.3, 1.0},
			Proj:     &Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Viewport: Vec4(0.0, 0.0, 1.0, 1.0),
			Expected: Vec3(0.78125, 0.011628, 0.383721),
		},
	}
	for _, value := range tests {
		prj := Project(value.Value, value.Model, value.Proj, value.Viewport)
		c.Check(prj, Vector3Check, value.Expected)
	}
}

func (s *S) TestUnProject(c *C) {
	tests := []ProjectMatrix4TestValue{
		ProjectMatrix4TestValue{
			Value:    Vec3(1.0, 2.0, 3.0),
			Model:    &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 0.965926, 0.258819, 0.0, 0.0, -0.258819, 0.965926, 0.0, 0.0, 0.0, 0.0, 1.0},
			Proj:     &Matrix4{0.562500, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Viewport: Vec4(0.0, 0.0, 16.0, 9.0),
			Expected: Vec3(-0.311111, -0.159089, -0.164428),
		},
		ProjectMatrix4TestValue{
			Value:    Vec3(1.0, 2.0, 3.0),
			Model:    &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			Proj:     &Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Viewport: Vec4(0.0, 0.0, 16.0, 9.0),
			Expected: Vec3(-0.311111, -0.111111, -0.2),
		},
	}

	for _, value := range tests {
		unProj, err := UnProject(value.Value, value.Model, value.Proj, value.Viewport)
		c.Check(err, IsNil)
		c.Check(unProj, Vector3Check, value.Expected)
	}
}

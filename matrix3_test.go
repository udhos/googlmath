package math

import (
	. "gopkg.in/check.v1"
)

type MulMatrix3TestValue struct {
	Matrix   *Matrix3
	Matrix2  *Matrix3
	Expected *Matrix3
}

type DeterminantMatrix3TestValue struct {
	Matrix   *Matrix3
	Expected float32
}

type Matrix3TestValue struct {
	Matrix   *Matrix3
	Expected *Matrix3
}

type ArrayMatrix3TestValue struct {
	Matrix   *Matrix3
	Expected []float32
}

type FloatMatrix3TestValue struct {
	Value    float32
	Expected *Matrix3
}

type MatrixVector3Matrix3TestValue struct {
	Matrix   *Matrix3
	Vector   Vector3
	Expected *Matrix3
}

type MatrixFloatMatrix3TestValue struct {
	Matrix   *Matrix3
	Value    float32
	Expected *Matrix3
}

type Vector2Matrix3TestValue struct {
	Vector   Vector2
	Expected *Matrix3
}

type Vector3FloatMatrix3TestValue struct {
	Vector   Vector3
	Value    float32
	Expected *Matrix3
}

type Matrix3TestSuite struct {
	newXRotationTestTable   []FloatMatrix3TestValue
	newYRotationTestTable   []FloatMatrix3TestValue
	newZRotationTestTable   []FloatMatrix3TestValue
	newRotationTestTable    []Vector3FloatMatrix3TestValue
	newTranslationTestTable []Vector2Matrix3TestValue
	scaleTestTable          []Vector2Matrix3TestValue
	mulTestTable            []MulMatrix3TestValue
	determinantTestTable    []DeterminantMatrix3TestValue
	inverseTestTable        []Matrix3TestValue
	toArrayTestTable        []ArrayMatrix3TestValue
	transposeTestTable      []Matrix3TestValue
	proj2DTestTable         []MatrixVector3Matrix3TestValue
	shearX2DTestTable       []MatrixFloatMatrix3TestValue
	shearY2DTestTable       []MatrixFloatMatrix3TestValue
}

var matrixTest3Suite = Suite(&Matrix3TestSuite{})

func (test *Matrix3TestSuite) SetUpTest(c *C) {
	test.newXRotationTestTable = []FloatMatrix3TestValue{
		FloatMatrix3TestValue{
			30,
			&Matrix3{1.0, 0, 0, 0, 0.866, -0.5, 0, 0.5, 0.866},
		},
		FloatMatrix3TestValue{
			-99,
			&Matrix3{1.0, 0, 0, 0, -Sin(Pi / 20), Cos(Pi / 20), 0, -Cos(Pi / 20), -Sin(Pi / 20)},
		},
	}

	test.newYRotationTestTable = []FloatMatrix3TestValue{
		FloatMatrix3TestValue{
			-99,
			&Matrix3{-Sin(Pi / 20), 0, -Cos(Pi / 20), 0, 1, 0, Cos(Pi / 20), 0, -Sin(Pi / 20)},
		},
		FloatMatrix3TestValue{
			11,
			&Matrix3{Cos(11 * Pi / 180), 0, Sin(11 * Pi / 180), 0, 1, 0, -Sin(11 * Pi / 180), 0, Cos(11 * Pi / 180)},
		},
	}

	test.newZRotationTestTable = []FloatMatrix3TestValue{
		FloatMatrix3TestValue{
			-99,
			&Matrix3{-Sin(Pi / 20), Cos(Pi / 20), 0, -Cos(Pi / 20), -Sin(Pi / 20), 0, 0, 0, 1},
		},
		FloatMatrix3TestValue{
			11,
			&Matrix3{Cos(11 * Pi / 180), -Sin(11 * Pi / 180), 0, Sin(11 * Pi / 180), Cos(11 * Pi / 180), 0, 0, 0, 1},
		},
	}

	test.newRotationTestTable = []Vector3FloatMatrix3TestValue{
		Vector3FloatMatrix3TestValue{
			Vec3(0, 2.5, 0),
			25.0,
			&Matrix3{0.906308, 0.0, -0.422618, 0.0, 1.0, 0.0, 0.422618, 0.0, 0.906308},
		},
		Vector3FloatMatrix3TestValue{
			Vec3(2.0, 2.5, 0.0),
			-45.0,
			&Matrix3{0.821407, 0.142875, 0.552158, 0.142875, 0.885700, -0.441726, -0.552158, 0.441726, 0.707107},
		},
	}

	test.newTranslationTestTable = []Vector2Matrix3TestValue{
		Vector2Matrix3TestValue{
			Vec2(1.2, 2.0),
			&Matrix3{
				1.0, 0.0, 0.0,
				0.0, 1.0, 0.0,
				1.2, 2.0, 1.0,
			},
		},
		Vector2Matrix3TestValue{
			Vec2(-2.2, 3.3),
			&Matrix3{
				1.0, 0.0, 0.0,
				0.0, 1.0, 0.0,
				-2.2, 3.3, 1.0,
			},
		},
	}

	test.scaleTestTable = []Vector2Matrix3TestValue{
		Vector2Matrix3TestValue{
			Vec2(1.2, -2.0),
			&Matrix3{1.2, 0.0, 0.0, -0.0, -2.0, -0.0, 0.0, 0.0, 1.0},
		},
		Vector2Matrix3TestValue{
			Vec2(-1.2, -2.3),
			&Matrix3{-1.2, -0.0, -0.0, -0.0, -2.3, -0.0, 0.0, 0.0, 1.0},
		},
	}

	test.mulTestTable = []MulMatrix3TestValue{
		MulMatrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			&Matrix3{1.0, 0.0, 0.0, 0.0, 2.5, 0.0, 2.2, 0.2, 1.0},
			&Matrix3{11.0, 0.0, 3.0, -7.5, 6.25, 0.0, 25.800001, 0.8, 18.6},
		},
		MulMatrix3TestValue{
			&Matrix3{11.0, 0.34, -3.0, -3.0, 2.5, 0.0, 32.200001, 3.32, 12.0},
			&Matrix3{1.0, 0.0, 3.0, 5.0, -1.5, -2.0, -2.2, 0.2, 2.0},
			&Matrix3{107.600006, 10.3, 33.0, -4.900002, -8.69, -39.0, 39.599998, 6.392, 30.6},
		},
	}

	test.determinantTestTable = []DeterminantMatrix3TestValue{
		DeterminantMatrix3TestValue{
			&Matrix3{11.0, 0.34, -3.0, -3.0, 2.5, 0.0, 32.200001, 3.32, 12.0},
			613.619995,
		},
		DeterminantMatrix3TestValue{
			&Matrix3{1.0, 0.0, 3.0, 5.0, -1.5, -2.0, -2.2, 0.2, 2.0},
			-9.5,
		},
	}

	test.inverseTestTable = []Matrix3TestValue{
		Matrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			&Matrix3{0.096525, 0.002896, -0.024131, 0.11583, 0.403475, -0.028958, -0.020592, -0.010618, 0.088481},
		},
	}

	test.toArrayTestTable = []ArrayMatrix3TestValue{
		ArrayMatrix3TestValue{
			&Matrix3{1, 0, 0, 0, 1, 0, 0, 0, 1},
			[]float32{1, 0, 0, 0, 1, 0, 0, 0, 1},
		},
		ArrayMatrix3TestValue{
			&Matrix3{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]float32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	test.transposeTestTable = []Matrix3TestValue{
		Matrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			&Matrix3{11.0, -3.0, 2.2, 0.0, 2.5, 0.3, 3.0, 0.0, 12.0},
		},
	}

	test.proj2DTestTable = []MatrixVector3Matrix3TestValue{
		MatrixVector3Matrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			Vec3(1, 2, 3),
			&Matrix3{6.0, -5.0, 0.0, -13.0, -7.5, -6.0, 2.2, 0.3, 12.0},
		},
	}

	test.shearX2DTestTable = []MatrixFloatMatrix3TestValue{
		MatrixFloatMatrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			1.2,
			&Matrix3{7.4, 3.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
		},
		MatrixFloatMatrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			-1.2,
			&Matrix3{14.6, -3.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
		},
	}

	test.shearY2DTestTable = []MatrixFloatMatrix3TestValue{
		MatrixFloatMatrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			1.2,
			&Matrix3{11.0, 0.0, 3.0, 10.200001, 2.5, 3.6, 2.2, 0.3, 12.0},
		},
		MatrixFloatMatrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			-1.2,
			&Matrix3{11.0, 0.0, 3.0, -16.200001, 2.5, -3.6, 2.2, 0.3, 12.0},
		},
	}
}

func (test *Matrix3TestSuite) TestNewMatrix3(c *C) {
	var m11, m12, m13, m21, m22, m23, m31, m32, m33 float32 = 1, 2, 3, 4, 5, 6, 7, 8, 9
	m := NewMatrix3(m11, m12, m13, m21, m22, m23, m31, m32, m33)
	m2 := &Matrix3{m11, m12, m13, m21, m22, m23, m31, m32, m33}
	c.Check(m, Matrix3Check, m2)
}

func (test *Matrix3TestSuite) TestNewIdentityMatrix3(c *C) {
	expected := &Matrix3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	obtained := NewIdentityMatrix3()
	c.Check(obtained, Matrix3Check, expected)
}

func (test *Matrix3TestSuite) TestNewXRotationMatrix3(c *C) {
	for _, value := range test.newXRotationTestTable {
		matrix := NewXRotationMatrix3(value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewYRotationMatrix3(c *C) {
	for _, value := range test.newYRotationTestTable {
		matrix := NewYRotationMatrix3(value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewZRotationMatrix3(c *C) {
	for _, value := range test.newZRotationTestTable {
		matrix := NewZRotationMatrix3(value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewRotationMatrix3(c *C) {
	for _, value := range test.newRotationTestTable {
		matrix := NewRotationMatrix3(value.Vector, value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewTranslationMatrix3(c *C) {
	for _, value := range test.newTranslationTestTable {
		matrix := NewTranslationMatrix3(value.Vector.X, value.Vector.Y)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewScaleMatrix3(c *C) {
	for _, value := range test.scaleTestTable {
		matrix := NewScaleMatrix3(value.Vector.X, value.Vector.Y)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestSet(c *C) {
	m := &Matrix3{}
	m2 := &Matrix3{0, 1, 2, 3, 4, 5, 6, 7, 8}
	m.Set(m2)
	c.Check(m, Matrix3Check, m2)
}

func (test *Matrix3TestSuite) TestMul(c *C) {
	for _, value := range test.mulTestTable {
		matrix := value.Matrix.Mul(value.Matrix2)
		c.Check(value.Matrix, Not(Matrix3Check), value.Matrix2)
		c.Check(value.Matrix, Not(Matrix3Check), matrix)
		c.Check(value.Matrix2, Not(Matrix3Check), matrix)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestDeterminant(c *C) {
	for _, value := range test.determinantTestTable {
		det := value.Matrix.Determinant()
		c.Check(det, EqualsFloat32, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestInverse(c *C) {
	for _, value := range test.inverseTestTable {
		inv, err := value.Matrix.Inverse()
		c.Check(err, IsNil)
		c.Check(inv, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestToArray(c *C) {
	for _, value := range test.toArrayTestTable {
		a := value.Matrix.ToArray()
		c.Check(a, DeepEquals, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestTranspose(c *C) {
	for _, value := range test.transposeTestTable {
		m := value.Matrix.Transpose()
		c.Check(m, DeepEquals, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestProj2D(c *C) {
	for _, value := range test.proj2DTestTable {
		m := value.Matrix.Proj2D(value.Vector)
		c.Check(m, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestShearX2D(c *C) {
	for _, value := range test.shearX2DTestTable {
		m := value.Matrix.ShearX2D(value.Value)
		c.Check(m, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestShearY2D(c *C) {
	for _, value := range test.shearY2DTestTable {
		m := value.Matrix.ShearY2D(value.Value)
		c.Check(m, Matrix3Check, value.Expected)
	}
}

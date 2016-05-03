package geo

import (
	"math"
)

// 3x3 column major matrix.
type Mat3 struct {
	Values [9]float64
}

// Creates a new 3x3 matrix with initial values.
func NewMat3(m00, m10, m20, m01, m11, m21, m02, m12, m22 float64) *Mat3 {
	return &Mat3{[9]float64{m00, m10, m20, m01, m11, m21, m02, m12, m22}}
}

// Creates a copy of actual matrix.
func (m *Mat3) Copy() *Mat3 {
	return &Mat3{m.Values}
}

// Sets the matrix to zeros.
func (m *Mat3) Clear() {
	for i := range m.Values {
		m.Values[i] = 0
	}
}

// Sets the matrix to identity matrix.
func (m *Mat3) Identity() {
	m.Clear()
	m.Values[0] = 1
	m.Values[4] = 1
	m.Values[8] = 1
}

// Multiplies actual matrix with given matrix and saves result.
func (m *Mat3) Mult(mat *Mat3) {
	this := m.Copy()

	m.Values[0] = mat.Values[0]*this.Values[0] + mat.Values[1]*this.Values[3] + mat.Values[2]*this.Values[6]
	m.Values[1] = mat.Values[0]*this.Values[1] + mat.Values[1]*this.Values[4] + mat.Values[2]*this.Values[7]
	m.Values[2] = mat.Values[0]*this.Values[2] + mat.Values[1]*this.Values[5] + mat.Values[2]*this.Values[8]

	m.Values[3] = mat.Values[3]*this.Values[0] + mat.Values[4]*this.Values[3] + mat.Values[5]*this.Values[6]
	m.Values[4] = mat.Values[3]*this.Values[1] + mat.Values[4]*this.Values[4] + mat.Values[5]*this.Values[7]
	m.Values[5] = mat.Values[3]*this.Values[2] + mat.Values[4]*this.Values[5] + mat.Values[5]*this.Values[8]

	m.Values[6] = mat.Values[6]*this.Values[0] + mat.Values[7]*this.Values[3] + mat.Values[8]*this.Values[6]
	m.Values[7] = mat.Values[6]*this.Values[1] + mat.Values[7]*this.Values[4] + mat.Values[8]*this.Values[7]
	m.Values[8] = mat.Values[6]*this.Values[2] + mat.Values[7]*this.Values[5] + mat.Values[8]*this.Values[8]
}

// Multiplies given vector with actual matrix and returns result.
func (m *Mat3) MultVec(v Vec3) Vec3 {
	vec := Vec3{}

	vec.X = m.Values[0]*v.X + m.Values[3]*v.X + m.Values[6]*v.X
	vec.Y = m.Values[1]*v.Y + m.Values[4]*v.Y + m.Values[7]*v.Y
	vec.Z = m.Values[2]*v.Z + m.Values[5]*v.Z + m.Values[8]*v.Z

	return vec
}

// Returns the determinate of actual matrix.
func (m *Mat3) Determinate() float64 {
	var d float64

	d = m.Values[0]*m.Values[4]*m.Values[8] + m.Values[3]*m.Values[7]*m.Values[2] + m.Values[6]*m.Values[1]*m.Values[5]
	d -= m.Values[2]*m.Values[4]*m.Values[6] - m.Values[5]*m.Values[7]*m.Values[0] - m.Values[8]*m.Values[1]*m.Values[3]

	return d
}

// Sets the inverse of actual matrix.
func (m *Mat3) Inverse() {
	d := 1 / m.Determinate()
	mat := m.Copy()

	m.Values[0] = (mat.Values[4]*mat.Values[8] - mat.Values[7]*mat.Values[5]) * d
	m.Values[1] = (mat.Values[7]*mat.Values[2] - mat.Values[1]*mat.Values[8]) * d
	m.Values[2] = (mat.Values[1]*mat.Values[5] - mat.Values[4]*mat.Values[2]) * d

	m.Values[3] = (mat.Values[6]*mat.Values[5] - mat.Values[3]*mat.Values[8]) * d
	m.Values[4] = (mat.Values[0]*mat.Values[8] - mat.Values[6]*mat.Values[2]) * d
	m.Values[5] = (mat.Values[3]*mat.Values[2] - mat.Values[0]*mat.Values[5]) * d

	m.Values[6] = (mat.Values[3]*mat.Values[7] - mat.Values[6]*mat.Values[4]) * d
	m.Values[7] = (mat.Values[6]*mat.Values[1] - mat.Values[0]*mat.Values[7]) * d
	m.Values[8] = (mat.Values[0]*mat.Values[4] - mat.Values[3]*mat.Values[1]) * d
}

// Calculates and saves the transpose of actual matrix.
func (m *Mat3) Transpose() {
	mat := m.Copy()

	m.Values[1] = mat.Values[3]
	m.Values[2] = mat.Values[6]
	m.Values[3] = mat.Values[1]

	m.Values[5] = mat.Values[7]
	m.Values[6] = mat.Values[2]
	m.Values[7] = mat.Values[5]
}

// Translates and saves actual matrix by given vector.
func (m *Mat3) Translate(v Vec2) {
	mat := &Mat3{}
	mat.Identity()

	mat.Values[6] = v.X
	mat.Values[7] = v.Y

	m.Mult(mat)
}

// Scales and saves actual matrix by given vector.
func (m *Mat3) Scale(v Vec2) {
	mat := &Mat3{}
	mat.Identity()

	mat.Values[0] = v.X
	mat.Values[4] = v.Y

	m.Mult(mat)
}

// Rotates and saves actual matrix by given vector.
func (m *Mat3) Rotate(angle float64) {
	mat := &Mat3{}
	var co, si float64

	angle = angle * (math.Pi / 180)
	si = math.Sin(angle)
	co = math.Cos(angle)

	mat.Values[0] = co
	mat.Values[1] = si
	mat.Values[2] = 0

	mat.Values[3] = -si
	mat.Values[4] = co
	mat.Values[5] = 0

	mat.Values[6] = 0
	mat.Values[7] = 0
	mat.Values[8] = 1

	m.Mult(mat)
}

// Sets actual matrix to orthogonal projection with given viewport.
func (m *Mat3) Ortho(viewport Vec4) {
	if viewport.X != viewport.Z && viewport.Y != viewport.W {
		m.Identity()

		m.Values[0] = 2 / (viewport.Z - viewport.X)
		m.Values[4] = 2 / (viewport.W - viewport.Y)
		m.Values[6] = -(viewport.Z + viewport.X) / (viewport.Z - viewport.X)
		m.Values[7] = -(viewport.W + viewport.Y) / (viewport.W - viewport.Y)
		m.Values[8] = 1
	}
}

// Multiplies both matrices and returns a new Mat3.
func MultMat3(a, b *Mat3) *Mat3 {
	c := a.Copy()
	c.Mult(b)

	return c
}

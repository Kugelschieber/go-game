package geo

import (
	"math"
)

// 4x4 column major matrix.
type Mat4 struct {
	Values [16]float64
}

// Creates a new 4x4 matrix with initial values.
func NewMat4(m00, m10, m20, m30, m01, m11, m21, m31, m02, m12, m22, m32, m03, m13, m23, m33 float64) *Mat4 {
	return &Mat4{[16]float64{m00, m10, m20, m30, m01, m11, m21, m31, m02, m12, m22, m32, m03, m13, m23, m33}}
}

// Creates a copy of actual matrix.
func (m *Mat4) Copy() *Mat4 {
	return &Mat4{m.Values}
}

// Sets the matrix to zeros.
func (m *Mat4) Clear() {
	for i := range m.Values {
		m.Values[i] = 0
	}
}

// Sets the matrix to identity matrix.
func (m *Mat4) Identity() {
	m.Clear()
	m.Values[0] = 1
	m.Values[5] = 1
	m.Values[10] = 1
	m.Values[15] = 1
}

// Multiplies actual matrix with given matrix and saves result.
func (m *Mat4) Mult(mat *Mat4) {
	this := m.Copy()

	m.Values[0] = mat.Values[0]*this.Values[0] + mat.Values[1]*this.Values[4] + mat.Values[2]*this.Values[8] + mat.Values[3]*this.Values[12]
	m.Values[1] = mat.Values[0]*this.Values[1] + mat.Values[1]*this.Values[5] + mat.Values[2]*this.Values[9] + mat.Values[3]*this.Values[13]
	m.Values[2] = mat.Values[0]*this.Values[2] + mat.Values[1]*this.Values[6] + mat.Values[2]*this.Values[10] + mat.Values[3]*this.Values[14]
	m.Values[3] = mat.Values[0]*this.Values[3] + mat.Values[1]*this.Values[7] + mat.Values[2]*this.Values[11] + mat.Values[3]*this.Values[15]

	m.Values[4] = mat.Values[4]*this.Values[0] + mat.Values[5]*this.Values[4] + mat.Values[6]*this.Values[8] + mat.Values[7]*this.Values[12]
	m.Values[5] = mat.Values[4]*this.Values[1] + mat.Values[5]*this.Values[5] + mat.Values[6]*this.Values[9] + mat.Values[7]*this.Values[13]
	m.Values[6] = mat.Values[4]*this.Values[2] + mat.Values[5]*this.Values[6] + mat.Values[6]*this.Values[10] + mat.Values[7]*this.Values[14]
	m.Values[7] = mat.Values[4]*this.Values[3] + mat.Values[5]*this.Values[7] + mat.Values[6]*this.Values[11] + mat.Values[7]*this.Values[15]

	m.Values[8] = mat.Values[8]*this.Values[0] + mat.Values[9]*this.Values[4] + mat.Values[10]*this.Values[8] + mat.Values[11]*this.Values[12]
	m.Values[9] = mat.Values[8]*this.Values[1] + mat.Values[9]*this.Values[5] + mat.Values[10]*this.Values[9] + mat.Values[11]*this.Values[13]
	m.Values[10] = mat.Values[8]*this.Values[2] + mat.Values[9]*this.Values[6] + mat.Values[10]*this.Values[10] + mat.Values[11]*this.Values[14]
	m.Values[11] = mat.Values[8]*this.Values[3] + mat.Values[9]*this.Values[7] + mat.Values[10]*this.Values[11] + mat.Values[11]*this.Values[15]

	m.Values[12] = mat.Values[12]*this.Values[0] + mat.Values[13]*this.Values[4] + mat.Values[14]*this.Values[8] + mat.Values[15]*this.Values[12]
	m.Values[13] = mat.Values[12]*this.Values[1] + mat.Values[13]*this.Values[5] + mat.Values[14]*this.Values[9] + mat.Values[15]*this.Values[13]
	m.Values[14] = mat.Values[12]*this.Values[2] + mat.Values[13]*this.Values[6] + mat.Values[14]*this.Values[10] + mat.Values[15]*this.Values[14]
	m.Values[15] = mat.Values[12]*this.Values[3] + mat.Values[13]*this.Values[7] + mat.Values[14]*this.Values[11] + mat.Values[15]*this.Values[15]
}

// Multiplies given vector with actual matrix and returns result.
func (m *Mat4) MultVec(v Vec3) Vec3 {
	vec := Vec3{}

	vec.X = m.Values[0]*v.X + m.Values[4]*v.X + m.Values[8]*v.X + m.Values[12]*v.X
	vec.Y = m.Values[1]*v.Y + m.Values[5]*v.Y + m.Values[9]*v.Y + m.Values[13]*v.Y
	vec.Z = m.Values[2]*v.Z + m.Values[6]*v.Z + m.Values[10]*v.Z + m.Values[14]*v.Z

	return vec
}

// Returns the determinate of actual matrix.
func (m *Mat4) Determinate() float64 {
	var d float64

	d = m.Values[0]*m.Values[4]*m.Values[8] + m.Values[1]*m.Values[5]*m.Values[6] + m.Values[2]*m.Values[3]*m.Values[7]
	d -= m.Values[2]*m.Values[4]*m.Values[6] + m.Values[0]*m.Values[5]*m.Values[7] + m.Values[1]*m.Values[3]*m.Values[8]

	return d
}

// Sets the inverse of actual matrix.
func (m *Mat4) Inverse() {
	mat := m.Copy()

	m.Values[0] = mat.Values[0]
	m.Values[1] = mat.Values[4]
	m.Values[2] = mat.Values[8]
	m.Values[4] = mat.Values[1]
	m.Values[6] = mat.Values[9]
	m.Values[8] = mat.Values[2]
	m.Values[9] = mat.Values[6]

	m.Values[12] = m.Values[0]*-mat.Values[12] + m.Values[4]*-mat.Values[13] + m.Values[8]*-mat.Values[14]
	m.Values[13] = m.Values[1]*-mat.Values[12] + m.Values[5]*-mat.Values[13] + m.Values[9]*-mat.Values[14]
	m.Values[14] = m.Values[2]*-mat.Values[12] + m.Values[6]*-mat.Values[13] + m.Values[10]*-mat.Values[14]

	m.Values[3] = 0
	m.Values[7] = 0
	m.Values[11] = 0
	m.Values[15] = 1
}

// Calculates and saves the transpose of actual matrix.
func (m *Mat4) Transpose() {
	mat := m.Copy()

	m.Values[1] = mat.Values[4]
	m.Values[2] = mat.Values[8]
	m.Values[3] = mat.Values[12]

	m.Values[4] = mat.Values[1]
	m.Values[6] = mat.Values[9]
	m.Values[7] = mat.Values[13]

	m.Values[8] = mat.Values[2]
	m.Values[9] = mat.Values[6]
	m.Values[11] = mat.Values[14]

	m.Values[12] = mat.Values[3]
	m.Values[13] = mat.Values[2]
	m.Values[14] = mat.Values[11]
}

// Translates and saves actual matrix by given vector.
func (m *Mat4) Translate(v Vec3) {
	mat := Mat4{}
	mat.Identity()

	mat.Values[12] = v.X
	mat.Values[13] = v.Y
	mat.Values[14] = v.Z

	m.Mult(&mat)
}

// Scales and saves actual matrix by given vector.
func (m *Mat4) Scale(v Vec3) {
	mat := Mat4{}
	mat.Identity()

	mat.Values[0] = v.X
	mat.Values[5] = v.Y
	mat.Values[10] = v.Z

	m.Mult(&mat)
}

// Rotates and saves actual matrix by given vector.
func (m *Mat4) Rotate(angle float64, axis Vec3) {
	mat := Mat4{}
	var co, si float64

	axis.Normalize()
	angle = angle * (math.Pi / 180)
	si = float64(math.Sin(float64(angle)))
	co = float64(math.Cos(float64(angle)))

	mat.Values[0] = axis.X*axis.X*(1-co) + co
	mat.Values[1] = axis.Y*axis.X*(1-co) + axis.Z*si
	mat.Values[2] = axis.X*axis.Z*(1-co) - axis.Y*si
	mat.Values[3] = 0

	mat.Values[4] = axis.X*axis.Y*(1-co) - axis.Z*si
	mat.Values[5] = axis.Y*axis.Y*(1-co) + co
	mat.Values[6] = axis.Y*axis.Z*(1-co) + axis.X*si
	mat.Values[7] = 0

	mat.Values[8] = axis.X*axis.Z*(1-co) + axis.Y*si
	mat.Values[9] = axis.Y*axis.Z*(1-co) - axis.X*si
	mat.Values[10] = axis.Z*axis.Z*(1-co) + co
	mat.Values[11] = 0

	mat.Values[12] = 0
	mat.Values[13] = 0
	mat.Values[14] = 0
	mat.Values[15] = 1

	m.Mult(&mat)
}

// Sets actual matrix to orthogonal projection with given viewport.
func (m *Mat4) Ortho(viewport Vec4, znear, zfar float64) {
	if viewport.X != viewport.Z && viewport.Y != viewport.W && znear != zfar {
		m.Identity()

		m.Values[0] = 2 / (viewport.Z - viewport.X)
		m.Values[5] = 2 / (viewport.W - viewport.Y)
		m.Values[10] = -2 / (zfar - znear)
		m.Values[12] = -(viewport.Z + viewport.X) / (viewport.Z - viewport.X)
		m.Values[13] = -(viewport.W + viewport.Y) / (viewport.W - viewport.Y)
		m.Values[14] = -(zfar + znear) / (zfar - znear)
		m.Values[15] = 1
	}
}

// Sets actual matrix to project to specified locations with given upper axis.
func (m *Mat4) LookAt(pos, lookAt, up Vec3) {
	dir := lookAt.Copy()
	dir.Sub(pos)
	dir.Normalize()

	right := CrossVec3(dir, up)
	right.Normalize()

	up = CrossVec3(right, dir)
	up.Normalize()

	mat := Mat4{}

	mat.Values[0] = right.X
	mat.Values[4] = right.Y
	mat.Values[8] = right.Z
	mat.Values[12] = -right.DotVec(pos)

	mat.Values[1] = up.X
	mat.Values[5] = up.Y
	mat.Values[9] = up.Z
	mat.Values[13] = -up.DotVec(pos)

	mat.Values[2] = -dir.X
	mat.Values[6] = -dir.Y
	mat.Values[10] = -dir.Z
	mat.Values[14] = dir.DotVec(pos)

	mat.Values[3] = 0
	mat.Values[7] = 0
	mat.Values[11] = 0
	mat.Values[15] = 1

	m.Mult(&mat)
}

// Sets actual matrix to perspective projection with given viewport.
func (m *Mat4) Perspective(fov, ratio, znear, zfar float64) {
	f := 1 / math.Tan(float64(fov*(math.Pi/360)))
	m.Identity()

	m.Values[0] = f / ratio
	m.Values[5] = f
	m.Values[10] = (zfar + znear) / (znear - zfar)
	m.Values[11] = -1
	m.Values[14] = (2 * zfar * znear) / (znear - zfar)
	m.Values[15] = 0
}

// Multiplies both matrices and returns a new Mat4.
func MultMat4(a, b *Mat4) *Mat4 {
	c := a.Copy()
	c.Mult(b)

	return c
}

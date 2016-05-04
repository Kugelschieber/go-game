package goga

import (
	"math"
)

// A 3D vector.
type Vec3 struct {
	X, Y, Z float64
}

// Creates a copy of actual vector.
func (v *Vec3) Copy() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

// Adds and saves the given vector to actual vector.
func (v *Vec3) Add(vec Vec3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

// Subtracts and saves the given vector to actual vector.
func (v *Vec3) Sub(vec Vec3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

// Multiplies and saves the given vector to actual vector.
func (v *Vec3) Mult(vec Vec3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

// Divides and saves the given vector to actual vector.
func (v *Vec3) Div(vec Vec3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}

// Calculates and returns the dot product of actual vector.
func (v *Vec3) Dot() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Calculates and returns the dot product of combination of given vector and actual vector.
func (v *Vec3) DotVec(vec Vec3) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Returns the length of actual vector.
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.Dot())
}

// Normalizes actual vector to length 1.
func (v *Vec3) Normalize() {
	l := v.Length()

	v.X /= l
	v.Y /= l
	v.Z /= l
}

// Calculates and saves cross product of given and actual vector.
func (v *Vec3) Cross(vec Vec3) {
	this := Vec3{v.X, v.Y, v.Z}

	v.X = this.Y*vec.Z - this.Z*vec.Y
	v.Y = this.Z*vec.X - this.X*vec.Z
	v.Z = this.X*vec.Y - this.Y*vec.X
}

// Calulates the cross product of given vectors and returns result as a new vector.
func CrossVec3(a, b Vec3) Vec3 {
	vec := Vec3{}

	vec.X = (a.Y * b.Z) - (a.Z * b.Y)
	vec.Y = (a.Z * b.X) - (a.X * b.Z)
	vec.Z = (a.X * b.Y) - (a.Y * b.X)

	return vec
}

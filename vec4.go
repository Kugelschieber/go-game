package goga

import (
	"math"
)

// A 4D vector.
type Vec4 struct {
	X, Y, Z, W float64
}

// Creates a copy of actual vector.
func (v *Vec4) Copy() Vec4 {
	return Vec4{v.X, v.Y, v.Z, v.W}
}

// Adds and saves the given vector to actual vector.
func (v *Vec4) Add(vec Vec4) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
	v.W += vec.W
}

// Subracts and saves the given vector to actual vector.
func (v *Vec4) Sub(vec Vec4) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
	v.W -= vec.W
}

// Multiplies and saves the given vector to actual vector.
func (v *Vec4) Mult(vec Vec4) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
	v.W *= vec.W
}

// Divides and saves the given vector to actual vector.
func (v *Vec4) Div(vec Vec4) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
	v.W /= vec.W
}

// Calculates and returns the dot product of actual vector.
func (v *Vec4) Dot() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

// Calculates and returns the dot product of combination of given vector and actual vector.
func (v *Vec4) DotVec(vec Vec4) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z + v.W*vec.W
}

// Returns the length of actual vector.
func (v *Vec4) Length() float64 {
	return math.Sqrt(v.Dot())
}

// Normalizes actual vector to length 1.
func (v *Vec4) Normalize() {
	l := v.Length()

	v.X /= l
	v.Y /= l
	v.Z /= l
	v.W /= l
}

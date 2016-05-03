package geo

import (
	"math"
)

// A 2D vector.
type Vec2 struct {
	X, Y float64
}

// Creates a copy of actual vector.
func (v *Vec2) Copy() Vec2 {
	return Vec2{v.X, v.Y}
}

// Adds and saves the given vector to actual vector.
func (v *Vec2) Add(vec Vec2) {
	v.X += vec.X
	v.Y += vec.Y
}

// Subtracts and saves the given vector to actual vector.
func (v *Vec2) Sub(vec Vec2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

// Multiplies and saves the given vector to actual vector.
func (v *Vec2) Mult(vec Vec2) {
	v.X *= vec.X
	v.Y *= vec.Y
}

// Divides and saves the given vector to actual vector.
func (v *Vec2) Div(vec Vec2) {
	v.X /= vec.X
	v.Y /= vec.Y
}

// Calculates and returns the dot product of actual vector.
func (v *Vec2) Dot() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Calculates and returns the dot product of combination of given vector and actual vector.
func (v *Vec2) DotVec(vec Vec2) float64 {
	return v.X*vec.X + v.Y*vec.Y
}

// Returns the length of actual vector.
func (v *Vec2) Length() float64 {
	return float64(math.Sqrt(float64(v.Dot())))
}

// Normalizes actual vector to length 1.
func (v *Vec2) Normalize() {
	l := v.Length()

	v.X /= l
	v.Y /= l
}

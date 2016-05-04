package goga

import (
	"math"
)

// Returns the distance between two 2D vectors.
func DistanceVec2(a, b Vec2) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

// Returns the distance between two 3D vectors.
func DistanceVec3(a, b Vec3) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z)))
}

// Returns the distance between two 4D vectors.
func DistanceVec4(a, b Vec4) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z) + (a.W-b.W)*(a.W-b.W)))
}

// Package geometry implements 3D vector math operations.
//
// It provides the Vec3 type for representing three-dimensional vectors
// and implements common vector operations such as addition, subtraction,
// scaling, dot product, cross product, normalization, and more.
// Both in-place modifications and functions that return new vector instances
// are available to suit different programming needs.
package geometry

import (
	"fmt"
	"math"
)

// Vec3 represents a three-dimensional vector with X, Y, and Z components.
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// ZERO_VEC3 is the zero vector (0, 0, 0).
// UNIT_X, UNIT_Y, and UNIT_Z are the unit vectors along the X, Y, and Z axes respectively.
var (
	ZERO_VEC3 = NewVec3(0, 0, 0)
	UNIT_X    = NewVec3(1, 0, 0)
	UNIT_Y    = NewVec3(0, 1, 0)
	UNIT_Z    = NewVec3(0, 0, 1)
)

// NewVec3 creates and returns a new Vec3 initialized with the given x, y, and z components.
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// Add adds the specified vector to the current vector, modifying it in place.
func (v *Vec3) Add(v2 Vec3) {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
}

// Add returns a new Vec3 that is the element-wise sum of the two provided vectors.
func Add(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// Sub subtracts the specified vector from the current vector, modifying it in place.
func (v *Vec3) Sub(v2 Vec3) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}

// Sub returns a new Vec3 that is the element-wise difference between the two provided vectors.
func Sub(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// Mul multiplies the current vector by the given scalar, modifying it in place.
func (v *Vec3) Mul(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

// Mul returns a new Vec3 resulting from scaling the provided vector by the given scalar.
func Mul(v1 Vec3, scalar float64) Vec3 {
	return Vec3{v1.X * scalar, v1.Y * scalar, v1.Z * scalar}
}

// Div divides the current vector by the provided scalar, modifying it in place.
func (v *Vec3) Div(scalar float64) {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
}

// Div returns a new Vec3 resulting from dividing the provided vector by the given scalar.
func Div(v1 Vec3, scalar float64) Vec3 {
	return Vec3{v1.X / scalar, v1.Y / scalar, v1.Z / scalar}
}

// Dot computes and returns the dot product of the current vector with the given vector.
func (v *Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Dot returns the dot product of the two provided vectors.
func Dot(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Negate inverts the current vector (i.e., multiplies all components by -1) in place.
func (v *Vec3) Negate() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Neg returns a new Vec3 that is the negation of the current vector.
func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Length calculates and returns the magnitude (length) of the current vector.
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.Dot(*v))
}

// Length calculates and returns the magnitude (length) of the provided vector.
func Length(v Vec3) float64 {
	return math.Sqrt(v.Dot(v))
}

// SqrLength returns the squared length of the current vector (avoiding the cost of a square root).
func (v *Vec3) SqrLength() float64 {
	return v.Dot(*v)
}

// Normalize adjusts the current vector to have a unit length (normalizes it) in place.
func (v *Vec3) Normalize() {
	v.Div(v.Length())
}

// Normal returns a normalized copy of the current vector (i.e., with a length of 1).
func (v Vec3) Normal() Vec3 {
	v.Normalize()
	return v
}

// IsNormalized checks if the current vector is approximately of unit length,
// returning true if the squared length is nearly 1 within a small epsilon.
func (v *Vec3) IsNormalized() bool {
	return math.Abs(v.SqrLength()-1) < 0.00001
}

// Copy returns a duplicate of the current vector.
func (v *Vec3) Copy() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

// Cross computes and returns a new Vec3 that is the cross product
// of the current vector with the specified vector.
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

// Cross returns the cross product of the two provided vectors.
func Cross(v1, v2 Vec3) Vec3 {
	return Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

// String returns a string representation of the vector in the format "(X, Y, Z)".
func (v Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

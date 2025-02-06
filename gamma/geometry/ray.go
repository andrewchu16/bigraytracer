package geometry

import "fmt"

type Ray struct {
	orig Vec3
	dir  Vec3
}

func NewRay(origin, direction Vec3) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Origin() Vec3 {
	return r.orig
}

func (r *Ray) Direction() Vec3 {
	return r.dir
}

func (r *Ray) At(t float64) Vec3 {
	return Add(r.orig, Mul(r.dir, t))
}

func (r *Ray) String() string {
	return fmt.Sprintf("Ray(Origin: %v, Direction: %v)", r.orig, r.dir)
}

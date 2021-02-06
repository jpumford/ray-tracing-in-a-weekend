package main

import "math"

type vec3 struct {
	x float64
	y float64
	z float64
}

func (x *vec3) dot(y *vec3) float64 {
	return x.x * y.x + x.y * y.y + x.z * y.z
}

func (x *vec3) mult(y float64) *vec3 {
	return &vec3{
		x.x*y,
		x.y*y,
		x.z*y,
	}
}

func (x *vec3) add(y *vec3) *vec3 {
	return &vec3{
		x.x + y.x,
		x.y + y.y,
		x.z + y.z,
	}
}

func (x *vec3) sub(y *vec3) *vec3 {
	return &vec3{
		x.x - y.x,
		x.y - y.y,
		x.z - y.z,
	}
}

func (x *vec3) norm() float64 {
	sum := x.x * x.x + x.y * x.y + x.z * x.z
	return math.Sqrt(sum)
}

func (x *vec3) unit() *vec3 {
	reciprocalNorm := 1 / x.norm()
	return x.mult(reciprocalNorm)
}

type ray struct {
	origin vec3
	direction vec3
}
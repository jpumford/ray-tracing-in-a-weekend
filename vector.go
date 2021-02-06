package main

import (
	"math"
	"math/rand"
)

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
	return math.Sqrt(x.normSquared())
}

func (x *vec3) unit() *vec3 {
	reciprocalNorm := 1 / x.norm()
	return x.mult(reciprocalNorm)
}

func (x *vec3) normSquared() float64 {
	return x.x * x.x + x.y * x.y + x.z * x.z
}

func (x *vec3) div(y float64) *vec3 {
	return &vec3{
		x.x / y,
		x.y / y,
		x.z / y,
	}
}

func randomVec(r *rand.Rand) *vec3 {
	return &vec3{
		r.Float64(),
		r.Float64(),
		r.Float64(),
	}
}

func randomInUnitSphere(r *rand.Rand) *vec3 {
	for {
		p := randomVec(r)
		// draw p in unit box
		p = p.mult(2).sub(&vec3{1, 1, 1})

		if p.normSquared() < 1 {
			return p
		}
	}
}

type ray struct {
	origin vec3
	direction vec3
}

func (r *ray) at(t float64) *vec3 {
	return r.origin.add(r.direction.mult(t))
}
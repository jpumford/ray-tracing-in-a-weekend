package main

import (
	"fmt"
	"math"
)

var sphereCenter = &vec3{0, 0, -1}
var sphereRadius = 0.5

func main() {
	nx := 200
	ny := 100
	fmt.Println("P3")
	fmt.Printf("%d %d\n", nx, ny)
	fmt.Println("255")
	origin := vec3{0,0,0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := &ray{
				origin:    origin,
				direction: vec3{u * 4 - 2, v * 2 - 1, -1},
			}
			col := getColor(r)

			ir := int32(255.99 * col.x)
			ig := int32(255.99 * col.y)
			ib := int32(255.99 * col.z)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}

func getColor(r *ray) *vec3 {
	hitT := hitSphere(sphereCenter, sphereRadius, r)
	if hitT != -1 {
		hitPoint := r.origin.add(r.direction.mult(hitT))
		normal := hitPoint.sub(sphereCenter).unit()
		return (&vec3{normal.x + 1, normal.y + 1, normal.z + 1}).mult(0.5)
	}
	unitVec := r.direction.unit()
	t := 0.5 * (unitVec.y + 1.0)
	return (&vec3{1, 1, 1}).mult(1.0 - t).add((&vec3{0.5, 0.7, 1.0}).mult(t))
}

func hitSphere(center *vec3, radius float64, r *ray) float64 {
	a := r.direction.dot(&r.direction)
	b := r.direction.dot(r.origin.sub(center)) * 2
	c := r.origin.sub(center).dot(r.origin.sub(center)) - radius * radius
	discriminant := b * b - 4 * a * c
	if discriminant < 0 {
		return -1
	}
	return (-b - math.Sqrt(discriminant)) / (2 * a)
}
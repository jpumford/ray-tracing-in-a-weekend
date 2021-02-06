package main

import (
	"fmt"
	"math"
)

var widthUnits = 4.0
var heightUnits = 2.0

func main() {
	nx := 200
	ny := 100
	fmt.Println("P3")
	fmt.Printf("%d %d\n", nx, ny)
	fmt.Println("255")
	origin := vec3{0,0,0}

	sphere1 := &sphere{&vec3{0, 0, -1}, 0.5}
	sphere2 := &sphere{&vec3{0, -100.5, -1}, 100}
	world := &hitableList{[]hitable{sphere1, sphere2}}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := &ray{
				origin:    origin,
				direction: vec3{u * widthUnits - widthUnits / 2, v * heightUnits - heightUnits / 2, -1},
			}
			col := getColor(r, world)

			ir := int32(255.99 * col.x)
			ig := int32(255.99 * col.y)
			ib := int32(255.99 * col.z)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}

func getColor(r *ray, world hitable) *vec3 {
	if wasHit, hitRecord := world.hit(r, 0, math.MaxFloat64); wasHit {
		return (&vec3{hitRecord.normal.x + 1, hitRecord.normal.y + 1, hitRecord.normal.z + 1}).mult(0.5)
	}
	unitVec := r.direction.unit()
	t := 0.5 * (unitVec.y + 1.0)
	return (&vec3{1, 1, 1}).mult(1.0 - t).add((&vec3{0.5, 0.7, 1.0}).mult(t))
}


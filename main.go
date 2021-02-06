package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var widthUnits = 4.0
var heightUnits = 2.0

// for shadow acne
var epsilon = 0.001

func main() {
	nx := 800
	ny := 400
	numSamples := 100
	randSource := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(randSource)

	fmt.Println("P3")
	fmt.Printf("%d %d\n", nx, ny)
	fmt.Println("255")

	mat1 := &lambertian{&vec3{0.8, 0.3, 0.3}}
	mat2 := &lambertian{&vec3{0.8, 0.8, 0}}
	mat3 := &metal{&vec3{0.8, 0.6, 0.2}}
	mat4 := &metal{&vec3{0.8, 0.8, 0.8}}
	sphere1 := &sphere{&vec3{0, 0, -1}, 0.5, mat1}
	sphere2 := &sphere{&vec3{0, -100.5, -1}, 100, mat2}
	sphere3 := &sphere{&vec3{1, 0, -1}, 0.5, mat3}
	sphere4 := &sphere{&vec3{-1, 0, -1}, 0.5, mat4}
	world := &hitableList{[]hitable{sphere1, sphere2, sphere3, sphere4}}

	cam := &camera{
		&vec3{0, 0, 0},
		&vec3{0, heightUnits, 0},
		&vec3{widthUnits, 0, 0},
		&vec3{-widthUnits / 2.0, -heightUnits / 2.0, -1.0},
	}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := &vec3{0, 0, 0}
			for sample := 0; sample < numSamples; sample++ {
				jitteredI := float64(i) + randNum.Float64()
				jitteredJ := float64(j) + randNum.Float64()
				u := jitteredI / float64(nx)
				v := jitteredJ / float64(ny)
				r := cam.castRay(u, v)
				color = color.add(getColor(r, world, 0, randNum))
			}

			// average out samples
			color = color.div(float64(numSamples))

			// correct for gamma 2
			color = &vec3{math.Sqrt(color.x), math.Sqrt(color.y), math.Sqrt(color.z)}

			ir := int32(255.99 * color.x)
			ig := int32(255.99 * color.y)
			ib := int32(255.99 * color.z)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}

func getColor(r *ray, world hitable, depth int, randGen *rand.Rand) *vec3 {
	if wasHit, hitRecord := world.hit(r, epsilon, math.MaxFloat64); wasHit {
		if depth < 50 {
			if wasScattered, attenuation, newRay := hitRecord.mat.scatter(r, hitRecord, randGen); wasScattered {
				return attenuation.componentMult(getColor(newRay, world, depth+1, randGen))
			} else {
				return &vec3{0, 0, 0}
			}
		} else {
			return &vec3{0, 0, 0}
		}
	}

	unitVec := r.direction.unit()
	t := 0.5 * (unitVec.y + 1.0)
	return (&vec3{1, 1, 1}).mult(1.0 - t).add((&vec3{0.5, 0.7, 1.0}).mult(t))
}


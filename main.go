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

	sphere1 := &sphere{&vec3{0, 0, -1}, 0.5}
	sphere2 := &sphere{&vec3{0, -100.5, -1}, 100}
	world := &hitableList{[]hitable{sphere1, sphere2}}

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
				color = color.add(getColor(r, world, randNum))
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

func getColor(r *ray, world hitable, randGen *rand.Rand) *vec3 {
	if wasHit, hitRecord := world.hit(r, epsilon, math.MaxFloat64); wasHit {
		newRayStart := hitRecord.p
		centerOfUnitSphereTangentToSurface := hitRecord.p.add(hitRecord.normal)
		newRayDirection := centerOfUnitSphereTangentToSurface.add(randomInUnitSphere(randGen))
		newRayStrength := 0.5
		return getColor(&ray{*newRayStart, *(newRayDirection.sub(newRayStart))}, world, randGen).mult(newRayStrength)
	}
	unitVec := r.direction.unit()
	t := 0.5 * (unitVec.y + 1.0)
	return (&vec3{1, 1, 1}).mult(1.0 - t).add((&vec3{0.5, 0.7, 1.0}).mult(t))
}


package main

import "math/rand"

type material interface {
	// todo: remove hitRecord arg
	scatter(in *ray, hit *hitRecord, randGen *rand.Rand) (bool, *vec3, *ray)
}

type lambertian struct {
	albedo *vec3
}

func (m *lambertian) scatter(in *ray, hit *hitRecord, randGen *rand.Rand) (bool, *vec3, *ray) {
	newRayStart := hit.p
	centerOfUnitSphereTangentToSurface := newRayStart.add(hit.normal)
	newRayDirection := centerOfUnitSphereTangentToSurface.add(randomInUnitSphere(randGen))
	newRay := &ray{*newRayStart, *(newRayDirection.sub(newRayStart))}
	return true, m.albedo, newRay
}

type metal struct {
	albedo *vec3
}

func (m *metal) scatter(in *ray, hit *hitRecord, randGen *rand.Rand) (bool, *vec3, *ray) {
	// todo: could probably skip reflecting and check orthogonality first
	reflected := in.direction.unit().reflect(hit.normal)
	scattered := &ray{*hit.p, *reflected}
	// can't reflect sideways, make sure direction and normal are not orthogonal
	wasScattered := scattered.direction.dot(hit.normal) > 0
	return wasScattered, m.albedo, scattered
}
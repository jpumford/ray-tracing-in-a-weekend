package main

import "math"

type hitRecord struct {
	t float64
	p *vec3
	normal *vec3
}

type hitable interface {
	hit(r *ray, tmin float64, tmax float64) (bool, *hitRecord)
}

type sphere struct {
	center *vec3
	radius float64
}

func (s *sphere) hit(r *ray, tmin float64, tmax float64) (bool, *hitRecord) {
	a := r.direction.dot(&r.direction)
	b := r.direction.dot(r.origin.sub(s.center)) * 2
	c := r.origin.sub(s.center).dot(r.origin.sub(s.center)) - s.radius * s.radius
	discriminant := b * b - 4 * a * c
	if discriminant > 0 {
		// choose smallest t first
		possibleT := (-b - math.Sqrt(discriminant)) / (2 * a)
		if possibleT > tmin && possibleT < tmax {
			hitPoint := r.at(possibleT)
			record := &hitRecord{
				possibleT,
				hitPoint,
				hitPoint.sub(s.center).unit(),
			}
			return true, record
		}
		possibleT = (-b + math.Sqrt(discriminant)) / (2 * a)
		if possibleT > tmin && possibleT < tmax {
			hitPoint := r.at(possibleT)
			record := &hitRecord{
				possibleT,
				hitPoint,
				hitPoint.sub(s.center).unit(),
			}
			return true, record
		}
	}
	return false, nil
}

type hitableList struct {
	hitables []hitable
}

func (l *hitableList) hit(r *ray, tmin float64, tmax float64) (bool, *hitRecord) {
	hitAnything := false
	closestSoFar := tmax
	var toReturn *hitRecord = nil
	for _, item := range l.hitables {
		wasHit, hitRecord := item.hit(r, tmin, closestSoFar)
		if wasHit {
			hitAnything = true
			closestSoFar = hitRecord.t
			toReturn = hitRecord
		}
	}
	return hitAnything, toReturn
}
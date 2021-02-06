package main

type camera struct {
	origin *vec3
	height *vec3
	width *vec3
	lowerLeftCorner *vec3
}

func (c *camera) castRay(u float64, v float64) *ray {
	return &ray{
		*c.origin,
		*c.lowerLeftCorner.add(c.width.mult(u)).add(c.height.mult(v)).sub(c.origin),
	}
}

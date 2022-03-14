package main

import (
	"github.com/jakecoffman/cp"
	"math"
	"math/rand"
)

func uniform(a, b float64) float64 {
	return a + (b-a)*rand.Float64()
}

func irange(a, b int) int {
	return a + rand.Intn(b-a)
}

type Dot struct {
	body  *cp.Body
	moves []cp.Vector
}

func NewRandomDot() *Dot {
	dot := &Dot{moves: make([]cp.Vector, 0)}
	for i := 0; i < irange(10, 20); i++ {
		vector := cp.
			ForAngle(2 * math.Pi * rand.Float64()).Mult(uniform(2, 100))
		dot.moves = append(dot.moves, vector)
	}
	return dot
}

func (d *Dot) CreatePhysicsBody(space *cp.Space) {
	// haha

	d.body = space.AddBody(cp.NewBody(0, 0))
	shape := cp.NewCircle(d.body, 1, cp.Vector{})
	shape.SetDensity(10)
	d.body.AddShape(space.AddShape(shape))
	d.body.UserData = d
}

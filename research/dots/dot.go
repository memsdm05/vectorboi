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
	body    *cp.Body
	moves   []cp.Vector
	dead    bool
	scored  bool
	fitness float64
}

func NewRandomDot() *Dot {
	dot := &Dot{moves: make([]cp.Vector, 0)}
	for i := 0; i < 20; i++ {
		vector := cp.
			ForAngle(2 * math.Pi * rand.Float64()).Mult(uniform(50, 100))
		dot.moves = append(dot.moves, vector)
	}
	return dot
}

func (d *Dot) CreatePhysicsBody(space *cp.Space) {
	// haha

	d.body = cp.NewBody(0, 0)
	shape := cp.NewCircle(d.body, 1, cp.Vector{})
	shape.SetMass(1)
	shape.SetFilter(cp.SHAPE_FILTER_NONE)
	d.body.AccumulateMassFromShapes()

	space.AddShape(shape)
	space.AddBody(d.body)
	d.body.UserData = d
}

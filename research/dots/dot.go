package main

import (
	"fmt"
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
	Moves   []cp.Vector
	history []cp.Vector
	dead    bool
	scored  bool
	fitness float64
	Age     int
}

func NewRandomDot() *Dot {
	dot := &Dot{Moves: make([]cp.Vector, 0)}
	for i := 0; i < irange(5, 15); i++ {
		vector := cp.
			ForAngle(2 * math.Pi * rand.Float64()).Mult(uniform(50, 100))
		dot.Moves = append(dot.Moves, vector)
	}
	return dot
}

func (d *Dot) Kick(move int) {

}

func (d *Dot) CreatePhysicsBody(space *cp.Space) {
	// haha

	d.body = cp.NewBody(0, 0)
	shape := cp.NewCircle(d.body, 1, cp.Vector{})
	shape.SetCollisionType(1)
	shape.SetMass(1)
	shape.SetFilter(cp.ShapeFilter{
		Group: 1,
		Categories: cp.ALL_CATEGORIES,
		Mask:       cp.ALL_CATEGORIES,
	})
	d.body.AccumulateMassFromShapes()

	space.AddShape(shape)
	space.AddBody(d.body)
	d.body.UserData = d
}

func (d *Dot) String() string {
	return fmt.Sprintf("Dot %.2f", d.fitness)
}

func (d *Dot) SetScored()  {
	d.scored = true
	d.body.SetType(cp.BODY_STATIC)
}
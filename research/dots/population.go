package main

import (
	"github.com/jakecoffman/cp"
)

type Eval func(dot *Dot, population *Population) float64

func ConstantFitness(dot *Dot, pop *Population) float64 {
	return 1
}

type Population struct {
	Dots   []*Dot
	Num    int
	Space  *cp.Space
	Time   float64
	OnMove int

	spawn   cp.Vector
	fitness Eval
	running bool

	bestDot        *Dot
	bestDotFitness float64
}

func NewRandomPopulation(num int, spawn cp.Vector, fitness Eval) *Population {
	if fitness == nil {
		fitness = ConstantFitness
	}

	pop := &Population{
		Dots:    make([]*Dot, num, num),
		Num:     num,
		Space:   cp.NewSpace(),
		fitness: fitness,
		spawn:   spawn,
	}
	pop.Space.UseSpatialHash(2, 100)

	for i := 0; i < num; i++ {
		ndot := NewRandomDot()
		ndot.CreatePhysicsBody(pop.Space)
		ndot.body.SetPosition(spawn)
		pop.Dots[i] = ndot
	}

	return pop
}

func (p *Population) IsBest(dot *Dot) bool {
	return dot == p.bestDot
}

func (p *Population) Step(dt float64) {
	p.Space.Step(dt)
	p.Time += dt

	for _, dot := range p.Dots {
		//if p.OnMove < len(dot.moves) {
		//	dot.body.ApplyImpulseAtLocalPoint(dot.moves[p.OnMove], cp.Vector{})
		//}

		if f := p.fitness(dot, p); f > p.bestDotFitness {
			p.bestDot = dot
			p.bestDotFitness = f
		}
	}

	//p.OnMove++
}

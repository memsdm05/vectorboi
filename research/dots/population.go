package main

import (
	"github.com/jakecoffman/cp"
)

type Eval func(dot *Dot, population *Population) float64

func ConstantFitness(dot *Dot, pop *Population) float64 {
	return 1
}

type Population struct {
	Dots  []*Dot
	Num   int
	Space *cp.Space
	Time  float64

	onmove  int
	fitness Eval
	running bool

	bestDot        *Dot
	bestDotFitness float64
}

func NewRandomPopulation(num int, fitness Eval) *Population {
	if fitness == nil {
		fitness = ConstantFitness
	}

	pop := &Population{
		Dots:    make([]*Dot, num, num),
		Num:     num,
		Space:   cp.NewSpace(),
		fitness: fitness,
	}

	for i := 0; i < num; i++ {
		pop.Dots[i] = NewRandomDot()
	}

	return pop
}

func (p *Population) Step(dt float64) {
	p.Space.Step(dt)
	p.Time += dt

	for _, dot := range p.Dots {
		if p.onmove < len(dot.moves) {
			dot.body.ApplyImpulseAtLocalPoint(dot.moves[p.onmove], cp.Vector{})
		}

		if f := p.fitness(dot, p); f > p.bestDotFitness {
			p.bestDot = dot
			p.bestDotFitness = f
		}
	}

	p.onmove++
}

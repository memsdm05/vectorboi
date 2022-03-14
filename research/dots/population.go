package main

import (
	"github.com/jakecoffman/cp"
)

type Eval func(dot *Dot, population *Population) float64

func ConstantFitness(dot *Dot, pop *Population) float64 {
	return 1
}

type Population struct {
	Dots          []*Dot
	Num           int
	Space         *cp.Space
	OnMove        int
	Width, Height int

	cycletime      float64
	spawn          cp.Vector
	fitness        Eval
	running        bool
	bestDot        *Dot
	bestDotFitness float64
}

func NewRandomPopulation(num, width, height int, fitness Eval) *Population {
	if fitness == nil {
		fitness = ConstantFitness
	}

	pop := &Population{
		Dots:   make([]*Dot, num, num),
		Num:    num,
		Space:  cp.NewSpace(),
		Width:  width,
		Height: height,
		spawn: cp.Vector{
			X: Width / 2,
			Y: Height - Height/10,
		},
		fitness: fitness,
	}

	pop.Space.UseSpatialHash(2, 100)
	for i := 0; i < num; i++ {
		ndot := NewRandomDot()
		ndot.CreatePhysicsBody(pop.Space)
		ndot.body.SetPosition(pop.spawn)
		pop.Dots[i] = ndot
	}

	return pop
}

func (p *Population) kill(dot *Dot) {
	dot.dead = true
	p.Space.Deactivate(dot.body)
}

func (p *Population) IsBest(dot *Dot) bool {
	return dot == p.bestDot
}

func (p *Population) Step(dt float64) {
	p.Space.Step(dt)

	hitnow := false
	p.cycletime += dt
	if p.cycletime >= 1.2 {
		p.cycletime = 0
		hitnow = true
	}

	p.bestDot = nil
	p.bestDotFitness = 0

	for _, dot := range p.Dots {
		//if p.OnMove < len(dot.moves) {
		//	dot.body.ApplyImpulseAtLocalPoint(dot.moves[p.OnMove], cp.Vector{})
		//}
		if hitnow && p.OnMove < len(dot.moves) {
			dot.body.ApplyImpulseAtLocalPoint(dot.moves[p.OnMove], cp.Vector{})
			//dot.body.ApplyForceAtLocalPoint(, cp.Vector{})
		}

		pos := dot.body.Position()
		if pos.X > float64(p.Width)-2 || pos.X < 2 || pos.Y > float64(p.Height)-2 || pos.Y < 2 {
			p.kill(dot)
		}

		if f := p.fitness(dot, p); f > p.bestDotFitness {
			p.bestDot = dot
			p.bestDotFitness = f
		}
	}

	if hitnow {
		p.OnMove++
	}

	//p.OnMove++
}

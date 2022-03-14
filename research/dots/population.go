package main

import "github.com/jakecoffman/cp"

type Eval func(dot *Dot, population *Population) float64

func ConstantFitness(dot *Dot, pop *Population) float64 {
	return 1
}

type Population struct {
	Dots []*Dot
	Num int
	Space *cp.Space

	fitness Eval
}

func NewRandomPopulation(num int, fitness Eval) *Population {
	if fitness == nil {
		fitness = ConstantFitness
	}

	pop := &Population{
		Dots: make([]*Dot, num, num),
		Num:  num,
		Space: cp.NewSpace(),
		fitness: fitness,
	}

	for i := 0; i < num; i++ {
		pop.Dots[i] = NewRandomDot()
	}

	return pop
}
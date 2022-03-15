package main

import (
	"github.com/jakecoffman/cp"
	"math/rand"
)

func Clone(dot *Dot) *Dot {
	ret := &Dot{
		moves:  make([]cp.Vector, len(dot.moves), len(dot.moves)),
	}
	for i := range dot.moves {
		ret.moves[i] = dot.moves[i]
	}
	return ret
}

func Mutate(dot *Dot) *Dot{
	for i := range dot.moves {
		if rand.Float64() < .3 {
			dot.moves[i].X += rand.NormFloat64() * 3
			dot.moves[i].Y += rand.NormFloat64() * 3
		}
	}
	return dot
}

func Crossover(a, b *Dot) (*Dot, *Dot) {
	return nil, nil
}
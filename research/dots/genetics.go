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
			move := &dot.moves[i]
			move.X += rand.Float64() * 5
			move.Y += rand.Float64() * 5
			*move = move.Clamp(100)
		}
	}
	return dot
}

func Crossover(a, b *Dot) (*Dot, *Dot) {
	childA := Clone(a)
	childB := Clone(b)

	// A should be larger than B
	if len(childA.moves) < len(childB.moves) {
		childA, childB = childB, childA
	}

	alen, blen := len(childA.moves), len(childB.moves)
	n := rand.Intn(blen)

	cull := 0
	for i := n; i < alen; i++ {
		if i < blen {
			childA.moves[i], childB.moves[i] = childB.moves[i], childA.moves[i]
		} else {
			childB.moves = append(childB.moves, childA.moves[i])
			cull++
		}
	}
	childA.moves = childA.moves[:alen - cull]


	return childA, childB
}
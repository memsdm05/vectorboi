package main

import (
	"github.com/jakecoffman/cp"
	"math"
	"math/rand"
)

func Clone(dot *Dot) *Dot {
	ret := &Dot{
		Moves: make([]cp.Vector, len(dot.Moves), len(dot.Moves)),
	}
	for i := range dot.Moves {
		ret.Moves[i] = dot.Moves[i]
	}
	return ret
}

func Mutate(dot *Dot) *Dot{
	for i := len(dot.Moves) - 1; i >= 0; i-- {
		// 10% chance of changing the vector
		if rand.Float64() < .1 {
			move := &dot.Moves[i]
			move.X += uniform(-5, 5)
			move.Y += uniform(-5, 5)
			*move = move.Clamp(100)
		}

		// 5% change of deleting the move (if there is more than 1 move left)
		if rand.Float64() < .05 && len(dot.Moves) > 1 {
			dot.Moves = append(dot.Moves[:i], dot.Moves[i+1:]...)
		}

		// 5% change of adding a move
		if rand.Float64() < .05 {
			dot.Moves = append(dot.Moves[:i+1], dot.Moves[i:]...)
			dot.Moves[i] = cp.
				ForAngle(2 * math.Pi * rand.Float64()).Mult(uniform(50, 100))
		}

	}
	return dot
}

func Crossover(a, b *Dot) (*Dot, *Dot) {
	childA := Clone(a)
	childB := Clone(b)

	// A should be larger than B
	if len(childA.Moves) < len(childB.Moves) {
		childA, childB = childB, childA
	}

	alen, blen := len(childA.Moves), len(childB.Moves)
	n := rand.Intn(blen)

	cull := 0
	for i := n; i < alen; i++ {
		if i < blen {
			childA.Moves[i], childB.Moves[i] = childB.Moves[i], childA.Moves[i]
		} else {
			childB.Moves = append(childB.Moves, childA.Moves[i])
			cull++
		}
	}
	childA.Moves = childA.Moves[:alen - cull]


	return childA, childB
}
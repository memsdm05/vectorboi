package main

import "github.com/jinzhu/copier"

func Clone(dot *Dot) *Dot {
	ret := &Dot{}
	copier.Copy(ret, dot)
	return ret
}

func Mutate(dot *Dot) *Dot{
	for _, move := range dot.moves {
		move.X *= uniform(-2, 2)
		move.Y *= uniform(-2, 2)
	}
	return dot
}

func Crossover(a, b *Dot) (*Dot, *Dot) {
	return nil, nil
}
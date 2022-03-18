package genetics

import (
	"github.com/jakecoffman/cp"
	"vectorboi/app/dot"
	"vectorboi/app/utils"
)

type Genetics struct {
	rand *utils.ExRand
	chances *dot.Chances
}

func New(p *dot.Population) *Genetics {
	return &Genetics{
		rand:    p.Rand,
		chances: &p.Scenario.Chances,
	}
}

func (g *Genetics) Clone(d *dot.Dot) *dot.Dot {
	ret := &dot.Dot{
		Kicks: make([]cp.Vector, len(d.Kicks)),
	}
	copy(ret.Kicks, d.Kicks)
	return ret
}

func (g *Genetics) Mutate(d *dot.Dot) *dot.Dot {
	for i := len(d.Kicks) - 1; i >= 0; i-- {
		// change a kick
		if g.rand.Roll(g.chances.ChangeKick) {
			move := &d.Kicks[i]
			move.X += g.rand.Uniform(-5, 5)
			move.Y += g.rand.Uniform(-5, 5)
			*move = move.Clamp(100)
		}

		// remove a kick (if there is more than 1 kick left)
		if g.rand.Roll(g.chances.RemoveKick) && len(d.Kicks) > 1 {
			d.Kicks = append(d.Kicks[:i], d.Kicks[i+1:]...)
		}

		// add a kick
		if g.rand.Roll(g.chances.AddKick) {
			d.Kicks = append(d.Kicks[:i+1], d.Kicks[i:]...)
			d.Kicks[i] = dot.RandomVector(g.rand)
		}

		// todo implement swap kick
		// swap a kick
		//if g.rand.Roll(g.chances.SwapKick) {
		//	other := g.rand.IntRange(i, len())
		//	d.Kicks[i], d.Kicks[j] =
		//}

	}
	return d
}

func (g *Genetics) Crossover(a, b *dot.Dot) (*dot.Dot, *dot.Dot) {
	childA := g.Clone(a)
	childB :=  g.Clone(b)

	// A should be larger than B
	if len(childA.Kicks) < len(childB.Kicks) {
		childA, childB = childB, childA
	}

	alen, blen := len(childA.Kicks), len(childB.Kicks)
	n := g.rand.Intn(blen)

	cull := 0
	// iterate from crossover point to end of parentA
	for i := n; i < alen; i++ {
		//
		if i < blen {
			childA.Kicks[i], childB.Kicks[i] = childB.Kicks[i], childA.Kicks[i]
		} else {
			childB.Kicks = append(childB.Kicks, childA.Kicks[i])
			cull++
		}
	}
	childA.Kicks = childA.Kicks[:alen-cull]

	return childA, childB
}

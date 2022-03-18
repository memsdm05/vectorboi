package dot

import (
	"github.com/jakecoffman/cp"
)

func (p *Population) clone(d *Dot) *Dot {
	ret := &Dot{
		Kicks: make([]cp.Vector, len(d.Kicks)),
	}
	copy(ret.Kicks, d.Kicks)
	return ret
}

func (p *Population) mutate(d *Dot) *Dot {
	for i := len(d.Kicks) - 1; i >= 0; i-- {
		// change a kick
		if p.exr.Roll(p.Scenario.Chances.ChangeKick) {
			move := &d.Kicks[i]
			move.X += p.exr.Uniform(-5, 5)
			move.Y += p.exr.Uniform(-5, 5)
			*move = move.Clamp(100)
		}

		// remove a kick (if there is more than 1 kick left)
		if p.exr.Roll(p.Scenario.Chances.RemoveKick) && len(d.Kicks) > 1 {
			d.Kicks = append(d.Kicks[:i], d.Kicks[i+1:]...)
		}

		// add a kick
		if p.exr.Roll(p.Scenario.Chances.AddKick) {
			d.Kicks = append(d.Kicks[:i+1], d.Kicks[i:]...)
			d.Kicks[i] = RandomVector(p.exr)
		}

		// todo implement swap kick
		// swap a kick
		//if p.exr.Roll(p.Scenario.Chances.SwapKick) {
		//	var j int
		//	for {
		//		j = p.exr.Intn(len(d.Kicks))
		//		if j == i { continue }
		//		break
		//	}
		//	d.Kicks[i], d.Kicks[j] = d.Kicks[j], d.Kicks[i]
		//}

	}
	return d
}

func (p *Population) crossover(parentA, parentB *Dot) (*Dot, *Dot) {
	childA := p.clone(parentA)
	childB := p.clone(parentB)

	// A should be larger than B
	if len(childA.Kicks) < len(childB.Kicks) {
		childA, childB = childB, childA
	}

	alen, blen := len(childA.Kicks), len(childB.Kicks)
	n := p.exr.Intn(blen)

	cull := 0
	// iterate from crossover point to end of parentA
	for i := n; i < alen; i++ {
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

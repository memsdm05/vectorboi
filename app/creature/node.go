package creature

import (
	"fmt"
	"github.com/jakecoffman/cp"
	"math"
	"math/rand"
	"strings"
)

type Node struct {
	Parent    *Node
	Children  []*Node
	Thrusters []Thruster
	Angle     float64
	Radius    float64

	cachedPos cp.Vector
	cachedDepth int
}

func NewRandomNode() *Node {
	return &Node{
		Children: make([]*Node, 0),
		Angle:    rand.Float64() * 2 * math.Pi,
		Radius:   10 + rand.Float64() * 10,
	}
}

func (n *Node) resetCache()  {
	n.cachedPos = cp.Vector{}
	n.cachedDepth = 0
}

func (n *Node) String() string {
	if n.Orphan() {
		return "A0"
	}

	counter := make([]int, 26, 26)
	n.measureDepth('A'-1, counter)

	var sb strings.Builder
	for i, count := range counter {
		if count > 0 {
			sb.WriteRune(rune(i) + 'A')
			sb.WriteString(fmt.Sprintf("%d", counter[i]))
		}
	}

	return sb.String()
}

func (n *Node) measureDepth(depth rune, counter []int) {
	if depth > 'Z' {
		return
	}

	if !n.Root() && n.Leaf() {
		counter[depth-'A']++
	}

	for _, child := range n.Children {
		child.measureDepth(depth+1, counter)
	}
}

func (n *Node) Do(f func(n *Node)) {
	f(n)
	for _, child := range n.Children {
		child.Do(f)
	}
}

func (n *Node) Root() bool {
	return n.Parent == nil
}

func (n *Node) Leaf() bool {
	return len(n.Children) == 0
}

func (n *Node) Orphan() bool {
	return n.Root() && n.Leaf()
}

func (n *Node) Intersects(other *Node) bool {
	if n.Parent == other {
		return false
	}
	return n.Position().Near(other.Position(), n.Radius * 2 + other.Radius)
}

func (n *Node) Attach(child *Node) {
	n.Children = append(n.Children, child)
	child.Parent = n
}

func (n *Node) Detach()  {
	pc := n.Parent.Children
	for i := len(pc); i >= 0; i-- {
		if pc[i] == n {
			last := pc[len(pc) - 1]
			pc[i], last = last, nil
			break
		}
	}
	n.Parent = nil
}

func (n *Node) Position() cp.Vector {
	if n.Root() {
		return cp.Vector{}
	}

	if !n.cachedPos.Equal(cp.Vector{}) {
		return n.cachedPos
	}

	p := n.Parent
	n.cachedPos = p.Position().Add(cp.ForAngle(n.Angle).Mult(p.Radius + n.Radius))
	return n.cachedPos
}

type Thruster struct {
	Direction float64
	MaxThrust float64
}

func (t *Thruster) Power(throttle float64) cp.Vector {
	return cp.ForAngle(t.Direction).Mult(throttle * t.MaxThrust)
}
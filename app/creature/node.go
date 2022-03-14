package creature

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"math"
	"math/rand"
	"strings"
	"vectorboi/helpers"
)

type Node struct {
	Parent    *Node
	Children  []*Node
	Thrusters []Thruster
	Angle     float64
	Radius    float64
}

func (n *Node) ScaleMe() bool {
	return false
}

func (n *Node) Draw(zoom float64) (*ebiten.Image, *ebiten.DrawImageOptions) {
	if n.Root() {
		return helpers.CircleImage(n.Radius * zoom, colornames.Red), nil
	}
	return helpers.CircleImage(n.Radius * zoom, colornames.White), nil
}

func NewNode() *Node {
	return &Node{
		Children: make([]*Node, 0),
	}
}

func NewRandomNode() *Node {
	node := NewNode()
	node.Angle = rand.Float64() * 2 * math.Pi
	node.Radius = 10 + rand.Float64() * 10
	return node
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


	p := n.Parent
	return p.Position().Add(cp.ForAngle(n.Angle).Mult(p.Radius + n.Radius))
}

type Thruster struct {
	Direction float64
	MaxThrust float64
}

func (t *Thruster) Power(throttle float64) cp.Vector {
	return cp.ForAngle(t.Direction).Mult(throttle * t.MaxThrust)
}
package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
	"strings"
	"time"
	"vectorboi/helpers"
)

var (
	mainCreature *Creature
	NumNodes     = 3
)

type Creature struct {
	Body *Node
}

func NewRandomCreature(num int) *Creature {
	ret := &Creature{Body: NewRandomNode()}
	all := []*Node{ret.Body}

	for i := 0; i < num; i++ {
	adder:
		for {
			node := NewRandomNode()
			node.Parent = all[rand.Intn(len(all))]

			for _, other := range all {
				if node.Intersects(other) {
					continue adder
				}
			}

			node.Parent.Attach(node)
			all = append(all, node)
			break
		}
	}

	return ret
}

type Node struct {
	Parent    *Node
	Children  []*Node
	Thrusters []Thruster
	Angle     float64
	Radius    float64
}

func NewRandomNode() *Node {
	return &Node{
		Children: make([]*Node, 0),
		Angle:    rand.Float64() * 2 * math.Pi,
		Radius:   10 + rand.Float64()*10,
	}
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
	return n.Position().Near(other.Position(), n.Radius*2+other.Radius)
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

type CreatureViewGame struct{}

func reload() {
	rand.Seed(time.Now().UnixNano())
	mainCreature = NewRandomCreature(NumNodes)
}

func (c CreatureViewGame) Init() {
	reload()
}

func (c CreatureViewGame) Shutdown() {}

func (c CreatureViewGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		reload()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		NumNodes++
		reload()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		NumNodes--
		reload()
	}

	return nil
}

func (c CreatureViewGame) Draw(screen *ebiten.Image) {
	w, h := ebiten.WindowSize()
	offset := cp.Vector{
		X: float64(w) / 2,
		Y: float64(h) / 2,
	}

	mainCreature.Body.Do(func(node *Node) {
		var c color.Color
		if node.Root() {
			c = colornames.Red
		} else {
			c = colornames.Whitesmoke
		}
		helpers.DrawCircle(screen, node.Position().Add(offset), node.Radius+1.2, c)
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("num: %v, %s", NumNodes, mainCreature.Body))
}

func (c CreatureViewGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowResizable(true)
	helpers.RunGame(new(CreatureViewGame))
}

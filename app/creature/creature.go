package creature

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Creature struct {
	Body []*Node
}

func NewRandomCreature(num int) *Creature {
	creature := &Creature{
		Body: []*Node{
			NewRandomNode(),
		},
	}

	for i := 0; i < num; i++ {
	check:
		for {
			node := NewRandomNode()
			node.Parent = creature.randomBodyNode()

			for _, other := range creature.Body {
				if node.Intersects(other) {
					continue check
				}
			}

			node.Parent.Attach(node)
			creature.Body = append(creature.Body, node)
			break
		}
	}

	return creature
}

func (c *Creature) root() *Node {
	return c.Body[0]
}

func (c *Creature) randomBodyNode() *Node {
	return c.Body[rand.Intn(len(c.Body))]
}

func (c *Creature) Draw(dst *ebiten.Image)  {
}

//func CreatePhysicsBody(space cp.Space) *cp.Body {
//	cbody := cp.NewBody()
//}
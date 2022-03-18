package structures

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

type KillWall struct {
	A cp.Vector
	B cp.Vector
	added bool
}

func MakeKillWall(ax, ay, bx, by float64) KillWall {
	return KillWall{
		A: cp.Vector{ax, ay},
		B: cp.Vector{bx, by},
	}
}

func (k KillWall) Draw(dst *ebiten.Image) {
	ebitenutil.DrawLine(dst,
		k.A.X, k.A.Y, k.B.X, k.B.Y, colornames.Orange)
}

func (k *KillWall) PhysicsShape(space *cp.Space) *cp.Shape {
	if k.added {
		return nil
	}

	shape := cp.NewSegment(space.StaticBody, k.B, k.A, 3)
	shape.SetCollisionType(2)
	shape.SetSensor(true)
	shape.SetFilter(cp.ShapeFilter{
		Group:      2,
		Categories: cp.ALL_CATEGORIES,
		Mask:       cp.ALL_CATEGORIES,
	})
	space.AddShape(shape)
	k.added = true
	return shape
}

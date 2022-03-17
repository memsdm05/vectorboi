package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

type KillWall struct {
	A cp.Vector
	B cp.Vector
}

func MakeKillWall(ax, ay, by, bx float64) KillWall {
	return KillWall{
		A: cp.Vector{ax, ay},
		B: cp.Vector{bx, by},
	}
}

func (k KillWall) Draw(dst *ebiten.Image)  {
	ebitenutil.DrawLine(dst,
		k.A.X, k.A.Y, k.B.X, k.B.Y, colornames.Orange)
}

func (k KillWall) PhysicsShape(space *cp.Space) *cp.Shape {
	shape := space.AddShape(cp.NewSegment(space.StaticBody, k.B, k.A, 3))
	shape.SetCollisionType(2)
	shape.SetSensor(false)
	shape.SetFilter(cp.ShapeFilter{
		Group: 2,
		Categories: cp.ALL_CATEGORIES,
		Mask:       cp.ALL_CATEGORIES,
	})
	shape.UserData = k
	return shape
}

// lol
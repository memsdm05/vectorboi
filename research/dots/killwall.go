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
	shape := space.AddShape(cp.NewSegment(space.StaticBody, k.A, k.B, 5))
	shape.SetSensor(true)
	return shape
}

// lol
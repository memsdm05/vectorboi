package app

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"image/color"
	_ "image/png"
	"vectorboi/app/dot"
	"vectorboi/app/structures"
	"vectorboi/app/utils"
)

//go:embed spawn_ball.png
var spawnBallData []byte

var spawnBall, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(spawnBallData))
var sbw, sbh = spawnBall.Size()

type Editor struct {
	work *dot.Scenario
	buildingWall bool
	wallStart cp.Vector
	sbOP *ebiten.DrawImageOptions
}

func NewEditor(scenario *dot.Scenario) *Editor {
	return &Editor{
		work:         scenario,
		buildingWall: false,
		sbOP:  &ebiten.DrawImageOptions{},
	}
}


// AHHHHHHHHHHHHHHHHHHH
func (e *Editor) Interact() {
	x, y := ebiten.CursorPosition()
	mpos := cp.Vector{float64(x), float64(y)}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyZ) && ebiten.IsKeyPressed(ebiten.KeyControl):
		e.work.Walls = e.work.Walls[:len(e.work.Walls) - 1]
	case inpututil.IsKeyJustPressed(ebiten.KeyAlt):
		e.wallStart = cp.Vector{}
	case inpututil.IsKeyJustPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyControl):
		fmt.Printf("Set Name (currently \"%v\"): ", e.work.Name)
		fmt.Scanln(&e.work.Name)
		utils.Export("scenario", e.work)
	}


	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		switch {
		case mpos.Near(e.work.Spawn, 15):
			e.work.Spawn = mpos
		case e.work.Target.ContainsVect(mpos):
			center := e.work.Target.Center()
			e.work.Target = e.work.Target.Offset(mpos.Sub(center))
		default:
			switch {
			case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft):
				if e.wallStart.Equal(cp.Vector{}) {
					e.wallStart = mpos
				}

				e.work.Walls = append(e.work.Walls, structures.KillWall{
					A: e.wallStart,
					B: mpos,
				})

				e.wallStart = mpos
			}
		}
	}
}

const size = 25

func (e *Editor) Draw(dst *ebiten.Image) {
	dst.Fill(color.RGBA{0, 20, 132, 255})

	for x := float64(0); x < float64(e.work.Width); x += size {
		ebitenutil.DrawLine(dst, x, 0, x, float64(e.work.Height), colornames.Darkgray)
	}

	for y := float64(0); y < float64(e.work.Width); y += size {
		ebitenutil.DrawLine(dst, 0, y, float64(e.work.Width), y, colornames.Darkgray)
	}

	t := e.work.Target
	ebitenutil.DrawRect(dst,
		t.L, t.B, t.R-t.L, t.T-t.B, colornames.Silver)

	e.sbOP.GeoM.Reset()
	e.sbOP.GeoM.Translate(float64(-sbw) / 2, float64(-sbh) / 2)
	e.sbOP.GeoM.Scale(0.2,0.2)
	e.sbOP.GeoM.Translate(e.work.Spawn.X, e.work.Spawn.Y)
	dst.DrawImage(spawnBall, e.sbOP)

	for _, k := range e.work.Walls {
		ebitenutil.DrawLine(dst,
			k.A.X, k.A.Y, k.B.X, k.B.Y, colornames.White)
	}
}
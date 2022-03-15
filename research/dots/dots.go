package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"vectorboi/helpers"
)

const (
	PopulationSize = 2000
	Width          = 640
	Height         = 480
	TimeStep       = 1 / 60.
)

type DotGame struct {
	pop *Population
}

func (d *DotGame) Init() {
	d.pop = NewRandomPopulation(PopulationSize, Width, Height, nil)
	d.pop.Space.SetDamping(0.5)
	//d.pop.Space.SetGravity(cp.Vector{Y: 1000})
}

func (d *DotGame) Shutdown() {}

func (d *DotGame) Update() error {
	d.pop.Step(TimeStep)

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		for _, dot := range d.pop.Dots {
			if d.pop.OnMove < len(dot.moves) {
				dot.body.ApplyImpulseAtLocalPoint(dot.moves[d.pop.OnMove], cp.Vector{})
				//dot.body.ApplyForceAtLocalPoint(, cp.Vector{})
			}
		}
		d.pop.OnMove++
	}

	//if inpututil.IsKeyJustPressed(ebiten.KeyR) {
	//	for _, dot := range d.pop.Dots {
	//		dot.body.SetPosition(cp.Vector{
	//			X: uniform(0, Width),
	//			Y: uniform(0, Height),
	//		})
	//	}
	//}

	return nil
}

func (d *DotGame) Draw(screen *ebiten.Image) {
	d.pop.Draw(screen)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("move %v, gen %v, dt %.2f", d.pop.OnMove, d.pop.Generation, d.pop.Time))
}

func (d *DotGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	helpers.RunGame(new(DotGame))
}

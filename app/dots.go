package app

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"time"
	"vectorboi/app/dot"
	"vectorboi/helpers"
)

const (
	PopulationSize = 1000
	Width          = 640
	Height         = 480
	TimeStep       = 1 / 60.
)

type DotGame struct {
	pop *dot.Population
	debug bool
}

func (d *DotGame) Init() {
	d.pop = dot.NewRandomPopulation(PopulationSize, Width, Height, nil)
	d.pop.Space.SetDamping(0.5)
	d.pop.AddKillWalls(
		app.MakeKillWall(2, 200, 300, 300),
		app.MakeKillWall(Width, 100, 220, 200),
		app.MakeKillWall(200, 200, Width - 200, 200),
	)

	start := time.Now()
	for time.Since(start) < 40 * time.Second {
		d.pop.Step(1./30)
	}
	//d.pop.Space.SetGravity(cp.Vector{Y: 1000})
}

func (d *DotGame) Shutdown() {}

func (d *DotGame) Update() error {
	d.pop.Step(TimeStep)

	if d.debug {
		x, y := ebiten.CursorPosition()
		info := d.pop.Space.PointQueryNearest(cp.Vector{float64(x), float64(y)}, 10, cp.SHAPE_FILTER_ALL)
		if info.Shape != nil {
			fmt.Println("the mouse is over something:", info.Point)
		}
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		d.pop.Paused = !d.pop.Paused
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		d.debug = !d.debug
	}

	//if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	//	for _, dot := range d.pop.Dots {
	//		if d.pop.OnMove < len(dot.Kicks) {
	//			dot.body.ApplyImpulseAtLocalPoint(dot.Kicks[d.pop.OnMove], cp.Vector{})
	//			//dot.body.ApplyForceAtLocalPoint(, cp.Vector{})
	//		}
	//	}
	//	d.pop.OnMove++
	//}

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
	msg := fmt.Sprintf("move %v, gen %v, dt %.2f", d.pop.OnMove, d.pop.Generation, d.pop.Time)

	if d.debug {
		d.pop.Space.EachShape(func(shape *cp.Shape) {
			pos := shape.BB().Center()
			ebitenutil.DrawRect(screen, pos.X-2, pos.Y-2, 4, 4, colornames.Aqua)
		})

		x, y := ebiten.CursorPosition()
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("(%d, %d)", x, y), x, y - 12)

		msg += fmt.Sprintf("\nFPS %.2f, TPS %.2f",
			ebiten.CurrentFPS(), ebiten.CurrentTPS())
	}



	ebitenutil.DebugPrint(screen, msg)
}

func (d *DotGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	helpers.RunGame(new(DotGame))
}

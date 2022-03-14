package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"vectorboi/helpers"
)

const (
	PopulationSize = 500
)

type DotGame struct {
	pop *Population
}

func (d *DotGame) Init() {
	d.pop = NewRandomPopulation(500, nil)
}

func (d DotGame) Shutdown() {}

func (d DotGame) Update() error {
	return nil
}

func (d DotGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.White)
	for _, dot := range d.pop.Dots {
		pos := dot.body.Position()
		screen.Set(int(pos.X), int(pos.Y), colornames.Black)
	}
}

func (d DotGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideWidth
}

func main() {
	helpers.RunGame(new(DotGame))
}

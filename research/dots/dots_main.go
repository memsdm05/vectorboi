package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"vectorboi/helpers"
)



type DotGame struct {}

func (d DotGame) Init() {}

func (d DotGame) Shutdown() {}

func (d DotGame) Update() error {
	return nil
}

func (d DotGame) Draw(screen *ebiten.Image) {

}

func (d DotGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideWidth
}

func main()  {
	helpers.RunGame(new(DotGame))
}
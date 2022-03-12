package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"vectorboi/app"
	"vectorboi/helpers"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

func main() {
	ebiten.SetWindowTitle("vectorboi")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(1920, 1080)

	helpers.RunGame(new(app.Game))
}

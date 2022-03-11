package app

import "github.com/hajimehoshi/ebiten/v2"

type SimpleGame struct {}

func (s SimpleGame) Init() {
	ebiten.SetWindowTitle("vectorboi")
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowSize(1920, 1080)
}

func (s SimpleGame) Shutdown() {}

func (s SimpleGame) Update() error {
	return nil
}

func (s SimpleGame) Draw(screen *ebiten.Image) {}

func (s SimpleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
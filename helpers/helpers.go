package helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type ContextGame interface {
	ebiten.Game
	Init()
	Shutdown()
}

func RunGame(game ContextGame) {
	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	} else {
		game.Shutdown()
	}
}
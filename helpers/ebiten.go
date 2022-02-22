package helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"io/ioutil"
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

func MustNewShader(src []byte) *ebiten.Shader {
	if shader, err := ebiten.NewShader(src); err != nil {
		panic(err)
	} else {
		return shader
	}
}

func MustLoadShader(path string) *ebiten.Shader {
	if src, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		return MustNewShader(src)
	}
}

func Color2Slice(color color.Color) []float32 {
	r, g, b, a := color.RGBA()
	return []float32{float32(r) / 0xffff, float32(g) / 0xffff, float32(b) / 0xffff, float32(a) / 0xffff}
}
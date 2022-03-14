package helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"image/color"
	"io/ioutil"
	"log"
	"math"
)

var CircleShader = MustLoadShader("public/circle.kage")

type ContextGame interface {
	Init()
	Shutdown()
	ebiten.Game
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

func CircleImage(radius float64, color color.Color) *ebiten.Image {
	side := int(math.Ceil(radius) * 2)

	circle := ebiten.NewImage(side, side)
	circle.DrawRectShader(side, side, CircleShader, &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Color": Color2Slice(color),
		},
	})

	return circle
}

// Kill me now
func DrawCircle(dst *ebiten.Image, pos cp.Vector, radius float64, color color.Color) {
	//side := int(math.Ceil(radius) * 2)
	//halfside := float64(side / 2)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.X+radius, pos.Y+radius)
	dst.DrawImage(CircleImage(radius, color), op)
}

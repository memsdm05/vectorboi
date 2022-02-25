package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"io/ioutil"
	"log"
	"vectorboi/helpers"
)

const SideLength = 400

type CircleGame struct {
	shader *ebiten.Shader
}

func (c *CircleGame) reload()  {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	sb, err := ioutil.ReadFile("research/circle/circle.kage")
	if err != nil {
		panic("could not load shader, check if it's in the right place")
	}

	s, err := ebiten.NewShader(sb)
	if err != nil {
		panic(err)
	}

	c.shader = s
}

func (c *CircleGame) Init() { c.reload() }
func (c *CircleGame) Shutdown() {}

func (c *CircleGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		c.reload()
	}

	return nil
}

func (c *CircleGame) Draw(screen *ebiten.Image) {
	if c.shader == nil {
		return
	}

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms["Side"] = float32(SideLength)
	w, h := screen.Size()
	screen.DrawRectShader(w, h, c.shader, op)
}

func (c CircleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(SideLength, SideLength)
	helpers.RunGame(new(CircleGame))
}

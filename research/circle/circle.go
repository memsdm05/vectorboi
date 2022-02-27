package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
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

	op := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Side": float32(SideLength),
			"Color": helpers.Color2Slice(colornames.Red),
		},
	}

	//w, h := screen.Size()
	screen.DrawRectShader(SideLength, SideLength, c.shader, op)
}

func (c CircleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(400, 400)
	helpers.RunGame(new(CircleGame))
}

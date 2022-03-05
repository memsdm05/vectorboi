package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"io/ioutil"
	"log"
	"vectorboi/helpers"
)

const SideLength = 400

type CircleGame struct {
	shader *ebiten.Shader
	middle cp.Vector
}

func (c *CircleGame) reload() {
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

func (c *CircleGame) Init()     {
	c.reload()
	c.middle = cp.Vector{SideLength / 2, SideLength / 2}
}
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

	x, y := ebiten.CursorPosition()
	mouse := cp.Vector{float64(x), float64(y)}
	radius := c.middle.Sub(mouse).Clamp(SideLength / 2).Length()
	side := int(radius * 2 + 5)

	temp := ebiten.NewImage(side, side)
	temp.DrawRectShader(side, side, c.shader, &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Side":  float32(side),
			"Color": helpers.Color2Slice(colornames.Red),
		},
	})
	geom := ebiten.GeoM{}
	geom.Translate(c.middle.X + 2 - float64(side) / 2, c.middle.Y + 1 - float64(side) / 2)
	screen.DrawImage(temp, &ebiten.DrawImageOptions{GeoM: geom})

	ebitenutil.DrawLine(screen, SideLength / 2, 0, SideLength / 2, SideLength, colornames.Aqua)
	ebitenutil.DrawLine(screen, 0, SideLength / 2, SideLength, SideLength / 2, colornames.Aqua)
}

func (c CircleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(400, 400)
	helpers.RunGame(new(CircleGame))
}

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"vectorboi/helpers"
)

const SideLength = 400

var (
	Shader *ebiten.Shader
	//shiftx = float64(1.81)
	//shifty = float64(2.67)
	shiftx = float64(0)
	shifty = float64(0)
)

type CircleGame struct {
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

	Shader = s
}

func (c *CircleGame) Init() {
	c.reload()
	c.middle = cp.Vector{200, 200}
}
func (c *CircleGame) Shutdown() {}

func (c *CircleGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		c.reload()
	}

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		shifty -= 0.01
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		shifty += 0.01
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		shiftx -= 0.01
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		shiftx += 0.01
	}

	return nil
}

func DrawCircle(dst *ebiten.Image, pos cp.Vector, radius float64, color color.Color) {
	side := int(math.Ceil(radius) * 2)
	halfside := float64(side / 2)

	temp := ebiten.NewImage(side, side)
	temp.DrawRectShader(side, side, Shader, &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Color": helpers.Color2Slice(color),
		},
	})

	geom := ebiten.GeoM{}
	geom.Translate(pos.X-halfside+shiftx, pos.Y-halfside+shifty)
	dst.DrawImage(temp, &ebiten.DrawImageOptions{GeoM: geom})
}

func (c *CircleGame) Draw(screen *ebiten.Image) {
	if Shader == nil {
		return
	}

	x, y := ebiten.CursorPosition()
	mouse := cp.Vector{float64(x), float64(y)}
	radius := c.middle.Sub(mouse).Clamp(SideLength / 2).Length()
	DrawCircle(screen, c.middle, radius, colornames.Red)

	ebitenutil.DrawLine(screen, SideLength/2, 0, SideLength/2, SideLength, colornames.Aqua)
	ebitenutil.DrawLine(screen, 0, SideLength/2, SideLength, SideLength/2, colornames.Aqua)
	ebitenutil.DebugPrint(screen, fmt.Sprint(shiftx, shifty))
}

func (c CircleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetScreenTransparent(true)
	helpers.RunGame(new(CircleGame))
}

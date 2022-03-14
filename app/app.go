package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"vectorboi/app/creature"
	"vectorboi/helpers"
)

const TimeStep = 1 / 60.

type Game struct {
	camera       *helpers.Camera
	space        *cp.Space
	testCreature *creature.Creature
}

func (g *Game) Init() {
	g.camera = helpers.NewCamera()
	g.camera.Position.X = -30
	g.camera.Position.Y = -10
	g.camera.SetZoom(0.8)

	g.space = cp.NewSpace()
	g.testCreature = creature.NewRandomCreature(5)
	g.testCreature.CreatePhysicsBody(g.space)
}

func (g *Game) Shutdown() {}

func (g *Game) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		g.camera.Position.Y -= 1
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		g.camera.Position.Y += 1
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		g.camera.Position.X -= 1
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		g.camera.Position.X += 1
	}

	_, ywheel := ebiten.Wheel()
	g.camera.Scale += ywheel / 100

	//g.space.Step(TimeStep)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.camera.Render(screen, g.space)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

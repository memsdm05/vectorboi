package main

import (
	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"vectorboi/helpers"
)

const PhysicsStep = 1. / 60

type PhysicsTestGame struct {
	world box2d.B2World
	ground *box2d.B2Body
	fallingBox *box2d.B2Body
}

func (p *PhysicsTestGame) Init()  {
	p.world = box2d.MakeB2World(box2d.MakeB2Vec2(0, -10))

	groundDef := box2d.NewB2BodyDef()
	groundDef.Position.Set(0, -10)
	groundBox := box2d.NewB2PolygonShape()
	groundBox.SetAsBox(50, 10)
	p.ground = box2d.NewB2Body(groundDef, &p.world)
	p.ground.CreateFixture(groundBox, 0)

	fallingBoxDef := box2d.NewB2BodyDef()
	fallingBoxDef.Type = box2d.B2BodyType.B2_dynamicBody
	fallingBoxDef.Position.Set(0, 4)
	fallingBoxBox := box2d.NewB2PolygonShape()
	fallingBoxBox.SetAsBox(1,1)
	p.fallingBox = box2d.NewB2Body(fallingBoxDef, &p.world)
	p.fallingBox.CreateFixture(fallingBoxBox, 1)
}

func (p *PhysicsTestGame) Shutdown()  {}

func (p *PhysicsTestGame) Update() error {
	p.world.Step(PhysicsStep, 10, 10)
	return nil
}

func (p *PhysicsTestGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "lol")
	box := ebiten.NewImage(10, 10)
}

func (p *PhysicsTestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	helpers.RunGame(new(PhysicsTestGame))
}
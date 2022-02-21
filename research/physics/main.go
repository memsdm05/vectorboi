package main

import (
	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"vectorboi/helpers"
)

const PhysicsStep = 1. / 60

type PhysicsTestGame struct {
	world *box2d.B2World
	ground *box2d.B2Body
	fallingBox *box2d.B2Body
}

func (p *PhysicsTestGame) Init()  {
	p.world.M_gravity.Set(0, -0.001)

	body := box2d.NewB2BodyDef()
	body.Type = box2d.B2BodyType.B2_dynamicBody
	p.fallingBox = box2d.NewB2Body(body, p.world)
}

func (p *PhysicsTestGame) Shutdown()  {

}

func (p *PhysicsTestGame) Update() error {
	//p.world.Step()
	return nil
}

func (p *PhysicsTestGame) Draw(screen *ebiten.Image) {
	panic("implement me")
}

func (p *PhysicsTestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 380
}

func main() {
	helpers.RunGame(new(PhysicsTestGame))
}
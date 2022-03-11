package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"vectorboi/app/creature"
	"vectorboi/helpers"
)

const TimeStep = 1 / 60.

type SimpleGame struct {
	camera *helpers.Camera

	space *cp.Space
	testCreature *creature.Creature
}

func (s *SimpleGame) Init() {
	ebiten.SetWindowTitle("vectorboi")
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowSize(1920, 1080)

	s.space = cp.NewSpace()
	s.testCreature = creature.NewRandomCreature(5)
	s.testCreature.CreatePhysicsBody(s.space)
}

func (s *SimpleGame) Shutdown() {}

func (s *SimpleGame) Update() error {
	s.space.Step(TimeStep)
	return nil
}

func (s *SimpleGame) Draw(screen *ebiten.Image) {
	s.space.EachShape(func(shape *cp.Shape) {
		screen.DrawImage(shape.UserData.(helpers.CameraObject).Draw(1))
	})
}

func (s *SimpleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
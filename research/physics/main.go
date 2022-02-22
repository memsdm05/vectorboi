package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"image/color"
	"vectorboi/helpers"
)

type vec = cp.Vector
const (
	TimeStep = 1. / 60
	PPM = float64(10)
)



// ball properties
var (
	num = 3
	radius = 1.828 / 2
	mass = 83.9146
	moment         = cp.MomentForCircle(mass, 0, radius, vec{})
	initialBallPos = vec{20, 5}
	initialBallVel = vec{-2, -1}
)

// ground properties
var (
	groundA = vec{X: 0, Y: 30}
	groundB = vec{X: 40, Y: 40}
)



var (
	circleShader = helpers.MustLoadShader("research/data/circle.vert")

	red = colornames.Red
	orange = colornames.Orange
)

func pixelize(v vec) vec {
	return v.Mult(PPM)
}

func drawCircle(img *ebiten.Image, pos vec, r float64, color color.Color) {
	op := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]interface{}{
			"Radius": float32(r),
			"Color": helpers.Color2Slice(color),
		},
	}

	d := int(r * 2)
	cimg := ebiten.NewImage(d, d)
	cimg.DrawRectShader(d, d, circleShader, op)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(pos.X - r, pos.Y - r)
	img.DrawImage(cimg, op2)
}

type PhysicsTestGame struct {
	space *cp.Space
	ball *cp.Body
	ground *cp.Shape

	paused bool
}

func (p *PhysicsTestGame) Init()  {
	p.space = cp.NewSpace()
	p.space.SetGravity(vec{X: 0, Y: 9.8})
	p.space.Iterations = 10

	ground := cp.NewSegment(p.space.StaticBody, groundA, groundB, 0)
	ground.SetFriction(1)
	p.ground = p.space.AddShape(ground)
	//p.ground.SetElasticity(0.9)

	p.ball = p.space.AddBody(cp.NewBody(mass, moment))
	p.ball.SetPosition(initialBallPos)
	p.ball.SetVelocityVector(initialBallVel)
	ballShape := p.space.AddShape(cp.NewCircle(p.ball, radius, vec{}))
	ballShape.SetFriction(0.7)
}

func (p *PhysicsTestGame) Shutdown()  {}

func (p *PhysicsTestGame) Update() error {
	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeyR):
		circleShader = helpers.MustLoadShader("research/data/circle.vert")
	case inpututil.IsKeyJustReleased(ebiten.KeySpace):
		p.paused = !p.paused
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		p.ball.SetPosition(initialBallPos)
		p.ball.SetVelocityVector(initialBallVel)
	}


	if !p.paused {
		p.space.Step(TimeStep)
	}

	return nil
}

func (p *PhysicsTestGame) Draw(screen *ebiten.Image) {
	ball := pixelize(p.ball.Position())
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.5v; %.5v, %.5v", ebiten.CurrentFPS(), ball.X, ball.Y))

	a := pixelize(groundA)
	b := pixelize(groundB)
	ebitenutil.DrawLine(screen, a.X, a.Y, b.X, b.Y, red)
	drawCircle(screen, ball, radius * PPM, orange)
	//ebitenutil.DrawRect(screen, ball.X, ball.Y, 10, 10, Red)

}

func (p *PhysicsTestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	//ebiten.SetWindowResizable(true)
	helpers.RunGame(new(PhysicsTestGame))
}
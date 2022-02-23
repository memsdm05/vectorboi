package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"math"
	"vectorboi/helpers"
)

type vec = cp.Vector
const (
	TimeStep = 1. / 60
)

var PPM = float64(10)


// mutli properties
var (
	num = 20
	radius = 1.
	mass = 10.
	initialBallPos = vec{20, 5}
	initialBallVel = vec{-2, -1}
)

// ground properties
var (
	groundA = vec{X: 0, Y: 30}
	groundB = vec{X: 40, Y: 40}
)


var (
	red = colornames.Red
	orange = colornames.Orange
)

func pixelize(v vec) vec {
	return v.Mult(PPM)
}

type MultishapeGame struct {
	space *cp.Space
	multi *cp.Body
	ground *cp.Shape

	off vec
	last vec

	paused bool
}

func addShape(space *cp.Space, body *cp.Body, where vec) {
	c := space.AddShape(cp.NewCircle(body, radius, where))
	c.SetMass(mass)
	c.SetFriction(0.7)
	//c.SetFilter(cp.ShapeFilter{Group: 1})
}

func (p *MultishapeGame) Init()  {
	p.space = cp.NewSpace()
	p.space.SetGravity(vec{X: 0, Y: 9.8})

	ground := cp.NewSegment(p.space.StaticBody, groundA, groundB, 0)
	ground.SetFriction(1)
	p.ground = p.space.AddShape(ground)
	//p.ground.SetElasticity(0.9)

	// Creates an infinite loop (adding to body then iterating over body's shapes to add to space reads the shape
	// to the body, cycle continues)
	//
	// FIXED: dont iterate over body's shapes in order to add them to space

	p.multi = cp.NewBody(0, 0)
	//p.multi.SetVelocityVector(initialBallVel)

	fnum := float64(num)
	for i := float64(0); i < fnum; i++ {
		addShape(p.space, p.multi, cp.ForAngle(2 * math.Pi / fnum * i).Mult(10))
	}
	//addShape(p.space, p.multi, vec{5, 0})
	//addShape(p.space, p.multi, vec{-5, 0})
	//addShape(p.space, p.multi, vec{0, -10})
	//p.multi.AccumulateMassFromShapes()
	p.space.AddBody(p.multi)
	p.multi.SetPosition(initialBallPos)
}

func (p *MultishapeGame) Shutdown() {}

func (p *MultishapeGame) Update() error {
	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeyR):
		helpers.CircleShader = helpers.MustLoadShader("public/circle.kage")
	case inpututil.IsKeyJustReleased(ebiten.KeySpace):
		p.paused = !p.paused
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		p.multi.SetPosition(initialBallPos)
		p.multi.SetVelocityVector(initialBallVel)
		p.multi.SetAngle(0)
		//p.multi.SetAngularVelocity(2)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tx, ty := ebiten.CursorPosition()
		x, y := float64(tx), float64(ty)
		if p.last.Length() > 0 {
			p.off.X += x - p.last.X
			p.off.Y += y - p.last.Y
		}
		p.last.X, p.last.Y = x, y
	}
	p.last.Mult(0)

	_, dppm := ebiten.Wheel()
	PPM += dppm
	if PPM < 1 {
		PPM = 1
	}

	if !p.paused {
		p.space.Step(TimeStep)
	}

	return nil
}

func (p *MultishapeGame) Draw(screen *ebiten.Image) {
	multi := pixelize(p.multi.Position())
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.5v; %.5v, %.5v", ebiten.CurrentTPS(), multi.X, multi.Y))

	a := pixelize(groundA).Add(p.off)
	b := pixelize(groundB).Add(p.off)
	ebitenutil.DrawLine(screen, a.X, a.Y, b.X, b.Y, red)
	p.multi.EachShape(func(shape *cp.Shape) {
		helpers.DrawCircle(
			screen,
			pixelize(shape.Class.(*cp.Circle).TransformC()).Add(p.off),
			radius * PPM,
			orange)
	})
	//helpers.DrawCircle(screen, multi.Add(p.off), radius * PPM, orange)
	//ebitenutil.DrawRect(screen, ball.X, ball.Y, 10, 10, Red)

}

func (p *MultishapeGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowResizable(true)
	helpers.RunGame(new(MultishapeGame))
}
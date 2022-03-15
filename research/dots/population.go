package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"sort"

	_ "image/jpeg"
	_ "image/png"
)

const (
	GenerationTime = 15
	KickTime = 1.2
)

var (
	gooblegop, _ = ebitenutil.NewImageFromURL("https://images.thdstatic.com/productImages/bf4a1fd8-aca2-4f0f-94a6-d188cf1ba7ea/svn/black-fence-armor-deck-post-caps-pf-acn-252b-4f_600.jpg")
	snorb, _     = ebitenutil.NewImageFromURL("https://www.pikpng.com/pngl/b/190-1905158_scary-eye-png-transparent-creepy-eyeball-png-clipart.png")
)

var gw, gh = gooblegop.Size()
var sw, sh = snorb.Size()

type Eval func(dot *Dot, population *Population) float64

func ConstantFitness(dot *Dot, pop *Population) float64 {
	return 1
}

func RandomFitness(dot *Dot, pop *Population) float64 {
	return uniform(1, 100)
}

func DistanceFitness(dot *Dot, pop *Population) float64 {
	constraint := pop.Spawn.Distance(pop.Target)
	return 1. - (dot.body.Position().Distance(pop.Target) / constraint)
}

type Population struct {
	Dots          []*Dot
	Space         *cp.Space
	OnMove        int
	Width, Height int
	Generation    int
	Time          float64

	Spawn   cp.Vector
	Target  cp.Vector
	fitness Eval
	running bool

	bestDot        *Dot
	bestDotFitness float64
}

func (p *Population) Len() int           { return len(p.Dots) }
func (p *Population) Less(i, j int) bool { return p.Dots[i].fitness <= p.Dots[j].fitness }
func (p *Population) Swap(i, j int)      { p.Dots[i], p.Dots[j] = p.Dots[j], p.Dots[i] }

func NewRandomPopulation(num, width, height int, fitness Eval) *Population {
	if fitness == nil {
		fitness = RandomFitness
	}

	pop := &Population{
		Dots:   make([]*Dot, num, num),
		Space:  cp.NewSpace(),
		Width:  width,
		Height: height,
		Spawn: cp.Vector{
			X: Width / 2,
			Y: Height - Height/10,
		},
		fitness: fitness,
	}

	pop.Space.SleepTimeThreshold = cp.INFINITY
	pop.Space.UseSpatialHash(2, 100)
	for i := 0; i < num; i++ {
		ndot := NewRandomDot()
		ndot.CreatePhysicsBody(pop.Space)
		pop.Dots[i] = ndot
	}

	pop.reset()
	return pop
}

func (p *Population) reset() {
	p.Generation++
	p.Time = 0
	p.OnMove = 0
	for _, dot := range p.Dots {
		dot.body.SetAngle(0)
		dot.body.SetTorque(0)
		dot.body.SetAngularVelocity(0)
		dot.body.SetPosition(p.Spawn)
		dot.body.SetVelocity(0, 0)
		dot.body.SetForce(cp.Vector{})
		p.unkill(dot)
	}
}

func (p *Population) evolve() {
	// evaluate fitness
	for _, dot := range p.Dots {
		dot.fitness = p.fitness(dot, p)
	}

	// sort by fitness
	sort.Sort(p)

	// kill lower half
	//middle := p.Dots[len(p.Dots) / 2].fitness
	//for i, dot := range p.Dots {
	//	if dot.fitness < middle {
	//		p.Dots[i] = nil // RIP
	//	}
	//}

	// todo crossover, mutate

	p.reset()
}

func (p *Population) unkill(dot *Dot) {
	dot.dead = false
	p.Space.Activate(dot.body)
}

func (p *Population) kill(dot *Dot) {
	dot.dead = true
	p.Space.Deactivate(dot.body)
}

func (p *Population) IsBest(dot *Dot) bool {
	return dot == p.bestDot
}

func (p *Population) Step(dt float64) {
	if p.Time > GenerationTime { // cahnge this for
		p.evolve()
	}

	p.Space.Step(dt)

	hitnow := false
	p.Time += dt
	if p.Time >= float64(p.OnMove) * KickTime { // change this for faster/slower kicking
		hitnow = true
	}

	p.bestDot = nil
	p.bestDotFitness = 0

	for _, dot := range p.Dots {
		//if p.OnMove < len(dot.moves) {
		//	dot.body.ApplyImpulseAtLocalPoint(dot.moves[p.OnMove], cp.Vector{})
		//}
		if hitnow && p.OnMove < len(dot.moves) {
			dot.body.ApplyImpulseAtLocalPoint(dot.moves[p.OnMove], cp.Vector{})
			//dot.body.ApplyForceAtLocalPoint(, cp.Vector{})
		}

		pos := dot.body.Position()
		if pos.X > float64(p.Width)-2 || pos.X < 2 || pos.Y > float64(p.Height)-2 || pos.Y < 2 {
			p.kill(dot)
		}

		if f := p.fitness(dot, p); f > p.bestDotFitness {
			p.bestDot = dot
			p.bestDotFitness = f
		}
	}

	if hitnow {
		p.OnMove++
	}
}

func (p *Population) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i, dot := range p.Dots {
		pos := dot.body.Position()

		op.GeoM.Reset()
		op.GeoM.Translate(float64(-sw) / 2, float64(-sh) / 2)
		op.GeoM.Scale(0.02, 0.02)
		op.GeoM.Translate(pos.X, pos.Y)

		op.ColorM.Reset()
		switch {
		case dot.dead:
			op.ColorM.RotateHue(p.Time * 10 + float64(i) * 3)
		case p.IsBest(dot):
			op.ColorM.Invert()
		default:
			op.ColorM.RotateHue(p.Time)
		}

		dst.DrawImage(snorb, op)
	}

	op.GeoM.Reset()
	op.GeoM.Translate(float64(-gw) / 2, float64(-gh) / 2)
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(p.Target.X, p.Target.Y)
	dst.DrawImage(gooblegop, op)
}

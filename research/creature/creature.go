package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"time"
	"vectorboi/app/creature"
	"vectorboi/helpers"
)

var mainCreature *creature.Creature
var NumNodes = 5

type CreatureViewGame struct{}

func reload() {
	rand.Seed(time.Now().UnixNano())
	mainCreature = creature.NewRandomCreature(NumNodes)
}

func (c CreatureViewGame) Init() {
	reload()
}

func (c CreatureViewGame) Shutdown() {}

func (c CreatureViewGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		reload()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		NumNodes++
		reload()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		NumNodes--
		reload()
	}

	return nil
}

func (c CreatureViewGame) Draw(screen *ebiten.Image) {
	w, h := ebiten.WindowSize()
	offset := cp.Vector{
		X: float64(w) / 2,
		Y: float64(h) / 2,
	}

	mainCreature.Clone().Body[0].Do(func(node *creature.Node) {
		var c color.Color
		if node.Root() {
			c = colornames.Red
		} else {
			c = colornames.Whitesmoke
		}
		helpers.DrawCircle(screen, node.Position().Add(offset), node.Radius, c)
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("num: %v, %s", NumNodes, mainCreature.Body[0]))
}

func (c CreatureViewGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowResizable(true)
	helpers.RunGame(new(CreatureViewGame))
}

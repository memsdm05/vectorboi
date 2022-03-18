package main

import (
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
	"time"
	"vectorboi/app"
	"vectorboi/app/dot"
	"vectorboi/helpers"
)

const TimeStep = 1 / 60.

var RequiredFolders = [...]string{
	"snapshots",
	"scenarios",
}

type SimulationGame struct {
	pop   *dot.Population
	editor *app.Editor
	editing bool
}

func (s *SimulationGame) Init() {
	for _, folder := range RequiredFolders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.Mkdir(folder, 777)
		}
	}

	//myscenario := dot.DefaultScenario
	//myscenario.Walls = []structures.KillWall{
	//	structures.MakeKillWall(2, 200, 300, 300),
	//	structures.MakeKillWall(float64(myscenario.Width), 100, 220, 200),
	//	structures.MakeKillWall(200, 200, float64(myscenario.Width-200), 200),
	//}
	s.pop = dot.NewPopulation(dot.DefaultScenario)
	s.editor = app.NewEditor(&s.pop.Scenario)
	ebiten.SetWindowSize(s.pop.Scenario.Width, s.pop.Scenario.Height)

	//start := time.Now()
	//for time.Since(start) < 15*time.Second {
	//	s.pop.Step(1. / 30)
	//}
}

func (s *SimulationGame) Shutdown() {}

func (s *SimulationGame) Update() error {
	if !s.editing {
		s.pop.Step(TimeStep)
	} else {
		s.editor.Interact()
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		s.pop.Paused = !s.pop.Paused
	case inpututil.IsKeyJustPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyControl):
		filename := fmt.Sprintf("snapshots/%v-snapshot.json", time.Now().UnixMilli())
		f, _ := os.Create(filename)
		enc := json.NewEncoder(f)
		enc.SetIndent("", "    ")
		enc.Encode(s.pop)
	case inpututil.IsKeyJustPressed(ebiten.KeyE):
		s.editing = !s.editing
	}

	return nil
}

func (s *SimulationGame) Draw(screen *ebiten.Image) {
	if !s.editing {
		s.pop.Draw(screen)
		msg := fmt.Sprintf("kick %v, gen %v, dt %.2f", s.pop.KickIndex, s.pop.Generation, s.pop.Time)
		ebitenutil.DebugPrint(screen, msg)
	} else {
		x, y := ebiten.CursorPosition()
		msg := fmt.Sprintf("(%d, %d)", x, y)
		s.editor.Draw(screen)
		ebitenutil.DebugPrint(screen, msg)
	}
}

func (s *SimulationGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	helpers.RunGame(new(SimulationGame))
}

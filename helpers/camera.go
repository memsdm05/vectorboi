package helpers

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
)


type CameraObject interface {
	ScaleMe() bool
	Draw(zoom float64) (*ebiten.Image, *ebiten.DrawImageOptions)
}

// Camera does things
type Camera struct {
	Position cp.Vector
	Scale     float64
	frustum  cp.BB

	minzoom float64
	maxzoom float64

	count int
}

func NewCamera() *Camera {
	return &Camera{
		Scale:    1,
		minzoom: 0.01,
		maxzoom: 10,
	}
}

func (c *Camera) SetZoom(zoom float64)  {
	if zoom < c.minzoom {
		zoom = c.minzoom
	}
	if zoom > c.maxzoom {
		zoom = c.maxzoom
	}

	c.Scale = 1 / zoom
}

func (c *Camera) ToScreen(global cp.Vector) cp.Vector {
	return global.Mult(c.Scale).Sub(c.Position)
}

func (c *Camera) ToGlobal(screen cp.Vector) cp.Vector {
	return cp.Vector{}
}

func (c *Camera) drawEach(shape *cp.Shape, data interface{}) {
	dst := data.(*ebiten.Image)
	spos := c.ToScreen(shape.BB().Center())

	// make sure that the shape's UserData is a camera object
	if co, ok := shape.UserData.(CameraObject); ok {
		img, op := co.Draw(c.Scale)

		if op == nil {
			op = &ebiten.DrawImageOptions{}
		}

		w, h := img.Size()
		op.GeoM.Translate(-float64(w) / 2, -float64(h) / 2)
		op.GeoM.Translate(spos.X, spos.Y) // todo translate by screen coordinates
		if co.ScaleMe() { op.GeoM.Scale(c.Scale, c.Scale) }
		dst.DrawImage(img, op)
	}

	c.count++
}

/*
BB{
		L: c.X - hw,
		B: c.Y - hh,
		R: c.X + hw,
		T: c.Y + hh,
	}
 */

func (c *Camera) Render(dst *ebiten.Image, space *cp.Space) {
	// resize frustum
	w, h := dst.Size()
	hw, hh := float64(w) / 2 * c.Scale, float64(h) / 2 * c.Scale
	c.frustum.L = c.Position.X - hw
	c.frustum.B = c.Position.Y - hh
	c.frustum.R = c.Position.X + hw
	c.frustum.T = c.Position.Y + hh

	c.count = 0

	// draw each shape within frustum
	space.BBQuery(c.frustum, cp.SHAPE_FILTER_ALL, c.drawEach, dst)

	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f
Num of sprites: %d`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), c.count)
	ebitenutil.DebugPrint(dst, msg)
}
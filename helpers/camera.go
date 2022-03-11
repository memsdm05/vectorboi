package helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)


type CameraObject interface {
	Draw(zoom float64) (*ebiten.Image, *ebiten.DrawImageOptions)
}

// Camera does things
type Camera struct {
	frustum cp.BB
	pos     cp.Vector
	zoom float64

	minzoom float64
	maxzoom float64
}

func (c *Camera) SetZoom(zoom float64)  {
	if zoom < c.minzoom {
		zoom = c.minzoom
	}
	if zoom > c.maxzoom {
		zoom = c.maxzoom
	}

	c.zoom = zoom
}

//func (c *Camera) ToScreenPos(global cp.Vector) cp.Vector {
//	return global.Mult(c.zoom).Add(c.pos)
//}
//
//func (c *Camera) ToGlobalPos(screen cp.Vector) cp.Vector {
//
//}

func (c *Camera) drawEach(shape *cp.Shape, data interface{}) {
	dst := data.(*ebiten.Image)
	shapePos := shape.BB().Center()

	// make sure that the shape's UserData is a camera object
	if co, ok := shape.UserData.(CameraObject); ok {
		img, op := co.Draw(c.zoom)

		if op == nil {
			op = &ebiten.DrawImageOptions{}
		}

		op.GeoM.Translate(shapePos.X, shapePos.Y) // todo translate by screen coordinates
		dst.DrawImage(img, op)
	}
}

/*
BB{
		L: c.X - hw,
		B: c.Y - hh,
		R: c.X + hw,
		T: c.Y + hh,
	}
 */

func (c *Camera) Render(dst *ebiten.Image, space *cp.Space) error {
	// resize frustum
	w, h := dst.Size()
	hw, hh := float64(w) / 2 * c.zoom, float64(h) / 2 * c.zoom
	c.frustum.L = c.pos.X - hw
	c.frustum.B = c.pos.Y - hh
	c.frustum.R = c.pos.X + hw
	c.frustum.T = c.pos.Y + hh

	// draw each shape within frustum
	space.BBQuery(c.frustum, cp.SHAPE_FILTER_ALL, c.drawEach, dst)
	return nil
}
package helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)


type CameraObject interface {
	ScaleMe() bool
	Draw(zoom float64) *ebiten.Image
}

// Camera
type Camera struct {
	pos cp.Vector
	zoom float64
}

func (c *Camera) SetZoom()  {
	
}

func (c *Camera) drawEach(shape *cp.Shape, data interface{}) {
	dst := data.(*ebiten.Image)
	if co, ok := shape.UserData.(CameraObject); ok {
		op := &ebiten.DrawImageOptions{}
		if co.ScaleMe() {
			op.GeoM.Scale(c.zoom, c.zoom)
		}
		dst.DrawImage(co.Draw(c.zoom), op)
	}
}

func (c *Camera) Render(dst *ebiten.Image, space *cp.Space) error {
	w, h := dst.Size()
	frustrum := cp.NewBBForExtents(c.pos,
		float64(w) / 2 * c.zoom,
		float64(h) / 2 * c.zoom)

	space.BBQuery(frustrum, cp.SHAPE_FILTER_ALL, c.drawEach, dst)
	return nil
}
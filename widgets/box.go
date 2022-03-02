package widgets

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

var (
	Red    = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	Green  = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	Blue   = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
	Yellow = color.NRGBA{R: 0x00, G: 0xFF, B: 0xFF, A: 0xFF}
)

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) D {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

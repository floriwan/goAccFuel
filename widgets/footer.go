package widgets

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

type FooterInfoStyle struct {
	textFont     text.Font
	textFontSize unit.Value
	textShaper   *text.Cache
	textColor    color.NRGBA
}

func FooterInfo(font text.Font, size unit.Value, color color.NRGBA, shaper *text.Cache) FooterInfoStyle {
	return FooterInfoStyle{
		textFont:     font,
		textFontSize: size,
		textShaper:   shaper,
		textColor:    color,
	}
}

func (f FooterInfoStyle) Layout(gtx C) D {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
		layout.Rigid(func(gtx C) D {
			paint.ColorOp{Color: f.textColor}.Add(gtx.Ops)
			return widget.Label{}.Layout(gtx, f.textShaper, f.textFont, f.textFontSize, "ACC 1.23.4")
		}))
}

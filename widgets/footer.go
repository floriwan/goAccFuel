package widgets

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

type FooterInfoStyle struct {
	textColor color.NRGBA
	version   string
	carModel  string
}

func FooterInfo(color color.NRGBA, carModel string, version string) FooterInfoStyle {
	return FooterInfoStyle{
		textColor: color,
		version:   version,
		carModel:  carModel,
	}
}

func (f FooterInfoStyle) Layout(gtx C) D {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
		layout.Rigid(func(gtx C) D {
			paint.ColorOp{Color: f.textColor}.Add(gtx.Ops)
			return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, f.carModel+" / ACC "+f.version)
		}))
}

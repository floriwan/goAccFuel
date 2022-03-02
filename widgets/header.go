package widgets

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

type HeaderInfoStyle struct {
	textFont     text.Font
	textFontSize unit.Value
	textShaper   *text.Cache
	textColor    color.NRGBA
}

func HeaderInfo(font text.Font, size unit.Value, color color.NRGBA, shaper *text.Cache) HeaderInfoStyle {
	return HeaderInfoStyle{
		textFont:     font,
		textFontSize: size,
		textShaper:   shaper,
		textColor:    color,
	}
}

func (h HeaderInfoStyle) Layout(gtx C) D {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
		layout.Rigid(func(gtx C) D {

			paint.ColorOp{Color: h.textColor}.Add(gtx.Ops)
			return layout.Flex{}.Layout(gtx,
				//layout.Rigid(func(gtx C) D {
				//	return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "Session")
				//}),
				layout.Rigid(func(gtx C) D {

					return layout.Inset{
						Left:  unit.Dp(5),
						Right: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.End,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize.Scale(0.5), "Session")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "95min")
							}),
						)
					})

				}),

				layout.Rigid(func(gtx C) D {

					return layout.Inset{
						Left:  unit.Dp(5),
						Right: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.End,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize.Scale(0.5), "Laps")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "73")
							}),
						)
					})
				}),

				layout.Rigid(func(gtx C) D {

					return layout.Inset{
						Left: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.End,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize.Scale(0.5), "Fuel")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "145l")
							}),
						)
					})
				}),
			)

		}))
	/*
		return layout.Flex{}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return ColorBox(gtx, gtx.Constraints.Min, Blue)
			}),
			layout.Rigid(func(gtx C) D {
				paint.ColorOp{Color: h.textColor}.Add(gtx.Ops)
				return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "Session Info")
			}))*/
}

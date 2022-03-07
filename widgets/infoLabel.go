package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

func InfoLabel(gtx C, label string, info string) D {
	return layout.Inset{
		Left:  unit.Dp(5),
		Right: unit.Dp(5),
	}.Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.End,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, label)
			}),
			layout.Rigid(func(gtx C) D {
				return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), info)
			}),
		)

	})
}

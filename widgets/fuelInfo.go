package widgets

import (
	"fmt"
	"image/color"
	"time"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

type FuelInfoStyle struct {
	textColor   color.NRGBA
	fuelLevel   float32
	fuelPerLap  float32
	sessionTime time.Duration
	lapTime     time.Duration
}

func FuelInfo(color color.NRGBA,
	fuelLevel float32, fuelPerLap float32, sessionTime time.Duration, lapTime time.Duration) FuelInfoStyle {
	return FuelInfoStyle{
		textColor:   color,
		fuelLevel:   fuelLevel,
		fuelPerLap:  fuelPerLap,
		sessionTime: sessionTime,
		lapTime:     lapTime,
	}
}

func (f FuelInfoStyle) Layout(gtx C) D {

	return layout.Flex{}.Layout(gtx,

		layout.Rigid(func(gtx C) D {

			paint.ColorOp{Color: textColor}.Add(gtx.Ops)

			return layout.Flex{}.Layout(gtx,

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
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Lap Time")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), fmtLapTime(f.lapTime))
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
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Session Time")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), fmtDuration(f.sessionTime))
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
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Fuel Level")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), fmt.Sprintf("%.1f", f.fuelLevel))
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
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Fuel Per Lap")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), fmt.Sprintf("%.2f", f.fuelPerLap))
							}),
						)

					})
				}),
			)
		}),
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
	)

}

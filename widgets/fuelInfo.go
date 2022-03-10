package widgets

import (
	"fmt"
	"image/color"
	"time"

	"gioui.org/layout"
	"gioui.org/op/paint"
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
					return InfoLabel(gtx, "Lap Time", fmtLapTime(f.lapTime))
				}),

				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Session Time", fmtDuration(f.sessionTime))
				}),

				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Fuel Level", fmt.Sprintf("%.1f", f.fuelLevel))
				}),

				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Fuel Per Lap", fmt.Sprintf("%.2f", f.fuelPerLap))
				}),
			)
		}),
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
	)

}

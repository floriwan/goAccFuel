package widgets

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

type BodyStyle struct {
	RaceProgress         float32
	RaceProgressWithFuel float32
	FuelLevel            float32
	FuelPerLap           float32
	SessionTime          time.Duration
	LaptTime             time.Duration
	BoxLap               int
	RefuelLevel          float32
}

func BodyInfo(color color.NRGBA,
	raceProgress float32,
	raceProgressWithFuel float32,
	fuelLevel float32,
	fuelPerLap float32,
	sessionTime time.Duration,
	lapTime time.Duration,
	boxLap int,
	refuelLevel float32) BodyStyle {
	return BodyStyle{
		RaceProgress:         raceProgress,
		RaceProgressWithFuel: raceProgressWithFuel,
		FuelLevel:            fuelLevel,
		FuelPerLap:           fuelPerLap,
		SessionTime:          sessionTime,
		LaptTime:             lapTime,
		BoxLap:               boxLap,
		RefuelLevel:          refuelLevel,
	}
}
func (bs BodyStyle) Layout(gtx C) D {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 40
			maxX := float32(gtx.Constraints.Max.X)
			progressPx := (float32(maxX) * (float32(bs.RaceProgressWithFuel))) / float32(100)
			xlabel := 0 // width of the label left of the pit stop line

			return layout.Flex{}.Layout(gtx,

				layout.Rigid(func(gtx C) D {
					paint.ColorOp{Color: textColor}.Add(gtx.Ops)
					return layout.Inset{
						Left:  unit.Dp(5),
						Right: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.End,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								dim := widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Box Lap")
								xlabel += dim.Size.X
								return dim
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), strconv.Itoa(bs.BoxLap))
							}),
						)

					})
				}),

				layout.Rigid(func(gtx C) D {
					paint.ColorOp{Color: textColor}.Add(gtx.Ops)
					return layout.Inset{
						Left:  unit.Dp(5),
						Right: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.End,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								dim := widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Refuel")
								//xlabel += dim.Size.X
								return dim
							}),
							layout.Rigid(func(gtx C) D {
								dim := widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), fmt.Sprintf(" %.1f", bs.RefuelLevel))
								xlabel += dim.Size.X
								return dim
							}),
						)

					})
				}),

				layout.Rigid(func(gtx C) D {

					rect := clip.RRect{
						Rect: f32.Rectangle{Min: f32.Point{X: progressPx - float32(xlabel) - 2 - 20, Y: 10},
							Max: f32.Point{X: progressPx - float32(xlabel) + 2 - 20, Y: 40}},
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, Red, rect)
					return layout.Dimensions{Size: gtx.Constraints.Max}
				}))

			//max := gtx.Constraints.Max
			//max.Y = 40
			//return ColorBox(gtx, max, Green)

			//minX := float32(gtx.Constraints.Min.X)
			//minY := float32(gtx.Constraints.Min.Y)

			//maxX := float32(gtx.Constraints.Max.X)
			//maxY := float32(gtx.Constraints.Max.Y)

		}),
		layout.Rigid(func(gtx C) D {
			return ProgressBarInfo(bs.RaceProgress, bs.RaceProgressWithFuel).Layout(gtx)
		}),
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Max, Green)
		}),
		layout.Rigid(func(gtx C) D {
			return FuelInfo(textColor,
				bs.FuelLevel,
				bs.FuelPerLap,
				bs.SessionTime,
				bs.LaptTime).Layout(gtx)

		}))

}
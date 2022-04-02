package widgets

import (
	"fmt"
	"goAccFuel/acc"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type BodyStyle struct {
	accData acc.AccData
}

func BodyInfo(color color.NRGBA,
	accData acc.AccData) BodyStyle {
	return BodyStyle{
		accData: accData,
	}
}
func (bs BodyStyle) Layout(gtx C) D {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 50
			maxX := float32(gtx.Constraints.Max.X)
			progressPx := (float32(maxX) * (float32(bs.accData.ProgressWithFuel))) / float32(100)
			xlabel := 0 // width of the label left of the pit stop line

			return layout.Flex{}.Layout(gtx,
				/*
					layout.Rigid(func(gtx C) D {
										dim := InfoLabel(gtx, "Box Lap", fmt.Sprintf(" %.1f", strconv.Itoa(bs.BoxLap))
					xlabel += dim.Size.X
					return dim
					}),
				*/

				layout.Rigid(func(gtx C) D {
					dim := InfoLabel(gtx, "Laps Done", fmt.Sprintf("%v", bs.accData.LapsDone))
					xlabel += dim.Size.X
					return dim
				}),

				layout.Rigid(func(gtx C) D {
					dim := InfoLabel(gtx, "Laps With Fuel", fmt.Sprintf("%.1f", bs.accData.LapsToGo))
					xlabel += dim.Size.X
					return dim
				}),

				layout.Rigid(func(gtx C) D {
					paint.ColorOp{Color: Red}.Add(gtx.Ops)
					if bs.accData.RefuelLevel <= 0 {
						paint.ColorOp{Color: Green}.Add(gtx.Ops)
					}
					dim := InfoLabel(gtx, "Refuel", fmt.Sprintf("%.1f", bs.accData.RefuelLevel))
					xlabel += dim.Size.X
					paint.ColorOp{Color: textColor}.Add(gtx.Ops)
					return dim
				}),

				// Refuel Bar
				layout.Rigid(func(gtx C) D {
					rect := clip.RRect{
						Rect: f32.Rectangle{Min: f32.Point{X: progressPx - float32(xlabel) - 2, Y: 40},
							Max: f32.Point{X: progressPx - float32(xlabel) + 2, Y: float32(gtx.Constraints.Max.Y)}},
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, Red, rect)
					return layout.Dimensions{Size: gtx.Constraints.Max}
				}))

		}),
		layout.Rigid(func(gtx C) D {
			return ProgressBarInfo(bs.accData.RaceProgress,
				bs.accData.ProgressWithFuel,
				bs.accData.PitWindowStart,
				bs.accData.PitWindowEnd).Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 20
			xlabel := 0
			return layout.Flex{}.Layout(gtx,

				layout.Rigid(func(gtx C) D {

					rect := clip.RRect{
						Rect: f32.Rectangle{Min: f32.Point{X: 100 - float32(xlabel), Y: 0},
							Max: f32.Point{X: 200 - float32(xlabel), Y: float32(gtx.Constraints.Max.Y - 35)}},
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, Green, rect)
					return layout.Dimensions{Size: gtx.Constraints.Max}
				}))

		}),

		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 40
			xlabel := 0
			return layout.Flex{}.Layout(gtx,

				layout.Rigid(func(gtx C) D {
					paint.ColorOp{Color: textColor}.Add(gtx.Ops)
					dim := InfoLabel(gtx, "Box Open At", fmtDuration(bs.accData.PitWindowStartTime))
					xlabel += dim.Size.X
					return dim
				}),

				layout.Rigid(func(gtx C) D {
					dim := InfoLabel(gtx, "Box Close At", fmtDuration(bs.accData.PitWindowCloseTime))
					xlabel += dim.Size.X
					return dim
				}),

				/*layout.Rigid(func(gtx C) D {

					rect := clip.RRect{
						Rect: f32.Rectangle{Min: f32.Point{X: 100 - float32(xlabel), Y: 0},
							Max: f32.Point{X: 200 - float32(xlabel), Y: float32(gtx.Constraints.Max.Y - 35)}},
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, Green, rect)
					return layout.Dimensions{Size: gtx.Constraints.Max}
				})*/)

		}),

		// add some space between box info and fuel info
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 8
			return layout.Dimensions{Size: gtx.Constraints.Max}
		}),

		layout.Rigid(func(gtx C) D {
			return FuelInfo(textColor,
				bs.accData.FuelLevel,
				bs.accData.FuelPerLap,
				bs.accData.SessionTime,
				bs.accData.LapTime).Layout(gtx)

		}))

}

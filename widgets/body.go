package widgets

import (
	"fmt"
	"image/color"
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type BodyStyle struct {
	RaceProgress         float32
	RaceProgressWithFuel float32
	FuelLevel            float32
	FuelPerLap           float32
	SessionTime          time.Duration
	LaptTime             time.Duration
	BoxLap               int
	LapsToGo             float32
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
	lapsToGo float32,
	refuelLevel float32) BodyStyle {
	return BodyStyle{
		RaceProgress:         raceProgress,
		RaceProgressWithFuel: raceProgressWithFuel,
		FuelLevel:            fuelLevel,
		FuelPerLap:           fuelPerLap,
		SessionTime:          sessionTime,
		LaptTime:             lapTime,
		BoxLap:               boxLap,
		LapsToGo:             lapsToGo,
		RefuelLevel:          refuelLevel,
	}
}
func (bs BodyStyle) Layout(gtx C) D {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 50
			maxX := float32(gtx.Constraints.Max.X)
			progressPx := (float32(maxX) * (float32(bs.RaceProgressWithFuel))) / float32(100)
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
					dim := InfoLabel(gtx, "Laps With Fuel", fmt.Sprintf(" %.1f", bs.LapsToGo))
					xlabel += dim.Size.X
					return dim
				}),

				layout.Rigid(func(gtx C) D {
					dim := InfoLabel(gtx, "Refuel", fmt.Sprintf(" %.1f", bs.RefuelLevel))
					xlabel += dim.Size.X
					return dim
				}),

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
			return ProgressBarInfo(bs.RaceProgress, bs.RaceProgressWithFuel).Layout(gtx)
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
					dim := InfoLabel(gtx, "Box Open Lap", "X")
					xlabel += dim.Size.X
					return dim
				}),

				layout.Rigid(func(gtx C) D {
					dim := InfoLabel(gtx, "Box Close Lap", "X")
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
				bs.FuelLevel,
				bs.FuelPerLap,
				bs.SessionTime,
				bs.LaptTime).Layout(gtx)

		}))

}

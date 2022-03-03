package widgets

import (
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
}

func BodyInfo(color color.NRGBA,
	raceProgress float32,
	raceProgressWithFuel float32,
	fuelLevel float32,
	fuelPerLap float32,
	sessionTime time.Duration) BodyStyle {
	return BodyStyle{
		RaceProgress:         raceProgress,
		RaceProgressWithFuel: raceProgressWithFuel,
		FuelLevel:            fuelLevel,
		FuelPerLap:           fuelPerLap,
		SessionTime:          sessionTime,
	}
}
func (bs BodyStyle) Layout(gtx C) D {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 40
			maxX := float32(gtx.Constraints.Max.X)
			progressPx := (float32(maxX) * (float32(bs.RaceProgressWithFuel))) / float32(100)
			//fmt.Printf("progress pixel %v\n", progressPx)

			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					gtx.Constraints.Max.Y = 40
					rect := clip.RRect{
						Rect: f32.Rectangle{Min: f32.Point{X: progressPx - 2, Y: 10},
							Max: f32.Point{X: progressPx + 2, Y: 40}},
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
			return ProgressBarInfo(bs.RaceProgress).Layout(gtx)
		}),
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Max, Green)
		}),
		layout.Rigid(func(gtx C) D {
			return FuelInfo(textColor,
				bs.FuelLevel,
				bs.FuelPerLap,
				bs.SessionTime).Layout(gtx)

		}))

}

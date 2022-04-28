package widgets

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type ProgressBarStyle struct {
	RaceProgress         float32 // race progress in percentage of total time
	RaceProgressWithFuel float32 // percentage of the race to reach with current fuel level
	PitWindowStart       float32 // percentage of the race when pit window will open
	PitWindowClose       float32 // percentage of the race when pit window is closing
}

func ProgressBarInfo(raceProgress float32,
	raceProgressWithFuel float32,
	pitWindowStart float32,
	pitWindowClose float32) ProgressBarStyle {
	return ProgressBarStyle{
		RaceProgress:         raceProgress,
		RaceProgressWithFuel: raceProgressWithFuel,
		PitWindowStart:       pitWindowStart,
		PitWindowClose:       pitWindowClose,
	}
}

func (f ProgressBarStyle) Layout(gtx C) D {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {

			gtx.Constraints.Max.Y = 20

			minX := float32(gtx.Constraints.Min.X)
			minY := float32(gtx.Constraints.Min.Y)

			maxX := float32(gtx.Constraints.Max.X)
			maxY := float32(gtx.Constraints.Max.Y)

			// percentage to pixel
			progressPx := (float32(maxX) * (f.RaceProgress)) / float32(100)
			progressWithFuelPx := (float32(maxX) * (float32(f.RaceProgressWithFuel))) / float32(100)

			pitWindowOpenPx := (float32(maxX) * (f.PitWindowStart)) / float32(100)
			pitWindowClosePx := (float32(maxX) * (f.PitWindowClose)) / float32(100)

			rect := clip.RRect{
				Rect: f32.Rectangle{Min: f32.Point{X: minX, Y: minY},
					Max: f32.Point{X: maxX, Y: maxY}},
			}.Op(gtx.Ops)
			paint.FillShape(gtx.Ops, LightGrey, rect)

			// pit window bar
			rect = clip.RRect{
				Rect: f32.Rectangle{Min: f32.Point{X: pitWindowOpenPx, Y: float32(gtx.Constraints.Max.Y - 5)},
					Max: f32.Point{X: pitWindowClosePx, Y: float32(gtx.Constraints.Max.Y)}},
			}.Op(gtx.Ops)
			paint.FillShape(gtx.Ops, Green, rect)

			// refuel bar
			rect = clip.RRect{
				Rect: f32.Rectangle{Min: f32.Point{X: progressWithFuelPx - 2, Y: 0},
					Max: f32.Point{X: progressWithFuelPx + 2, Y: float32(gtx.Constraints.Max.Y)}},
			}.Op(gtx.Ops)
			paint.FillShape(gtx.Ops, Red, rect)

			// race progress bar
			rect = clip.RRect{
				Rect: f32.Rectangle{Min: f32.Point{X: minX, Y: minY},
					Max: f32.Point{X: progressPx, Y: maxY}},
			}.Op(gtx.Ops)
			paint.FillShape(gtx.Ops, Wight, rect)

			return layout.Dimensions{Size: gtx.Constraints.Max}
		}))
}

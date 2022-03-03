package widgets

import (
	"gioui.org/layout"
)

type BodyStyle struct {
	RaceProgress float32
}

func BodyInfo(raceProgress float32) BodyStyle {
	return BodyStyle{RaceProgress: raceProgress}
}
func (f BodyStyle) Layout(gtx C) D {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			max := gtx.Constraints.Max
			max.Y = 40
			//fmt.Printf("fuel box 1 : %+v\n", max)
			return ColorBox(gtx, max, Green)
		}),
		layout.Rigid(func(gtx C) D {
			return ProgressBarInfo(f.RaceProgress).Layout(gtx)
		}),
		layout.Flexed(1, func(gtx C) D {
			//fmt.Printf("fuel box 3 : %+v\n", gtx.Constraints.Max)
			return ColorBox(gtx, gtx.Constraints.Max, Red)
		}),
		layout.Rigid(func(gtx C) D {
			max := gtx.Constraints.Max
			max.Y = 40
			//fmt.Printf("fuel box 1 : %+v\n", max)
			return ColorBox(gtx, max, Green)
		}))

}

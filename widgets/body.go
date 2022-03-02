package widgets

import (
	"fmt"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type BodyStyle struct {
}

func BodyInfo() BodyStyle {
	return BodyStyle{}
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

			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx C) D {

					fmt.Printf("progress bar : %v\n", gtx.Constraints)
					gtx.Constraints.Max.Y = 20

					minX := float32(gtx.Constraints.Min.X)
					minY := float32(gtx.Constraints.Min.Y)

					maxX := float32(gtx.Constraints.Max.X)
					maxY := float32(gtx.Constraints.Max.Y)

					rect := clip.RRect{
						Rect: f32.Rectangle{Min: f32.Point{X: minX, Y: minY},
							Max: f32.Point{X: maxX, Y: maxY}},
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, Yellow, rect)

					return layout.Dimensions{Size: gtx.Constraints.Max}

				}))

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

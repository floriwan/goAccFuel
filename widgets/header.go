package widgets

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

type HeaderInfoStyle struct {
	textColor color.NRGBA
	fuel      int
	length    time.Duration
	laps      int
}

func HeaderInfo(color color.NRGBA,
	fuel int, length time.Duration, laps int) HeaderInfoStyle {
	return HeaderInfoStyle{
		textColor: color,
		fuel:      fuel,
		length:    length,
		laps:      laps,
	}
}

func (h HeaderInfoStyle) Layout(gtx C) D {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
		layout.Rigid(func(gtx C) D {

			paint.ColorOp{Color: h.textColor}.Add(gtx.Ops)
			return layout.Flex{}.Layout(gtx,
				//layout.Rigid(func(gtx C) D {
				//	return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "Session")
				//}),
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
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Session")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2),
									fmtDuration(h.length))
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
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Laps")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), strconv.Itoa(h.laps))
							}),
						)
					})
				}),

				layout.Rigid(func(gtx C) D {

					return layout.Inset{
						Left: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Alignment: layout.End,
						}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize, "Fuel")
							}),
							layout.Rigid(func(gtx C) D {
								return widget.Label{}.Layout(gtx, textShaper, labelFont, labelFontSize.Scale(2), strconv.Itoa(h.fuel)+"l")
							}),
						)
					})
				}),
			)

		}))
	/*
		return layout.Flex{}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return ColorBox(gtx, gtx.Constraints.Min, Blue)
			}),
			layout.Rigid(func(gtx C) D {
				paint.ColorOp{Color: h.textColor}.Add(gtx.Ops)
				return widget.Label{}.Layout(gtx, h.textShaper, h.textFont, h.textFontSize, "Session Info")
			}))*/
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)

	h := d / time.Hour
	d -= h * time.Hour

	m := d / time.Minute
	d -= m * time.Minute

	s := d / time.Second

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func fmtLapTime(d time.Duration) string {
	d = d.Round(time.Second)

	h := d / time.Hour
	d -= h * time.Hour

	m := d / time.Minute
	d -= m * time.Minute

	s := d / time.Second
	d -= s * time.Second

	ms := d / time.Millisecond

	return fmt.Sprintf("%02d:%02d.%03d", m, s, ms)
}

package widgets

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"gioui.org/layout"
	"gioui.org/op/paint"
)

type HeaderInfoStyle struct {
	textColor color.NRGBA
	status    string
	fuel      int
	length    time.Duration
	laps      int
}

func HeaderInfo(color color.NRGBA,
	status string, fuel int, length time.Duration, laps int) HeaderInfoStyle {
	return HeaderInfoStyle{
		status:    status,
		textColor: color,
		fuel:      fuel,
		length:    length,
		laps:      laps,
	}
}

func (h HeaderInfoStyle) Layout(gtx C) D {
	return layout.Flex{}.Layout(gtx,

		layout.Rigid(func(gtx C) D {

			paint.ColorOp{Color: h.textColor}.Add(gtx.Ops)
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Status", h.status)
				}),
			)

		}),
		layout.Flexed(1, func(gtx C) D {
			return ColorBox(gtx, gtx.Constraints.Min, Blue)
		}),
		layout.Rigid(func(gtx C) D {

			paint.ColorOp{Color: h.textColor}.Add(gtx.Ops)
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Session", fmtDuration(h.length))
				}),

				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Laps", strconv.Itoa(h.laps))
				}),

				layout.Rigid(func(gtx C) D {
					return InfoLabel(gtx, "Fuel", strconv.Itoa(h.fuel)+"l")
				}),
			)

		}))

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

package main

import (
	"flag"
	"fmt"
	"goAccFuel/acc"
	"goAccFuel/widgets"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

var backgroundColor = color.NRGBA{R: 18, G: 18, B: 18, A: 255} // very dark gray
var textColor = color.NRGBA{R: 222, G: 222, B: 222, A: 222}

var accData acc.AccData
var sessionData []acc.AccData

type (
	D = layout.Dimensions
	C = layout.Context
)

func main() {

	accSim := flag.Bool("sim", false, "a bool")
	flag.Parse()

	accChan := make(chan acc.AccData)
	go acc.Update(*accSim, accChan)

	go func() {
		w := app.NewWindow(
			app.Title("Go ACC Fuel"),
			app.Size(unit.Dp(600), unit.Dp(250)),
			//app.MaxSize(unit.Dp(600), unit.Dp(200)),
			app.MinSize(unit.Dp(600), unit.Dp(200)),
		)

		if err := run(w, accChan); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)

	}()

	app.Main()

}

func run(w *app.Window, accChan <-chan acc.AccData) error {

	for {

		// ops are the operations from the UI
		ops := new(op.Ops)

		// read gio and acc data channel
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(ops, e)
				AccLayout(ops, gtx)
				e.Frame(gtx.Ops)
			}
		case a := <-accChan:
			//log.Printf("%+v\n", a)
			accData = a
			//saveSessionData(accData)
			//printSessionData(accData)
			w.Invalidate()
		}
	}

}

func printSessionData(accData acc.AccData) {
	if sessionData[len(sessionData)-1].Status == accData.Status {
		return
	}

	for _, v := range sessionData {
		fmt.Printf("%v;%v;%v;%v\n", v.LapsDone, v.LapTime, v.FuelLevel, v.FuelPerLap)
	}

}

func saveSessionData(accData acc.AccData) {
	//if sessionData[len(sessionData)-1].SessionLaps < accData.SessionLaps {
	//	sessionData = append(sessionData, accData)
	//}
}

func AccLayout(ops *op.Ops, gtx C) {

	paint.FillShape(gtx.Ops, backgroundColor, clip.Rect(image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)).Op())

	var rows []layout.FlexChild

	rows = append(rows, layout.Rigid(func(gtx C) D {
		max := gtx.Constraints.Max
		max.Y = 40
		return widgets.HeaderInfo(
			textColor,
			accData.Status,
			accData.SessionLiter,
			accData.SessionLength,
			accData.SessionLaps).Layout(gtx)
	}))

	rows = append(rows,
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = 10

			rect := clip.RRect{
				Rect: f32.Rectangle{Min: f32.Point{X: 0, Y: 4},
					Max: f32.Point{X: float32(gtx.Constraints.Max.X), Y: float32(6)}},
			}.Op(gtx.Ops)
			paint.FillShape(gtx.Ops, widgets.LightGrey, rect)
			return layout.Dimensions{Size: gtx.Constraints.Max}

		}),
	)

	rows = append(rows, layout.Flexed(1, func(gtx C) D {
		paint.ColorOp{Color: textColor}.Add(gtx.Ops)
		return widgets.BodyInfo(
			textColor,
			accData).Layout(gtx)
	}))

	rows = append(rows, layout.Rigid(func(gtx C) D {
		max := gtx.Constraints.Max
		max.Y = 20
		return widgets.FooterInfo(textColor, accData.CarModel, accData.AccVersion).Layout(gtx)

	}))

	layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
			})
		}),
	)

}

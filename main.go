package main

import (
	"fmt"
	"goAccFuel/widgets"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
)

var backgroundColor = color.NRGBA{R: 18, G: 18, B: 18, A: 255} // very dark gray

var textColor = color.NRGBA{R: 222, G: 222, B: 222, A: 222}

var fontCollection []text.FontFace = gofont.Collection()
var textShaper = text.NewCache(fontCollection)
var LabelFont = fontCollection[0].Font
var LabelFontSize = unit.Px(10)

type (
	D = layout.Dimensions
	C = layout.Context
)

func main() {

	go func() {
		w := app.NewWindow(
			app.Title("Go ACC Fuel"),
			app.Size(unit.Dp(600), unit.Dp(200)),
			app.MaxSize(unit.Dp(600), unit.Dp(200)),
			app.MinSize(unit.Dp(600), unit.Dp(200)),
		)

		if err := run(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)

	}()

	app.Main()

}

func run(w *app.Window) error {

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
		}
	}

}

func AccLayout(ops *op.Ops, gtx C) {

	paint.FillShape(gtx.Ops, backgroundColor, clip.Rect(image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)).Op())

	var rows []layout.FlexChild

	rows = append(rows, layout.Rigid(func(gtx C) D {
		max := gtx.Constraints.Max
		max.Y = 40
		fmt.Printf("green box : %v\n", max)
		//return widgets.ColorBox(gtx, max, widgets.Green)
		return widgets.HeaderInfo(LabelFont, LabelFontSize.Scale(2), textColor, textShaper).Layout(gtx)
	}))

	rows = append(rows, layout.Flexed(1, func(gtx C) D {
		fmt.Printf("red box : %v\n", gtx.Constraints)
		return widgets.ColorBox(gtx, gtx.Constraints.Max, widgets.Red)
	}))

	rows = append(rows, layout.Rigid(func(gtx C) D {
		max := gtx.Constraints.Max
		max.Y = 20
		fmt.Printf("blue box : %v\n", max)

		return widgets.FooterInfo(LabelFont, LabelFontSize, textColor, textShaper).Layout(gtx)

	}))

	layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			fmt.Printf("rigit : %v\n", gtx.Constraints)
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
				fmt.Printf("flex : %v\n", gtx.Constraints)
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
			})
		}),
	)

}

package widgets

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/unit"
)

var fontCollection []text.FontFace = gofont.Collection()
var textShaper = text.NewCache(fontCollection)
var labelFont = fontCollection[0].Font
var labelFontSize = unit.Px(10)
var textColor = color.NRGBA{R: 222, G: 222, B: 222, A: 222}

package simulation

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

const template = `
Black
Red
Green
Yellow
Blue
Magenta
Cyan
White

Bright Black
Bright Red
Bright Green
Bright Yellow
Bright Blue
Bright Magenta
Bright Cyan
Bright White

Background
Foreground

Cursor
`

var colors = []string{
	"Black",
	"Red",
	"Green",
	"Yellow",
	"Blue",
	"Magenta",
	"Cyan",
	"White",
}

func PrintColorPair(codes []string, index int, background string) {
	regular := fmt.Sprintf(" %-7s %s ", colors[index], codes[index])
	bright := fmt.Sprintf(" %-7s %s ", colors[index], codes[index+8])
	printColor(regular, codes[index], background)
	printColor(bright, codes[index+8], background)
	printColor(regular, codes[index], background, color.ReverseVideo)
	fmt.Println()
}

func printColor(label, fg, bg string, attr ...color.Attribute) {
	fc := getRGB(fg)
	bc := getRGB(bg)
	c := color.RGB(fc.r, fc.g, fc.b).AddBgRGB(bc.r, bc.g, bc.b).Add(attr...)
	c.Print(label)
}

type rgb struct {
	r, g, b int
}

func getRGB(code string) rgb {
	r, g, b := tcell.GetColor(code).RGB()
	return rgb{int(r), int(g), int(b)}
}

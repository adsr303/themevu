package simulation

import (
	"fmt"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

const width = 3 * (2*7 + 3) // 3 columns of " %-7s %7s "

func PrintTitle(name, fg, bg, cursor string) {
	nameLength := utf8.RuneCountInString(name)
	offset := (width - nameLength - 2) / 2 // Includes cursor
	paddedName := fmt.Sprintf("%*s%s ", offset, "", name)
	printColor(paddedName, fg, bg)
	printColor("\u2588", cursor, bg, color.BlinkSlow)
	completion := fmt.Sprintf("%*s", width-nameLength-offset-2, "")
	printColor(completion, fg, bg)
	fmt.Println()
}

var colors = []string{
	"Black *",
	"Red ***",
	"Green *",
	"Yellow",
	"Blue **",
	"Magenta",
	"Cyan **",
	"White *",
}

func PrintColorPair(codes []string, index int, background string) {
	regular := fmt.Sprintf(" %-7s %7s ", colors[index], codes[index])
	bright := fmt.Sprintf(" %-7s %7s ", colors[index], codes[index+8])
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

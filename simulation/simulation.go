package simulation

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

func PrintAsTable(codes []string, background string) {
	rows := make([][]string, len(colors))
	for i := range colors {
		regular := fmt.Sprintf(" %-7s %7s ", colors[i], codes[i])
		bright := fmt.Sprintf(" %-7s %7s ", colors[i], codes[i+8])
		rows[i] = []string{regular, bright, regular}
	}
	bg := lipgloss.Color(background)
	t := table.New().
		BorderColumn(false).
		BorderLeft(false).
		BorderRight(false).
		BorderTop(false).
		BorderBottom(false).
		StyleFunc(func(row, col int) lipgloss.Style {
			var offset int
			var reverse bool
			switch col {
			case 0:
				offset = 0
			case 1:
				offset = 8
			case 2:
				offset = 0
				reverse = true
			default:
				log.Fatalf("unexpected column index: %d", col)
			}
			fore := codes[row+offset]
			return lipgloss.NewStyle().
				Foreground(lipgloss.Color(fore)).
				Background(bg).
				Reverse(reverse)
		}).
		Rows(rows...)
	fmt.Println(t)
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

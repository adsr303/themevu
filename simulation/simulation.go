package simulation

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const width = 3 * (2*7 + 3) // 3 columns of " %-7s %7s "
const fullBlock = "\u2588"

func PrintTitle(name, fg, bg, cursor string) {
	title := lipgloss.PlaceHorizontal(width, lipgloss.Center, fmt.Sprintf("%s %s", name, fullBlock))
	s := lipgloss.NewStyle().
		Foreground(lipgloss.Color(fg)).
		Background(lipgloss.Color(bg))
	cs := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cursor)).
		Background(lipgloss.Color(bg))
	parts := strings.Split(title, fullBlock)
	fmt.Print(s.Render(parts[0]))
	fmt.Print(cs.Render(fullBlock))
	fmt.Println(s.Render(parts[1]))
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

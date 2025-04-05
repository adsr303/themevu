package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adsr303/themevu/gogh"
	"github.com/adsr303/themevu/simulation"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	var showCode bool
	var permutate bool
	var goghFile string
	flag.BoolVar(&showCode, "fg", false, "show colored text on default background")
	flag.BoolVar(&permutate, "permutate", false, "generate a color swatch of RGB permutations")
	flag.StringVar(&goghFile, "gogh", "", "display colors from a Gogh theme")
	flag.Parse()
	if goghFile == "" {
		showStdin(showCode)
	} else {
		showTheme(goghFile)
	}
}

func showTheme(goghFile string) {
	b, err := os.ReadFile(goghFile)
	if err != nil {
		log.Fatal(err)
	}
	g, err := gogh.ParseTheme(b)
	if err != nil {
		log.Fatal(err)
	}
	simulation.PrintTitle(g.Name, g.Foreground, g.Background, g.Cursor)
	c := g.NumberedColors()
	simulation.PrintAsTable(c, g.Background)
}

func showStdin(showCode bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		s := lipgloss.NewStyle().
			Foreground(lipgloss.Color(line))
		if !showCode {
			line = fullBlocks
		}
		fmt.Println(s.Render(line))
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

const fullBlocks = "\u2588\u2588\u2588"

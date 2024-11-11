package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/adsr303/themevu/gogh"
	"github.com/adsr303/themevu/simulation"
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

func main() {
	var showCode bool
	var forceColor bool
	var goghFile string
	flag.BoolVar(&showCode, "fg", false, "show colored text on default background")
	flag.BoolVar(&forceColor, "force", false, "force color output")
	flag.StringVar(&goghFile, "gogh", "", "display colors from a Gogh theme")
	flag.Parse()
	if forceColor {
		color.NoColor = false
	}
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
	for i := range 8 {
		simulation.PrintColorPair(c, i, g.Background)
	}
}

func showStdin(showCode bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		r, g, b := tcell.GetColor(line).RGB()
		if !showCode {
			line = fullBlocks
		}
		color.RGB(int(r), int(g), int(b)).Println(line)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

const fullBlocks = "\u2588\u2588\u2588"

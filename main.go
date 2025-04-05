package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adsr303/themevu/colors"
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
		showStdin(showCode, permutate)
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

func showStdin(showCode, permutate bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if permutate {
			codes, err := colors.PermutateRGB(line)
			if err != nil {
				log.Fatal(err)
			}
			for _, c := range codes {
				show(c, showCode)
			}
		} else {
			show(line, showCode)
		}
		fmt.Println()
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func show(code string, showCode bool) {
	s := lipgloss.NewStyle().
		Foreground(lipgloss.Color(code))
	if !showCode {
		code = fullBlocks
	}
	fmt.Print(s.Render(code))
}

const fullBlocks = "\u2588\u2588\u2588"

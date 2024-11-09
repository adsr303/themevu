package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

func main() {
	var showCode bool
	flag.BoolVar(&showCode, "fg", false, "show colored text on default background")
	flag.Parse()
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

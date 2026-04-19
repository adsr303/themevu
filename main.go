package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/adsr303/themevu/colors"
	"github.com/adsr303/themevu/simulation"
	"github.com/adsr303/themevu/themes"
)

func main() {
	var (
		showCode     bool
		permutate    bool
		themeVariant string
		themeFile    string
		themesDir    string
	)
	flag.BoolVar(&showCode, "fg", false, "show colored text on default background")
	flag.BoolVar(&permutate, "permutate", false, "generate a color swatch of RGB permutations")
	flag.StringVar(&themeVariant, "variant", "", "light or dark")
	flag.StringVar(&themeFile, "theme", "", "display colors from a theme file")
	flag.StringVar(&themesDir, "dir", "", "directory containing theme files")
	flag.Parse()

	switch themes.Variant(themeVariant) {
	case themes.Light, themes.Dark, "":
		// valid
	default:
		log.Fatalf("invalid theme variant: %s", themeVariant)
	}

	if themesDir != "" {
		showDir(themesDir, themeVariant)
	} else if themeFile == "" {
		showStdin(showCode, permutate)
	} else {
		showTheme(themeFile)
	}
}

func showTheme(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	if b[0] == '{' {
		t, err := themes.ParseTerminal(b)
		if err != nil {
			log.Fatal(err)
		}
		simulation.PrintTitle(t.Name, t.Foreground, t.Background, t.CursorColor)
		c := t.NumberedColors()
		simulation.PrintAsTable(c, t.Background)
	} else {
		g, err := themes.ParseGogh(b)
		if err != nil {
			log.Fatal(err)
		}
		simulation.PrintTitle(g.Name, g.Foreground, g.Background, g.Cursor)
		c := g.NumberedColors()
		simulation.PrintAsTable(c, g.Background)
	}
}

type variantor interface {
	Variant() (themes.Variant, error)
}

func showDir(dir, variant string) {
	ymlMatches, err := filepath.Glob(filepath.Join(dir, "*.yml"))
	if err != nil {
		log.Fatal(err)
	}
	jsonMatches, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		log.Fatal(err)
	}
	matches := append(ymlMatches, jsonMatches...)
	if len(matches) == 0 {
		log.Printf("no theme files found in %s", dir)
		return
	}
	sort.Strings(matches)
	for _, p := range matches {
		b, err := os.ReadFile(p)
		if err != nil {
			log.Fatal(err)
		}
		var theme variantor
		ext := strings.ToLower(filepath.Ext(p))
		switch ext {
		case ".yml", ".yaml":
			g, err := themes.ParseGogh(b)
			if err != nil {
				log.Printf("skipping %s: %v", p, err)
				continue
			}
			theme = g
		case ".json":
			t, err := themes.ParseTerminal(b)
			if err != nil {
				log.Printf("skipping %s: %v", p, err)
				continue
			}
			theme = t
		default:
			// ignore other files
		}
		v, err := theme.Variant()
		if err != nil {
			log.Printf("skipping %s (variant resolution failed): %v", p, err)
			continue
		}
		if variant == "" || v == themes.Variant(variant) {
			showTheme(p)
			fmt.Println()
		}
	}
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

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
		toGogh       bool
		toTerminal   bool
		themesDir    string
	)
	flag.BoolVar(&showCode, "fg", false, "show colored text on default background")
	flag.BoolVar(&permutate, "permutate", false, "generate a color swatch of RGB permutations")
	flag.StringVar(&themeVariant, "variant", "", "light or dark")
	flag.StringVar(&themeFile, "theme", "", "display colors from a theme file")
	flag.BoolVar(&toGogh, "gogh", false, "convert to Gogh format")
	flag.BoolVar(&toTerminal, "terminal", false, "convert to terminal format")
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
		if toGogh {
			term, err := themes.LoadTheme(themeFile)
			if err != nil {
				log.Fatal(err)
			}
			theme, ok := term.(themes.Terminal)
			if !ok {
				log.Fatalf("theme %s is not a terminal theme", themeFile)
			}
			gogh, err := themes.ConvertToGogh(theme)
			if err != nil {
				log.Fatal(err)
			}
			b, err := gogh.ToYAML()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else if toTerminal {
			gogh, err := themes.LoadTheme(themeFile)
			if err != nil {
				log.Fatal(err)
			}
			theme, ok := gogh.(themes.Gogh)
			if !ok {
				log.Fatalf("theme %s is not a Gogh theme", themeFile)
			}
			term, err := themes.ConvertToTerminal(theme)
			if err != nil {
				log.Fatal(err)
			}
			b, err := term.ToJSON()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		} else {
			showTheme(themeFile)
		}
	}
}

func showTheme(path string) {
	theme, err := themes.LoadTheme(path)
	if err != nil {
		log.Fatal(err)
	}
	switch t := theme.(type) {
	case themes.Terminal:
		simulation.PrintTitle(t.Name, t.Foreground, t.Background, t.CursorColor)
		simulation.PrintAsTable(t.NumberedColors(), t.Background)
	case themes.Gogh:
		simulation.PrintTitle(t.Name, t.Foreground, t.Background, t.Cursor)
		simulation.PrintAsTable(t.NumberedColors(), t.Background)
	default:
		log.Fatalf("unexpected theme type: %T", theme)
	}
}

type varianter interface {
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
		ext := filepath.Ext(p)
		if ext != ".yml" && ext != ".json" {
			continue
		}
		loaded, err := themes.LoadTheme(p)
		if err != nil {
			log.Printf("skipping %s: %v", p, err)
			continue
		}
		theme, ok := loaded.(varianter)
		if !ok {
			log.Fatalf("skipping %s: %T does not implement varianter", p, loaded)
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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var showCode bool
	flag.BoolVar(&showCode, "fg", false, "show colored text on default background")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimPrefix(line, "#")
		var r, g, b string
		switch len(line) {
		case 6:
			r, g, b = line[0:2], line[2:4], line[4:6]
		case 3:
			r, g, b = line[0:1], line[1:2], line[2:3]
		default:
			log.Fatalf("invalid color code: %s", line)
		}
		if !showCode {
			line = fullBlocks
		}
		err := printlnRGB(r, g, b, line)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

const fullBlocks = "\u2588\u2588\u2588"

func printlnRGB(r, g, b, text string) error {
	rgb, err := hexToInt(r, g, b)
	if err != nil {
		return err
	}
	color.RGB(rgb[0], rgb[1], rgb[2]).Println(text)
	return nil
}

func hexToInt(values ...string) ([]int, error) {
	result := make([]int, len(values))
	for i, v := range values {
		ii, err := parseHex(v)
		if err != nil {
			return nil, fmt.Errorf("parsing at position %d: %w", i, err)
		}
		result[i] = ii
	}
	return result, nil
}

func parseHex(value string) (int, error) {
	ii, err := strconv.ParseInt(value, 16, 16)
	if err != nil {
		return 0, fmt.Errorf("parsing hex %s: %w", value, err)
	}
	if len(value) == 1 {
		ii = (ii << 4) | ii
	}
	return int(ii), nil
}

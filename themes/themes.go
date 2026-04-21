// Package themes implements parsing of terminal color theme files in various formats.
package themes

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

var numberedColor = regexp.MustCompile(`^color_[0-9]{2}$`)

type Variant string

const (
	Light Variant = "light"
	Dark  Variant = "dark"
)

func LoadTheme(path string) (any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if b[0] == '{' {
		return ParseTerminal(b)
	}
	return ParseGogh(b)
}

func numberedColors(theme any) ([]string, error) {
	result := make([]string, 16)
	termType := reflect.TypeOf(theme)
	for field := range termType.Fields() {
		tag := field.Tag.Get("theme")
		if numberedColor.MatchString(tag) {
			number, err := strconv.Atoi(tag[6:])
			if err != nil {
				return nil, fmt.Errorf("parsing theme %T: %w", theme, err)
			}
			if number < 1 || number > 16 {
				return nil, fmt.Errorf("color tag out of range: %s", tag)
			}
			result[number-1] = reflect.ValueOf(theme).FieldByName(field.Name).String()
		}
	}
	return result, nil
}

func getVariant(theme any) (Variant, error) {
	termType := reflect.TypeOf(theme)
	var foreground, background string
	for field := range termType.Fields() {
		tag := field.Tag.Get("theme")
		switch tag {
		case "background":
			background = reflect.ValueOf(theme).FieldByName(field.Name).String()
		case "foreground":
			foreground = reflect.ValueOf(theme).FieldByName(field.Name).String()
		}
	}

	if foreground == "" {
		return "", fmt.Errorf("no foreground color in theme %T", theme)
	}
	if background == "" {
		return "", fmt.Errorf("no background color in theme %T", theme)
	}

	bgValue, err := getColorValue(background)
	if err != nil {
		return "", fmt.Errorf("parsing theme %T background: %w", theme, err)
	}

	fgValue, err := getColorValue(foreground)
	if err != nil {
		return "", fmt.Errorf("parsing theme %T foreground: %w", theme, err)
	}

	if bgValue > fgValue {
		return Light, nil
	}
	return Dark, nil
}

func getColorValue(rgb string) (int, error) {
	if len(rgb) != 7 || rgb[0] != '#' {
		return 0, fmt.Errorf("invalid RGB color: %s", rgb)
	}

	val, err := strconv.ParseInt(rgb[1:], 16, 32)
	if err != nil {
		return 0, err
	}

	r := (val >> 16) & 0xFF
	g := (val >> 8) & 0xFF
	b := val & 0xFF

	average := (r + g + b) / 3
	return int(average), nil
}

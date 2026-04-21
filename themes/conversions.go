package themes

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertToGogh(theme Terminal) (Gogh, error) {
	gogh := Gogh{
		Name:       theme.Name,
		Author:     "themevu",
		Background: theme.Background,
		Foreground: theme.Foreground,
		Cursor:     theme.CursorColor,
	}
	err := setNumberedColors(&gogh, theme.NumberedColors())
	if err != nil {
		return gogh, err
	}
	v, err := getVariant(theme)
	gogh.LightnessVariant = string(v)
	return gogh, err
}

func ConvertToTerminal(theme Gogh) (Terminal, error) {
	term := Terminal{
		Name:                theme.Name,
		Foreground:          theme.Foreground,
		Background:          theme.Background,
		SelectionBackground: theme.Foreground, // TODO: no selection background in Gogh, use foreground as a fallback
		CursorColor:         theme.Cursor,
	}
	err := setNumberedColors(&term, theme.NumberedColors())
	if err != nil {
		return term, err
	}
	return term, nil
}

func setNumberedColors(theme any, colors []string) error {
	themeType := reflect.TypeOf(theme).Elem()
	themeValue := reflect.ValueOf(theme).Elem()
	for field := range themeType.Fields() {
		tag := field.Tag.Get("theme")
		if numberedColor.MatchString(tag) {
			number, err := strconv.Atoi(tag[6:])
			if err != nil {
				return fmt.Errorf("parsing theme %T: %w", theme, err)
			}
			if number < 1 || number > 16 {
				return fmt.Errorf("color tag out of range: %s", tag)
			}
			themeValue.FieldByName(field.Name).SetString(colors[number-1])
		}
	}
	return nil
}

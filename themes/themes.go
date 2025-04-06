// Package themes implements parsing of terminal color theme files in various formats.
package themes

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var numberedColor = regexp.MustCompile(`^color_[0-9]{2}$`)

func NumberedColors(theme any) ([]string, error) {
	result := make([]string, 16)
	termType := reflect.TypeOf(theme)
	for i := range termType.NumField() {
		field := termType.Field(i)
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

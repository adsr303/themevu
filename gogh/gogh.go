package gogh

import (
	"log"
	"reflect"
	"regexp"
	"strconv"

	"github.com/goccy/go-yaml"
)

// Type Gogh represents the YAML format from the [Gogh terminal themes collection].
//
// [Gogh terminal themes collection]: https://gogh-co.github.io/Gogh/
type Gogh struct {
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Variant string `yaml:"variant"` // dark or light

	Black   string `yaml:"color_01"` // Black (Host)
	Red     string `yaml:"color_02"` // Red (Syntax string)
	Green   string `yaml:"color_03"` // Green (Command)
	Yellow  string `yaml:"color_04"` // Yellow (Command second)
	Blue    string `yaml:"color_05"` // Blue (Path)
	Magenta string `yaml:"color_06"` // Magenta (Syntax var)
	Cyan    string `yaml:"color_07"` // Cyan (Prompt)
	White   string `yaml:"color_08"` // White

	BrightBlack   string `yaml:"color_09"` // Bright Black
	BrightRed     string `yaml:"color_10"` // Bright Red (Command error)
	BrightGreen   string `yaml:"color_11"` // Bright Green (Exec)
	BrightYellow  string `yaml:"color_12"` // Bright Yellow
	BrightBlue    string `yaml:"color_13"` // Bright Blue (Folder)
	BrightMagenta string `yaml:"color_14"` // Bright Magenta
	BrightCyan    string `yaml:"color_15"` // Bright Cyan
	BrightWhite   string `yaml:"color_16"` // Bright White

	Background string `yaml:"background"` // Background
	Foreground string `yaml:"foreground"` // Foreground (Text)

	Cursor string `yaml:"cursor"` // Cursor
}

func ParseTheme(yml []byte) (Gogh, error) {
	var g Gogh
	if err := yaml.Unmarshal(yml, &g); err != nil {
		return g, err
	}
	return g, nil
}

var numberedColor = regexp.MustCompile(`^color_[0-9]{2}$`)

func (g Gogh) NumberedColors() []string {
	result := make([]string, 16)
	gType := reflect.TypeOf(g)
	for i := range gType.NumField() {
		field := gType.Field(i)
		yamlTag := field.Tag.Get("yaml")
		if numberedColor.MatchString(yamlTag) {
			number, err := strconv.Atoi(yamlTag[6:])
			if err != nil {
				log.Fatal(err)
			}
			result[number-1] = reflect.ValueOf(g).FieldByName(field.Name).String()
		}
	}
	return result
}

package themes

import (
	"encoding/json"
	"log"
	"reflect"
	"regexp"
	"strconv"
)

// Type Terminal represents the JSON format of the Windows Terminal color schemes.
type Terminal struct {
	Name string `json:"name"`

	Foreground          string `json:"foreground"`
	Background          string `json:"background"`
	SelectionBackground string `json:"selectionBackground"`
	CursorColor         string `json:"cursorColor"`

	Black  string `json:"black" terminal:"color_01"`
	Red    string `json:"red" terminal:"color_02"`
	Green  string `json:"green" terminal:"color_03"`
	Yellow string `json:"yellow" terminal:"color_04"`
	Blue   string `json:"blue" terminal:"color_05"`
	Purple string `json:"purple" terminal:"color_06"` // a.k.a. magenta
	Cyan   string `json:"cyan" terminal:"color_07"`
	White  string `json:"white" terminal:"color_08"`

	BrightBlack  string `json:"brightBlack" terminal:"color_09"`
	BrightRed    string `json:"brightRed" terminal:"color_10"`
	BrightGreen  string `json:"brightGreen" terminal:"color_11"`
	BrightYellow string `json:"brightYellow" terminal:"color_12"`
	BrightBlue   string `json:"brightBlue" terminal:"color_13"`
	BrightPurple string `json:"brightPurple" terminal:"color_14"` // a.k.a. bright magenta
	BrightCyan   string `json:"brightCyan" terminal:"color_15"`
	BrightWhite  string `json:"brightWhite" terminal:"color_16"`
}

func ParseTerminal(jsonBytes []byte) (Terminal, error) {
	var t Terminal
	if err := json.Unmarshal(jsonBytes, &t); err != nil {
		return t, err
	}
	return t, nil
}

var numberedTerminalColor = regexp.MustCompile(`^color_[0-9]{2}$`)

func (t Terminal) NumberedColors() []string {
	result := make([]string, 16)
	termType := reflect.TypeOf(t)
	for i := range termType.NumField() {
		field := termType.Field(i)
		tag := field.Tag.Get("terminal")
		if numberedTerminalColor.MatchString(tag) {
			number, err := strconv.Atoi(tag[6:])
			if err != nil {
				log.Fatal(err)
			}
			result[number-1] = reflect.ValueOf(t).FieldByName(field.Name).String()
		}
	}
	return result
}

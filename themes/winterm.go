package themes

import (
	"encoding/json"
)

// Type Terminal represents the JSON format of the Windows Terminal color schemes.
type Terminal struct {
	Name string `json:"name"`

	Foreground          string `json:"foreground"`
	Background          string `json:"background"`
	SelectionBackground string `json:"selectionBackground"`
	CursorColor         string `json:"cursorColor"`

	Black  string `json:"black" theme:"color_01"`
	Red    string `json:"red" theme:"color_02"`
	Green  string `json:"green" theme:"color_03"`
	Yellow string `json:"yellow" theme:"color_04"`
	Blue   string `json:"blue" theme:"color_05"`
	Purple string `json:"purple" theme:"color_06"` // a.k.a. magenta
	Cyan   string `json:"cyan" theme:"color_07"`
	White  string `json:"white" theme:"color_08"`

	BrightBlack  string `json:"brightBlack" theme:"color_09"`
	BrightRed    string `json:"brightRed" theme:"color_10"`
	BrightGreen  string `json:"brightGreen" theme:"color_11"`
	BrightYellow string `json:"brightYellow" theme:"color_12"`
	BrightBlue   string `json:"brightBlue" theme:"color_13"`
	BrightPurple string `json:"brightPurple" theme:"color_14"` // a.k.a. bright magenta
	BrightCyan   string `json:"brightCyan" theme:"color_15"`
	BrightWhite  string `json:"brightWhite" theme:"color_16"`
}

func ParseTerminal(jsonBytes []byte) (Terminal, error) {
	var t Terminal
	if err := json.Unmarshal(jsonBytes, &t); err != nil {
		return t, err
	}
	return t, nil
}

func (t Terminal) NumberedColors() []string {
	result, err := NumberedColors(t)
	if err != nil {
		panic(err) // TODO!
	}
	return result
}

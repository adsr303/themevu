package colors

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/etnz/permute"
)

var hexRgbPattern = regexp.MustCompile(`^#?[[:xdigit:]]{6}$`)

func PermutateRGB(color string) ([]string, error) {
	if !hexRgbPattern.MatchString(color) {
		return nil, fmt.Errorf("color code format: %s", color)
	}
	hex, _ := strings.CutPrefix(color, "#")
	values := []string{hex[0:2], hex[2:4], hex[4:6]}
	result := []string{}
	for _, s := range permute.Permutations(values) {
		result = append(result, strings.Join(s, ""))
	}
	return result, nil
}

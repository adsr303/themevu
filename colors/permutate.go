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
	result := make([]string, 0)
	watch := make(map[string]bool)
	for _, s := range permute.Permutations(values) {
		next := "#" + strings.Join(s, "")
		if !watch[next] {
			result = append(result, next)
			watch[next] = true
		}
	}
	return result, nil
}

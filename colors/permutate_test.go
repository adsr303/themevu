package colors_test

import (
	"slices"
	"testing"

	"github.com/adsr303/themevu/colors"
)

func TestPermutateRGB(t *testing.T) {
	tests := []struct {
		color     string
		want      []string
		wantError bool
	}{
		{"#555", nil, true},
		{"fff", nil, true},
		{"abcdef", []string{"abcdef", "abefcd", "cdabef", "cdefab", "efabcd", "efcdab"}, false},
		{"#abcdef", []string{"abcdef", "abefcd", "cdabef", "cdefab", "efabcd", "efcdab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.color, func(t *testing.T) {
			swatch, err := colors.PermutateRGB(tt.color)
			if tt.wantError && err == nil {
				t.Errorf("%s should result in error", tt.color)
			} else {
				slices.Sort(swatch)
				if !slices.Equal(tt.want, swatch) {
					t.Errorf("wanted %v, got %v", tt.want, swatch)
				}
			}
		})
	}
}

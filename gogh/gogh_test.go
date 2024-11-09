package gogh_test

import (
	_ "embed"
	"testing"

	"github.com/adsr303/themevu/gogh"
)

//go:embed Nanosecond.yml
var nanosecond []byte

func TestParseTheme(t *testing.T) {
	g, err := gogh.ParseTheme(nanosecond)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if g.Name != "Nanosecond" {
		t.Errorf("expected name: Nanosecond, got: %s", g.Name)
	}
}

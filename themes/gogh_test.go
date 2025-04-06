package themes_test

import (
	_ "embed"
	"testing"

	"github.com/adsr303/themevu/themes"
)

//go:embed Nanosecond.yml
var nanosecond []byte

func TestParseGogh(t *testing.T) {
	g, err := themes.ParseGogh(nanosecond)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if g.Name != "Nanosecond" {
		t.Errorf("expected name: Nanosecond, got: %s", g.Name)
	}
}

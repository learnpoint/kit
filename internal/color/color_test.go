package color

import (
	"bytes"
	"os"
	"testing"
)

func TestRed(t *testing.T) {
	if got, want := Red("GET / 404"), "\x1b[31mGET / 404\x1b[0m"; got != want {
		t.Errorf("Red() = %q, want %q", got, want)
	}
}

func TestEnabledNotFile(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	if Enabled(&bytes.Buffer{}) {
		t.Error("Enabled(buffer) = true, want false")
	}
}

func TestEnabledNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	if Enabled(os.Stdout) {
		t.Error("Enabled(os.Stdout) with NO_COLOR = true, want false")
	}
}

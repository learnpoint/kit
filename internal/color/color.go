// Package color decides whether output may use ANSI colors and wraps text in them.
package color

import (
	"io"
	"os"
)

// Enabled reports whether w is a terminal that can show colors.
// It respects NO_COLOR (https://no-color.org). On Windows it turns on
// virtual terminal processing for w, which is needed for ANSI codes.
func Enabled(w io.Writer) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return enable(f)
}

// Red wraps s in the ANSI codes for red text.
func Red(s string) string {
	return "\x1b[31m" + s + "\x1b[0m"
}

//go:build !windows

package color

import "os"

// The standard library has no isatty; a character device is the closest
// check without a dependency.
func enable(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

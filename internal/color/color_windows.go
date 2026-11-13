package color

import (
	"os"
	"syscall"
)

const enableVirtualTerminalProcessing = 0x0004

// syscall has GetConsoleMode but not SetConsoleMode.
var setConsoleMode = syscall.NewLazyDLL("kernel32.dll").NewProc("SetConsoleMode")

func enable(f *os.File) bool {
	h := syscall.Handle(f.Fd())
	var mode uint32
	// Fails for pipes and files, so it doubles as the terminal check.
	if err := syscall.GetConsoleMode(h, &mode); err != nil {
		return false
	}
	if mode&enableVirtualTerminalProcessing != 0 {
		return true
	}
	if err := setConsoleMode.Find(); err != nil {
		return false
	}
	ok, _, _ := setConsoleMode.Call(uintptr(h), uintptr(mode|enableVirtualTerminalProcessing))
	return ok != 0
}

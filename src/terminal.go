//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build darwin dragonfly freebsd linux netbsd openbsd solaris zos

package src

import (
	"errors"
	"os"
	"syscall"
	"time"

	"github.com/charmbracelet/x/term"
	"golang.org/x/sys/unix"
)

func drainPendingTerminalInput() {
	fd := os.Stdin.Fd()
	if !term.IsTerminal(fd) {
		return
	}

	// Bubble Tea/Lip Gloss can leave terminal query responses pending if the
	// process is interrupted before they are read. Drain immediately after the
	// TUI exits so the shell prompt does not receive those bytes.
	time.Sleep(25 * time.Millisecond)

	state, err := term.MakeRaw(fd)
	if err != nil {
		return
	}
	defer term.Restore(fd, state) //nolint:errcheck

	flags, err := unix.FcntlInt(fd, unix.F_GETFL, 0)
	if err != nil {
		return
	}
	if _, err := unix.FcntlInt(fd, unix.F_SETFL, flags|unix.O_NONBLOCK); err != nil {
		return
	}
	defer unix.FcntlInt(fd, unix.F_SETFL, flags) //nolint:errcheck

	buf := make([]byte, 256)
	for {
		_, err := os.Stdin.Read(buf)
		if err == nil {
			continue
		}
		if errors.Is(err, syscall.EAGAIN) || errors.Is(err, syscall.EWOULDBLOCK) {
			return
		}
		return
	}
}

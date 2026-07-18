//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !zos
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!zos

package src

func drainPendingTerminalInput() {}

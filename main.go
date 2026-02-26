/*
Copyright © 2026 Saisathvik94
*/
package main

import (
	"runtime"

	"golang.org/x/sys/windows"

	"github.com/Saisathvik94/codemaxx/cmd"
)


func enableVirtualTerminal() {
	if runtime.GOOS != "windows" {
		return
	}

	stdout := windows.Handle(windows.Stdout)
	var mode uint32

	err := windows.GetConsoleMode(stdout, &mode)
	if err != nil {
		return
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	_ = windows.SetConsoleMode(stdout, mode)
}

func main() {
	enableVirtualTerminal()
	cmd.Execute()
}
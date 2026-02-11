package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	var cmd *exec.Cmd
	// This checks if the user is on Windows or Mac/Linux
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

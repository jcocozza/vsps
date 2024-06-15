package internal

import (
  "fmt"
  "os/exec"
  "runtime"
)

// Open a file in the default editor based on OS
func Open(filepath string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filepath)
	case "darwin":
		cmd = exec.Command("open", filepath)
	case "linux":
		cmd = exec.Command("xdg-open", filepath)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}

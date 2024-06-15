package internal

import (
  "fmt"
  "os/exec"
  "runtime"
  "strings"
)

// Send text to the clipboard
func Copy(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "echo "+text+"| clip")
	case "darwin":
		cmd = exec.Command("pbcopy")
		cmd.Stdin = strings.NewReader(text)
	case "linux":
		// Try xclip first
		cmd = exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err != nil {
			// If xclip fails, try xsel
			cmd = exec.Command("xsel", "--clipboard", "--input")
			cmd.Stdin = strings.NewReader(text)
		}
	default:
        return fmt.Errorf("failed to copy: unsupported platform")
	}

	if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to copy: %w", err)
	}
    return nil
}

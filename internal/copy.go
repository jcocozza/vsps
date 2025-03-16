package internal

import (
	"github.com/f1bonacc1/glippy"
)

func ClearClipboard() error {
	return Copy("")
}

// Send text to the clipboard
func Copy(text string) error {
	return glippy.Set(text)
}


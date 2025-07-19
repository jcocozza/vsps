package internal

import (
	"github.com/atotto/clipboard"
)

func ClearClipboard() error {
	return Copy("")
}

// Send text to the clipboard
func Copy(text string) error {
	return clipboard.WriteAll(text)
}


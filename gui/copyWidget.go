package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

type copyLabel struct {
	*widget.Label
}
func NewCopyLabel(w *widget.Label) *copyLabel {
	return &copyLabel{Label: w}
}
func (cw *copyLabel) DoubleTapped(ev *fyne.PointEvent) {
	internal.Copy(cw.Text)
	// TODO: handle error
}

type copyEntry struct {
	*widget.Entry
}
func NewCopyEntry(w *widget.Entry) *copyEntry {
	return &copyEntry{Entry: w}
}
func (ce *copyEntry) DoubleTapped(ev *fyne.PointEvent) {
	internal.Copy(ce.Text)
	// TODO: handle error
}

type copyPasswordEntry struct {
	*widget.Entry
}
func NewCopyPasswordEntry(w *widget.Entry) *copyPasswordEntry {
	w.Password = true
	return &copyPasswordEntry{Entry: w}
}
func (ce *copyPasswordEntry) DoubleTapped(ev *fyne.PointEvent) {
	internal.Copy(ce.Text)
	// TODO: handle error
}

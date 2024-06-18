package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func showErrorDialog(w fyne.Window, err error) {
	    dialog.ShowError(err, w)
}

func main() {
	app := app.New()
	window := app.NewWindow("vsps")

	tabs := makeTabs(window) 

	window.SetContent(tabs)
	window.ShowAndRun()
}

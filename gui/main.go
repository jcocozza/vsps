package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
  app := app.New()
  window := app.NewWindow("vsps")

  appState := NewAppState()
  //acc := makeAccountsViewer(appState, window)

  tabs := makeTabs(window, appState)

  window.SetContent(tabs)
  window.Resize(fyne.NewSize(400,300))
  window.ShowAndRun()
}

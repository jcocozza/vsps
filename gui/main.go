package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"github.com/jcocozza/vsps/internal"
)

func showErrorDialog(w fyne.Window, err error) {
	dialog.ShowError(err, w)
}

func removeStringFromSlice(strings []string, stringToRemove string) []string {
	result := []string{}
	for _, s := range strings {
		if s != stringToRemove {
			result = append(result, s)
		}
	}
	return result
}

type VspsAppState struct {
	Accounts             internal.Accounts
	BoundAccountNameList binding.ExternalStringList
	IsEncrypted          bool
	Masterpass           string
}

func (v *VspsAppState) RemoveAcct(name string) {
	lst, _ := v.BoundAccountNameList.Get()
    newLst := removeStringFromSlice(lst, name)
	v.BoundAccountNameList.Set(newLst)
}

func main() {
	app := app.New()
	window := app.NewWindow("vsps")

	tabs := makeTabs(window)
	window.SetContent(tabs)
	window.ShowAndRun()
}

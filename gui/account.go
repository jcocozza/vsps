package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func makeAccount(acct *internal.Account, appState *VspsAppState) *fyne.Container {
	name := NewCopyEntry(widget.NewEntry())
	name.Text = acct.Name

	username := NewCopyEntry(widget.NewEntry())
	username.Text = acct.Username

	password := NewCopyPasswordEntry(widget.NewEntry())
	password.Text = acct.Password

	left := container.NewVBox(
			name,
			widget.NewLabel("Username"),
			widget.NewLabel("Password"),
	)
	right := container.NewVBox(
			widget.NewLabel(""),
			username,
			password,
	)
	
	for key, value := range acct.Other {
		left.Add(widget.NewLabel(key))
		right.Add(widget.NewLabel(value))
	}

	b := container.NewHBox(left, right)

	editBtn := widget.NewButton("Edit", nil)
	saveBtn := widget.NewButton("Save", nil)
	saveBtn.Hide()

	editBtn.OnTapped = func () {
		name.Entry.Enable()
		username.Entry.Enable()
		password.Entry.Enable()

		editBtn.Hide()
		saveBtn.Show()
	}
	saveBtn.OnTapped = func () {
		name.Disable()
		username.Disable()
		password.Disable()

		acct.Name = name.Text
		acct.Username = username.Text
		acct.Password = password.Text

		p, _ := internal.GetFilePath(appState.IsEncrypted)
		appState.Accounts.Writer(p, appState.IsEncrypted, appState.Masterpass)

		saveBtn.Hide()
		editBtn.Show()
	}
	buttons := container.NewHBox(
		editBtn,
		saveBtn,
		widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
			internal.Copy(acct.Password)
		}),
		widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
			appState.Accounts.Remove(acct.Name)
			p, _ := internal.GetFilePath(appState.IsEncrypted)
			appState.Accounts.Writer(p, appState.IsEncrypted, appState.Masterpass)
			appState.RemoveAcct(acct.Name)
		}),
	)
	return container.NewVBox(b, buttons)
}


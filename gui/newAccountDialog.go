package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func makeNewFieldDialog(window fyne.Window, account *internal.Account) *dialog.FormDialog {
	name := widget.NewEntry()
	value := widget.NewEntry()
	check := widget.NewCheck("Randomize", func(checked bool) {
		if !checked {
			value.Text = ""
		} else {
			value.Text, _ = internal.GeneratePassword(25, true, true, true)
		}
		value.Refresh()
	})
	items := []*widget.FormItem{
		{Text: "Field Name", Widget: name},
		{Text: "Field Value", Widget: value},
		{Text: "", Widget: check},
	}
	dialogCallback := func(conf bool) {
		if conf {
			err := account.AddOtherField(name.Text, value.Text)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			name.Text = ""
			value.Text = ""
		}
	}
	return dialog.NewForm("Add Field", "Add", "Cancel", items, dialogCallback, window)
}

func makeNewAccountDialog(window fyne.Window, appState *VspsAppState) *dialog.FormDialog {
	acct := &internal.Account{
		Name:     "",
		Username: "",
		Password: "",
		Other:    make(map[string]string),
	}

	name := widget.NewEntry()
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()
	password.Text, _ = internal.GeneratePassword(25, true, true, true)
	check := widget.NewCheck("Custom Password", func(checked bool) {
		if checked {
			password.Text = ""
		} else {
			password.Text, _ = internal.GeneratePassword(25, true, true, true)
		}
		password.Refresh()
	})

	newFieldBtn := widget.NewButton("New Field", func() {
		dial := makeNewFieldDialog(window, acct)
		dial.Show()
	})

	items := []*widget.FormItem{
		{Text: "Name", Widget: name},
		{Text: "Username", Widget: username},
		{Text: "Password", Widget: password},
		{Text: "", Widget: check},
		{Text: "", Widget: newFieldBtn},
	}
	dialogCallback := func(conf bool) {
		if conf {
			acct.Name = name.Text
			acct.Username = username.Text
			acct.Password = password.Text

			appState.Accounts.Add(*acct)
			p, _ := internal.GetFilePath(appState.IsEncrypted)
			appState.Accounts.Writer(p, appState.IsEncrypted, appState.Masterpass)
			appState.BoundAccountNameList.Append(acct.Name)
		}
	}
	return dialog.NewForm("New Account", "Save", "Cancel", items, dialogCallback, window)
}

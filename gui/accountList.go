package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func createAccountContainer(account *internal.Account, accounts internal.Accounts, isEncrypted bool, masterpass string) *fyne.Container {
	var accountContainer *fyne.Container
	name := widget.NewEntry()
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()

	name.Text = account.Name
	username.Text = account.Username
	password.Text = account.Password

	name.Disable()
	username.Disable()
	password.Disable()

	editBtn := widget.NewButton("Edit", nil)
	saveBtn := widget.NewButton("Save", nil)
	saveBtn.Hide()

	editBtn.OnTapped = func() {
		name.Enable()
		username.Enable()
		password.Enable()
		editBtn.Hide()
		saveBtn.Show()
	}
	saveBtn.OnTapped = func() {
		name.Disable()
		username.Disable()
		password.Disable()

		account.Name = name.Text
		account.Username = username.Text
		account.Password = password.Text

		p, _ := internal.GetFilePath(isEncrypted)
		accounts.Writer(p, isEncrypted, masterpass)

		saveBtn.Hide()
		editBtn.Show()
	}

	copyPassBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		internal.Copy(account.Password)
	})
	deleteBtn := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		accounts.Remove(account.Name)
		p, _ := internal.GetFilePath(isEncrypted)
		accounts.Writer(p, isEncrypted, masterpass)
		accountContainer.Hide()
	})
	accountContainer = container.NewGridWithColumns(6, name, username, password, editBtn, saveBtn, copyPassBtn, deleteBtn)
	return accountContainer
}

func createAcctList(accounts internal.Accounts, isEncrypted bool, masterpass string) *fyne.Container {
	var items []fyne.CanvasObject
	labeler := container.NewGridWithColumns(5, widget.NewLabel("Name"), widget.NewLabel("Username"), widget.NewLabel("Password"))
	items = append(items, labeler)
	for _, account := range accounts {
		items = append(items, createAccountContainer(account, accounts, isEncrypted, masterpass))
	}
	return container.NewVBox(items...)
}

func createNewAcctForm(accounts internal.Accounts, accountList *fyne.Container, isEncrypted bool, masterpass string) *fyne.Container {
	name := widget.NewEntry()
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()

	password.Text, _ = internal.GeneratePassword(25, true, true, true)

	check := widget.NewCheck("Custom Password", func(value bool) {
		if value {
			password.Text = ""
		} else {
			password.Text, _ = internal.GeneratePassword(25, true, true, true)
		}
		password.Refresh()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name},
			{Text: "Username", Widget: username},
			{Text: "Password", Widget: password},
			{Text: "", Widget: check},
		},
	}
	form.Hide()

	form.OnSubmit = func() {
		acct := internal.Account{Name: name.Text, Username: username.Text, Password: password.Text}
		accounts.Add(acct)
		p, _ := internal.GetFilePath(isEncrypted)
		accounts.Writer(p, isEncrypted, masterpass)
		form.Hide()
		accountList.Add(createAccountContainer(&acct, accounts, isEncrypted, masterpass))
	}
	form.OnCancel = func() { form.Hide() }

	createNew := widget.NewButton("Create New Account", func() {
		form.Show()
	})

	return container.NewVBox(createNew, form)
}

func CreateAccountsContainer(accounts internal.Accounts, isEncrypted bool, masterpass string) *fyne.Container {
	accts := createAcctList(accounts, isEncrypted, masterpass)

	openBtn := widget.NewButton("open", func() {
		p, _ := internal.GetFilePath(false)
		internal.Open(p)
	})
	if !isEncrypted { // provide option to open password file if it is the non-encrypted file
		return container.NewVBox(accts, createNewAcctForm(accounts, accts, isEncrypted, masterpass), openBtn)	
	}
	return container.NewVBox(accts, createNewAcctForm(accounts, accts, isEncrypted, masterpass))	
}


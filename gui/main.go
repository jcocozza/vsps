package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func createAccountContainer(account *internal.Account, accounts internal.Accounts) *fyne.Container {
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

		p, _ := internal.GetFilePath()
		accounts.Write(p)

		saveBtn.Hide()
		editBtn.Show()
	}
	copyPassBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		internal.Copy(account.Password)
	})
	deleteBtn := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		accounts.Remove(account.Name)
		p, _ := internal.GetFilePath()
		accounts.Write(p)
		accountContainer.Hide()
	})
	accountContainer = container.NewGridWithColumns(6, name, username, password, editBtn, saveBtn, copyPassBtn, deleteBtn)
	return accountContainer
}

func createAcctList(accounts internal.Accounts) *fyne.Container {
	var items []fyne.CanvasObject
	labeler := container.NewGridWithColumns(5, widget.NewLabel("Name"), widget.NewLabel("Username"), widget.NewLabel("Password"))
	items = append(items, labeler)
	for _, account := range accounts {
		items = append(items, createAccountContainer(account, accounts))
	}

	// Create a vertical box to hold all the accounts
	return container.NewVBox(items...)
}

func createNew(accounts internal.Accounts, accountList *fyne.Container) *fyne.Container {
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
		OnSubmit: func() {
			accounts.Add(internal.Account{Name: name.Text, Username: username.Text, Password: password.Text})
			p, _ := internal.GetFilePath()
			accounts.Write(p)
		},
	}
	form.Hide()

	form.OnSubmit = func() {
		acct := internal.Account{Name: name.Text, Username: username.Text, Password: password.Text}
		accounts.Add(acct)
		p, _ := internal.GetFilePath()
		accounts.Write(p)
		form.Hide()

		accountList.Add(createAccountContainer(&acct, accounts))
	}
	form.OnCancel = func() { form.Hide() }

	createNew := widget.NewButton("Create New Account", func() {
		form.Show()
	})


	return container.NewVBox(createNew, form)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("vsps")
	p, _ := internal.GetFilePath()
	accounts, _ := internal.LoadAccounts(p)

	acc := createAcctList(accounts)

	lbl := widget.NewLabel("vsps account manager")
	// Set the content of the window
	myWindow.SetContent(container.NewVBox(lbl, acc, createNew(accounts, acc)))
	// Show the window
	myWindow.ShowAndRun()
}

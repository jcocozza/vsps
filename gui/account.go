package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func makeAccountViewer(acct *internal.Account, appState *VspsAppState, window fyne.Window) *fyne.Container {
  acctContainer := container.NewVBox()

  acctName := widget.NewLabel(acct.Name)
  name := NewCopyLabel(acctName)
  acctContainer.Add(name)

  usernameLabel := widget.NewLabel("Username")
  username := NewCopyEntry(widget.NewEntry())
  username.Text = acct.Username
  ucont := container.NewBorder(nil, nil, usernameLabel, nil, username) 
  acctContainer.Add(ucont)

  passwordLabel := widget.NewLabel("Password")
  password := NewCopyPasswordEntry(widget.NewEntry())
  password.Text = acct.Password
  upass := container.NewBorder(nil, nil, passwordLabel, nil, password) 
  acctContainer.Add(upass)

  for key, value := range acct.Other {
    l := NewCopyEntry(widget.NewEntry())
    l.Text = value
    v := container.NewBorder(nil, nil, widget.NewLabel(key), nil, l) 
    acctContainer.Add(v)
  }

  copyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
    internal.Copy(acct.Password)
  })
  acctContainer.Add(copyBtn)

  editBtn := widget.NewButton("Edit", func() {
    dial := makeEditAccountDialog(window, acct, appState)  
    dial.Show()
    acctContainer.Hide()
  })
  acctContainer.Add(editBtn)

  deleteBtn := widget.NewButton("Delete", func() {
    appState.RemoveAccount(acct.Name) 
    acctContainer.Hide()
  })

  acctContainer.Add(deleteBtn)

  return acctContainer 
}

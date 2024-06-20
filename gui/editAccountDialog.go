package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func makeEditAccountDialog(window fyne.Window, acct *internal.Account, appState *VspsAppState) *dialog.FormDialog {
  name := widget.NewEntry()
  name.Text = acct.Name
  username := widget.NewEntry()
  username.Text = acct.Username
  password := widget.NewPasswordEntry()
  password.Text = acct.Password

  newPass := widget.NewButton("Generate New Password", func() {
    password.Text, _ = internal.GeneratePassword(25, true, true, true)
    password.Refresh()
  })

  newOther := widget.NewButton("Add Field", func() {
    dial := makeNewFieldDialog(window, acct)
    dial.Show()
  })

  items := []*widget.FormItem{
    {Text: "Name", Widget: name},
    {Text: "Username", Widget: username},
    {Text: "Password", Widget: password},
    {Text: "", Widget: newPass},
    {Text: "", Widget: newOther},
  }

  markedForDelete := []string{}
  dialogCallback := func(conf bool) {
    if conf {
      oldName := acct.Name

      acct.Name = name.Text
      acct.Username = username.Text
      acct.Password = password.Text

      // the rest of the form elements will be Other data for the account
      // they will all be hbox containers with the first elt being an entry widget 
      // WARNING: THIS IS CURSED 
      // I need to figure out a better way to do this
      for _, elm := range items[5:] {
        // if there is an entry, then it has not been deleted (see delete func below)
        if _, ok := elm.Widget.(*fyne.Container).Objects[0].(*widget.Entry); ok {
          acct.UpdateOtherField(elm.Text, elm.Widget.(*fyne.Container).Objects[0].(*widget.Entry).Text)
        }
      }
      for _, field := range markedForDelete {
        acct.DeleteOtherField(field)
      }
      appState.UpdateAccount(oldName, acct)
    }
  }
  
  for field, value := range acct.Other {
    entry := widget.NewEntry()
    entry.Text = value
  
    cont := container.NewHBox()
    cont.Add(entry)
    cont.Add(widget.NewButtonWithIcon("", theme.DeleteIcon(), func() { // DELETE FUNC HERE
      markedForDelete = append(markedForDelete, field) 
      cont.RemoveAll()
      cont.Add(widget.NewLabel("will be deleted on save"))
    }))

    i := widget.NewFormItem(field, cont)
    items = append(items, i)
  }

  return dialog.NewForm("Edit", "Save", "Cancel", items, dialogCallback, window)
}

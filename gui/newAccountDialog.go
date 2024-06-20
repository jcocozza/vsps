package main

import (
  "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

func makeNewFieldDialog(window fyne.Window, account *internal.Account) *dialog.FormDialog {
  name := widget.NewEntry()
  value := widget.NewEntry()
  rand := widget.NewButton("Randomize", func() {
    value.Text, _ = internal.GeneratePassword(25, true, true, true)
    value.Refresh()
  })

  items := []*widget.FormItem{
    {Text: "Field Name", Widget: name},
    {Text: "Field Value", Widget: value},
    {Text: "", Widget: rand},
  }

  dialogCallback := func(conf bool) {
    if conf {
      err := account.AddOtherField(name.Text, value.Text)
      if err != nil {
        dialog.ShowError(err, window)
        return
      }
      dialog.ShowInformation("FYI", fmt.Sprintf("Field %s has been added to %s", name.Text, account.Name), window)
      name.Text = ""
      value.Text = ""
    }
  }
  return dialog.NewForm("Add Field", "Add", "Cancel", items, dialogCallback, window)
}

func makeNewAccountDialog(window fyne.Window, appState *VspsAppState) *dialog.FormDialog {
  newAcct := &internal.Account{
    Name: "",
    Username: "",
    Password: "",
    Other: make(map[string]string),
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
 
  newOther := widget.NewButton("Add Field", func() {
    dial := makeNewFieldDialog(window, newAcct)
    dial.Show()
  })
  
  items := []*widget.FormItem{
    {Text: "Name", Widget: name},
    {Text: "Username", Widget: username},
    {Text: "Password", Widget: password},
    {Text: "", Widget: check},
    {Text: "", Widget: newOther},
  }

  dialogCallback := func(conf bool) {
    if conf {
      newAcct.Name = name.Text
      newAcct.Username = username.Text
      newAcct.Password = password.Text
      
      appState.AddAccount(*newAcct)
    }
  }
  return dialog.NewForm("New Account", "Save", "Cancel", items, dialogCallback, window)
}

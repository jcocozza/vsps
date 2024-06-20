package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
  accountsTitle = "accounts"
  encryptedAccountsTitle = "encrypted accounts"
)

func makeMasterpassDialog(window fyne.Window, appState *VspsAppState, encryptTab *container.TabItem) *dialog.FormDialog {
  masterpassEntry := widget.NewPasswordEntry()
  items := []*widget.FormItem{
    {Text: "Master Password", Widget: masterpassEntry},
  }

  dialogCallback := func(conf bool) {
    masterpass := masterpassEntry.Text
    if conf {
      appState.IsEcrypted = true
      appState.Masterpass = masterpass

      appState.LoadAccounts()

      encryptTab.Content = makeAccountsViewer(appState, window)
    }
  }
  return dialog.NewForm("Master Password", "OK", "Cancel", items, dialogCallback, window)
}

func makeTabs(window fyne.Window, appState *VspsAppState) *container.AppTabs {
  c := makeAccountsViewer(appState, window)
  t1 := container.NewTabItemWithIcon(accountsTitle, theme.HomeIcon(), c)
  t2 := container.NewTabItemWithIcon(encryptedAccountsTitle, theme.WarningIcon(), widget.NewLabel("awaiting password..."))
  
  tabs := container.NewAppTabs(t1, t2)

  tabs.OnSelected = func(c *container.TabItem) {
    if tabs.Selected().Text == encryptedAccountsTitle {
      dial := makeMasterpassDialog(window, appState, t2)    
      dial.Show()
    } else {
      tabs.Select(t1)
    }
  }
  return tabs
}

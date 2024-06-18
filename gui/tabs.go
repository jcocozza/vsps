package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcocozza/vsps/internal"
)

const (
	accountsTitle = "accounts"
	encryptedAccountsTitle = "encrypted accounts"
)

// create tabs
func makeTabs(window fyne.Window) *container.AppTabs {
	p, _ := internal.GetFilePath(false)
	accounts, _ := internal.AccountLoader(p,false, "")
	c := CreateAccountsContainer(accounts, false, "")

	t1 := container.NewTabItemWithIcon(accountsTitle, theme.HomeIcon(), c)
	t2 := container.NewTabItemWithIcon(encryptedAccountsTitle, theme.WarningIcon(), widget.NewLabel("awaiting password..."))

	tabs := container.NewAppTabs(t1, t2)

	tabs.OnSelected = func (c *container.TabItem) {
		if tabs.Selected().Text == encryptedAccountsTitle {
			masterPassEntry := widget.NewPasswordEntry()
			items := []*widget.FormItem{
				{Text: "Master Password", Widget: masterPassEntry},
			}
			dialogCallback := func(conf bool) {
				// do logic here	
				masterpass := masterPassEntry.Text
				
				if conf {
					p, _ := internal.GetFilePath(true)
					accountsEncrypted, _ := internal.AccountLoader(p, true, masterpass)
					cEncrypted := CreateAccountsContainer(accountsEncrypted, true, masterpass) 
					t2.Content = cEncrypted	
				} else {
					tabs.Select(t1)
				}

			}
			dialog.ShowForm("Master password", "OK", "Cancel", items, dialogCallback, window)
		}
	}
	return tabs
}

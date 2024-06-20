package main

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/jcocozza/vsps/internal"
)

type VspsAppState struct {
	Accounts   internal.Accounts
    accountList []string
    BoundAccountList binding.ExternalStringList
	IsEcrypted bool
	Masterpass string
}

func NewAppState() *VspsAppState {
  v := &VspsAppState{
    Accounts: make(internal.Accounts),
    IsEcrypted: false,
    Masterpass: "",
  } 
  v.LoadAccounts()
  return v 
} 

func (vas *VspsAppState) LoadAccounts() {
	p, _ := internal.GetFilePath(vas.IsEcrypted)
	accounts, _ := internal.AccountLoader(p, vas.IsEcrypted, vas.Masterpass)

    vas.Accounts = accounts
    vas.accountList = vas.Accounts.List()
    vas.BoundAccountList = binding.BindStringList(&vas.accountList)  
}

func (vas *VspsAppState) AddAccount(acct internal.Account) {
  vas.Accounts.Add(acct)
  vas.accountList = vas.Accounts.List()
  vas.BoundAccountList.Reload()

  p, _ := internal.GetFilePath(vas.IsEcrypted)
  vas.Accounts.Writer(p, vas.IsEcrypted, vas.Masterpass)
}

func (vas *VspsAppState) RemoveAccount(acctName string) {
  vas.Accounts.Remove(acctName)
  vas.accountList = vas.Accounts.List()
  vas.BoundAccountList.Reload()
  
  p, _ := internal.GetFilePath(vas.IsEcrypted)
  vas.Accounts.Writer(p, vas.IsEcrypted, vas.Masterpass)
}

func (vas *VspsAppState) UpdateAccount(acctNameToUpdate string, acct *internal.Account) {
  vas.Accounts.UpdateAccount(acctNameToUpdate, acct)
  vas.accountList = vas.Accounts.List()
  vas.BoundAccountList.Reload()

  p, _ := internal.GetFilePath(vas.IsEcrypted)
  vas.Accounts.Writer(p, vas.IsEcrypted, vas.Masterpass)
}

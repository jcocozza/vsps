package cmd

import (
	"fmt"
	"os"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

const accountsFile string = ".vsps.yaml"

var accountsFilePath string

var rootCmd = &cobra.Command{
  Use:   "vsps",
  Short: "vsps is your Very Simple Password Manager",
  Long: "vsps in your Very Simple Password Manager. It's just a yaml file (edit it directly if you like!) with some extra fluff built on top of it.",
  Run: func(cmd *cobra.Command, args []string) {
    accounts, err := internal.LoadAccounts(accountsFilePath)
    if err != nil {
      fmt.Println(err.Error())
      return
    }
    
    if listFlag {
      for _, acct := range accounts {
        acct.Print()
      }
    }

    if showAccount != "" {
      acct, err := accounts.Get(showAccount)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
      acct.Print()
    }
  },
}

func initConfig() {
  home, err := os.UserHomeDir()
  cobra.CheckErr(err)

  accountsFilePath = home + "/" + accountsFile 
}


var listFlag bool
var showAccount string
func init() {
  cobra.OnInitialize(initConfig)


  rootCmd.Flags().BoolVarP(&listFlag, "list","l", false, "list all accounts")
  rootCmd.Flags().StringVarP(&showAccount, "show-account", "s", "", "show an account")
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

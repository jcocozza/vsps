package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var newAccount = &cobra.Command{
    Use: "new",
    Short: "create a new account",
    Run: func(cmd *cobra.Command, args []string) {
        var acctNameInput string
        var usernameInput string
        var passwordInput  string

        fmt.Print("Account Name:")
        fmt.Scanln(&acctNameInput)

        fmt.Print("Username:")
        fmt.Scanln(&usernameInput)

        fmt.Print("Password:")
        fmt.Scanln(&passwordInput)

        account := internal.Account{
            Name: acctNameInput,
            Username: usernameInput,
            Password: passwordInput,
        }
    
        err := account.Write(accountsFilePath)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
    }, 
}

func init() {
    rootCmd.AddCommand(newAccount)
}

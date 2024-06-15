package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var updateUsername bool
var updatePassword bool
var updateAccountName bool

var updateAccount = &cobra.Command{
    Use: "update [account name]",
    Short: "update an account",
    Args: cobra.ExactArgs(1),
    ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
        accounts, err := internal.LoadAccounts(accountsFilePath)    
        fmt.Println(accounts)
        if err != nil {
            return nil, cobra.ShellCompDirectiveError
        }

        acctNames := []string{}
        for acctName := range accounts {
            acctNames = append(acctNames, acctName)
        }
        return acctNames, cobra.ShellCompDirectiveNoFileComp
    }, 
    Run: func(cmd *cobra.Command, args []string) {
        accounts, err := internal.LoadAccounts(accountsFilePath)
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        accountName := args[0] 

        acct, err := accounts.Get(accountName)
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        var updateAccountNameInput string
        if updateAccountName {
            fmt.Print("Enter New Account Name:")
            fmt.Scanln(&updateAccountNameInput)
            acct.Name = updateAccountNameInput
        }
        var updateUsernameInput string
        if updateUsername {
            fmt.Print("Enter New Username:")
            fmt.Scanln(&updateUsernameInput)
            acct.Username = updateUsernameInput
        }
        var updatePasswordInput string
        if updatePassword {
            fmt.Print("Enter New Password:")
            fmt.Scanln(&updatePasswordInput)
            acct.Password = updatePasswordInput
        }
        err0 := accounts.Write(accountsFilePath)
        if err0 != nil {
            fmt.Println(err0.Error())
            return
        }
    },
}

func init() {
    updateAccount.Flags().BoolVarP(&updateUsername, "update-username","u", false, "update the account username")
    updateAccount.Flags().BoolVarP(&updatePassword, "update-password", "p", false, "update the account password")
    updateAccount.Flags().BoolVarP(&updateAccountName, "update-name","a", false, "update an account name")

    rootCmd.AddCommand(updateAccount)
}


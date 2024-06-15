package cmd

import (
    "fmt"

    "github.com/jcocozza/vsps/internal"
    "github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
    Use: "delete [account name]",
    Short: "delete an account",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        accounts, err := internal.LoadAccounts(accountsFilePath)
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        acctName := args[0]
        err = accounts.Remove(acctName)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        err = accounts.Write(accountsFilePath)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
    },
}

func init() {
    rootCmd.AddCommand(deleteCommand)
}

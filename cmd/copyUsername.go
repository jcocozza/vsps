package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var copyUsernameCommand = &cobra.Command{
	Use:               "uc [account name]",
	Short:             "copy the username of the account to your clipboard",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: ValidAccountNames,
	Run: func(cmd *cobra.Command, args []string) {
		accounts, err := internal.AccountLoader(accountsFilePath, encrypted, masterpassword)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		acctName := args[0]
		acct, err := accounts.Get(acctName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = acct.CopyUsername()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("successfully copied username for %s.", acct.Name)
	},
}

func init() {
	rootCmd.AddCommand(copyUsernameCommand)
}

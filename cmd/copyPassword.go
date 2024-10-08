package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var copyPasswordCommand = &cobra.Command{
	Use:               "pc [account name]",
	Short:             "copy the password of the account to your clipboard",
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

		err = acct.CopyPassword()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("successfully copied password for %s.", acct.Name)
	},
}

func init() {
	rootCmd.AddCommand(copyPasswordCommand)
}

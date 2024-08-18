package cmd

import (
	"fmt"
	"time"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

const (
	wait int = 5
)

var copy = &cobra.Command{
	Use: "copy [account name]",
	Short: fmt.Sprintf("copy the account username, then %d seconds later, copy password", wait),
	Args: cobra.ExactArgs(1),
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
		fmt.Printf("copied username for %s.\n", acctName)
		fmt.Print("will copy password in: ")
		for i := 0; i < wait; i ++ {
			fmt.Printf("%d...", wait - i)
			time.Sleep(1 * time.Second)
			if i + 1 == wait {
				fmt.Println("")
			}
		}
		err = acct.CopyPassword()
		fmt.Printf("copied password for %s.\n", acctName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(copy)
}

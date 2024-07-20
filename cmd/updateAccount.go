package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var updateUsername bool
var updatePassword bool
var updateAccountName bool
var fields []string
var addFields bool

var updateAccount = &cobra.Command{
	Use:               "update [account name]",
	Short:             "update an account",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: ValidAccountNames,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		accounts, err := internal.AccountLoader(accountsFilePath, encrypted, masterpassword)
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

		if updateAccountName {
			fmt.Print("Enter New Account Name: ")
			updateAccountNameInput, _ := readInput(reader)
			acct.Name = updateAccountNameInput
		}
		if updateUsername {
			fmt.Print("Enter New Username: ")
			updateUsernameInput, _ := readInput(reader)
			acct.Username = updateUsernameInput
		}
		if updatePassword {
			fmt.Print("Enter New Password: ")
			updatePasswordInput, _ := readInput(reader)
			acct.Password = updatePasswordInput
		}

		if len(fields) != 0 {
			fmt.Println("Updating other fields. Leave blank to delete.")
		}
		for _, field := range fields {
			if acct.HasOtherField(field) {
				fmt.Printf("Enter New value for %s (previously was %s): ", field, acct.Other[field])
				newFieldInput, _ := readInput(reader)
				if newFieldInput == "" {
					acct.DeleteOtherField(field)
				} else {
					acct.UpdateOtherField(field, newFieldInput)
				}
			} else {
				fmt.Printf("field %s not found", field)
			}
		}

		if addFields {
			addField(acct)
		}

		err0 := accounts.Writer(accountsFilePath, encrypted, masterpassword)
		if err0 != nil {
			fmt.Println(err0.Error())
			return
		}

		// if no flags are marked, tell the user they need to include flags
		if !updateUsername && !updatePassword && !updateAccountName && !addFields && len(fields) == 0 {
			fmt.Println("need to include flags for update to specify what to update.")
			fmt.Println("use the following command to get flag updates:\n\tvsps update --help")
		}
	},
}

func init() {
	updateAccount.Flags().BoolVarP(&updateUsername, "update-username", "u", false, "update the account username")
	updateAccount.Flags().BoolVarP(&updatePassword, "update-password", "p", false, "update the account password")
	updateAccount.Flags().BoolVarP(&updateAccountName, "update-name", "a", false, "update an account name")
	updateAccount.Flags().StringSliceVarP(&fields, "update-fields", "f", []string{}, "update the value of extra fields in associated with the account. pass in a list of field names that you want to update.")
	updateAccount.Flags().BoolVarP(&addFields, "add-fields", "i", false, "add additional data fields to the account")

	rootCmd.AddCommand(updateAccount)
}

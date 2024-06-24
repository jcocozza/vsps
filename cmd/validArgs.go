package cmd

import (
	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

func ValidAccountNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	accounts, err := internal.AccountLoader(accountsFilePath, encrypted, masterpassword)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return accounts.List(), cobra.ShellCompDirectiveNoFileComp
}

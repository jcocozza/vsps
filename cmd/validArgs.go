package cmd

import (
    "runtime"
    "os"
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

func getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" && runtime.GOOS == "windows" {
		shell = os.Getenv("ComSpec")
	}
	return shell
}

func generateAndSourceCompletion(shell string) error {
	return nil
}

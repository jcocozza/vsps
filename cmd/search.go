package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

func searchForAccount(accts internal.Accounts) []string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Search for Account: ")
	searchInput, _ := readInput(reader)

	lst := accts.Search(searchInput)
	if len(lst) == 0 {
		fmt.Println("no accounts found")
	}
	return lst
}

var searchAccount = &cobra.Command{
	Use:   "search",
	Short: "search for an account",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		accts, err := internal.AccountLoader(accountsFilePath, encrypted, masterpassword)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		lst := []string{}
		for len(lst) == 0 {
			lst = searchForAccount(accts)
		}
		fmt.Println(strings.Join(lst, ", "))

		fmt.Print("enter acct: ")
		acctName, _ := readInput(reader)
		acct, err := accts.Get(acctName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("----------")
		acct.Print()
	},
}

func init() {
	rootCmd.AddCommand(searchAccount)
}

package cmd

import (
	"fmt"
	"strings"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)


func searchForAccount(accts internal.Accounts) []string {
    var searchInput string
    fmt.Print("Search for Account: ")
    fmt.Scanln(&searchInput)

    lst := accts.Search(searchInput)
    if len(lst) == 0 {
        fmt.Println("no accounts found")
    }
    return lst
}

var searchAccount = &cobra.Command{
    Use: "search",
    Short: "search for an account",
    Run: func(cmd *cobra.Command, args []string) {

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

        var acctName string
        fmt.Print("enter acct: ")
        fmt.Scanln(&acctName)
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

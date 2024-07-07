package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)


func printPasswords(password string, accounts []string) {
    fmt.Println("├── " + password)
    for i, acct := range accounts {
        if i == len(accounts) - 1 {
            fmt.Println("│   └── " + acct)
        } else {
            fmt.Println("│   ├── " + acct)
        }
    }
}

var checkPasswords = &cobra.Command{
    Use: "check-passwords",
    Short: "check to see if accounts share passwords",
    Run: func(cmd *cobra.Command, args []string) {
        accts, err := internal.AccountLoader(accountsFilePath, encrypted, masterpassword)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        passMap := accts.CheckDuplicatePasswords()
        fmt.Println("duplicate password check report:")
        for password, acctList := range passMap {
            //fmt.Printf("%s: %s\n",password, fmt.Sprint(acctList))
            printPasswords(password, acctList)

        }
    },
}

func init() {
    rootCmd.AddCommand(checkPasswords)
}

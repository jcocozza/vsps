package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

// these are global variables that are set on initialization/prerun
// they are NOT flags
// See initConfig() and PersistentPreRun
var accountsFilePath string
var masterpassword string

const version string = "v0.2.2"

var rootCmd = &cobra.Command{
	Use:   "vsps",
	Short: "vsps is your Very Simple Password Manager",
	Long: `vsps is your Very Simple Password Manager.
It's just a simple file (edit it directly if you like!) with some extra fluff built on top of it.`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		accounts, err := internal.AccountLoader(accountsFilePath, encrypted, masterpassword)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if listFlag {
			for _, acct := range accounts {
				acct.Print()
			}
		} else if showAccount != "" {
			acct, err := accounts.Get(showAccount)
			if err != nil {
				similar := accounts.FindSimilar(showAccount)
				if len(similar) != 0 {
					fmt.Printf("Nothing found. Perhaps you meant: %v\n", similar)
				} else {
					fmt.Println(err.Error())
				}
				return
			}
			acct.Print()
		} else {
			cmd.Help()
		}
	},
	// ask for the master password when encrypted flag is included
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		if encrypted {
			_, err := os.Stat(accountsFilePath)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) { // if the encrypted file DNE, then explain the master pass to user
					fmt.Println("About to prompt for master password.")
					fmt.Println("Make sure to keep this somewhere safe and secure. If you lose it, then you will be unable to decrypt your encrypted passwords.")
				} else {
					fmt.Print("an unexpected error occured")
					os.Exit(1)
				}
			}
			fmt.Print("enter master password: ")
			masterpassInput, _ := readInput(reader)
			if masterpassInput == "" {
				fmt.Println("master password must have a non-zero length")
				os.Exit(1)
			}
			masterpassword = masterpassInput
		}
	},
}

func initConfig() {
	// load in file
	path, err := internal.GetFilePath(encrypted)
	cobra.CheckErr(err)
	accountsFilePath = path
}

var listFlag bool
var showAccount string
var encrypted bool

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "list all accounts")
	rootCmd.Flags().StringVarP(&showAccount, "show-account", "s", "", "show an account")
	rootCmd.PersistentFlags().BoolVarP(&encrypted, "encrypted", "e", false, "include this flag to deal with encrypted accounts. requires master password.")
	rootCmd.MarkFlagsMutuallyExclusive("list", "show-account")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var openCommand = &cobra.Command{
	Use:   "open",
	Short: "open the password file in default text editor",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.Open(accountsFilePath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(openCommand)
}

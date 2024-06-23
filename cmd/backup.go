package cmd

import (
    "fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var backup = &cobra.Command{
    Use: "backup",
    Short: "write your accounts file to your downloads folder.",
    Long: "write your accounts file to your downloads folder. including the encrypted flag will write an UNENCRYPTED version of your encryped files to your downloads folder",
    Run: func (cmd *cobra.Command, args []string) {
        err := internal.Backup(masterpassword, encrypted)
        if err != nil {
            fmt.Println("failed to backup file: " + err.Error())
            return
        }
        fmt.Println("successfully wrote accounts file to downloads folder")
    },
}

func init() {
    rootCmd.AddCommand(backup)
}

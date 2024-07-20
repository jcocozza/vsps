package cmd

import (
	"fmt"
	"os"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)

var resetEncryption = &cobra.Command{
	Use:   "reset-encryption",
	Short: "if you forgot your master password, you can reset the encrypted file",
	Long: `if you forgot your master password, you can reset the encrypted file.
this will delete your encrypted file, removing any possibility of retrieving your encrypted passwords.
non-encrypted passwords will be unaffected.
the encrypted flag is not required for this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		encryptedFilePath, err := internal.GetFilePath(true)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		answer, err := askYesNo("are you sure you want to delete encrypted file? This action is irreversible.")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		switch answer {
		case no:
			fmt.Println("aborting...")
			return
		case yes:
			err := os.Remove(encryptedFilePath)
			if err != nil {
				fmt.Println("failed to delete encrypted file: " + err.Error())
				return
			}
			fmt.Println("successfully deleted encrypted file")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(resetEncryption)
}

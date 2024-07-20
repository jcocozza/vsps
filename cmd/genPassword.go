package cmd

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
	"github.com/spf13/cobra"
)


var (
    excludeCapsFlag bool
    excludeDigitsFlag bool
    excludeSpecialCharsFlag bool
    includeExtraSpecialCharsFlag bool
    passwordLengthFlag int
)

var genPassword = &cobra.Command{
    Use: "gen-pass",
    Short: "generate a random password",
    Run: func(cmd *cobra.Command, args []string) {
        password, err := internal.GeneratePassword(passwordLengthFlag, !excludeSpecialCharsFlag, !excludeDigitsFlag, !excludeCapsFlag, includeExtraSpecialCharsFlag)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        fmt.Println(password)
    },
}

func init() {
	genPassword.Flags().IntVarP(&passwordLengthFlag, "pass-length", "l", 25, "set the auto generated password length")
	genPassword.Flags().BoolVarP(&excludeCapsFlag, "exclude-caps", "c", false, "exclude capital letters from autogenerated password")
	genPassword.Flags().BoolVarP(&excludeDigitsFlag, "exclude-digits", "d", false, "exclude digits [0-9] from autogenerated password")
	genPassword.Flags().BoolVarP(&excludeSpecialCharsFlag, "exclude-special-chars", "s", false, "exclude special characters from auto-generated password")
	genPassword.Flags().BoolVarP(&includeExtraSpecialCharsFlag, "include-extra-special-chars", "x", false, "include extra special characters like {} in the auto-generated password")
    rootCmd.AddCommand(genPassword)
}

package cmd

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

const yes = "y"
const no = "n"
const please = "Please enter 'yes' or 'no'"

// return y, n or request yes or no
func askYesNo() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("finished adding extra fields? (yes/no): ")
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("Error reading input: %w", err)
		}

		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		switch answer {
		case "yes", "y":
			return yes, nil
		case "no", "n":
			return no, nil
		default:
			return please, nil
		}
	}
}

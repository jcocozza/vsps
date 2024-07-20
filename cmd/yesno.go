package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const yes = "y"
const no = "n"
const please = "Please enter 'yes' or 'no'"
const yesno = "(yes/no): "

// return y, n or request yes or no
// only include the question. "(yes/no)" will be added to the question.
func askYesNo(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(question + " " + yesno)
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("error reading input: %w", err)
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

package cmd

import "bufio"

// read input from a bufio.Reader till a new line
//
// return the read string (without the '\n' character)
func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input[:len(input)-1], nil
}

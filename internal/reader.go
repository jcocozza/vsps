package internal

import (
	"bufio"
	"os"
)

func ReadAndParse(path string) (Accounts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	strData := string(data)
	return Parser(strData)
}

func GetFileLine(path string, line int) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currline := 0
	for scanner.Scan() {
		if currline == line {
			return scanner.Text()
		}
		currline++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ""
}

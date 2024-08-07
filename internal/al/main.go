package main

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
)

const (
	test = "/Users/josephcocozza/Repositories/vsps/internal/al/test.al"
)

func Parser(input string) internal.Accounts {
	tokens := initTokenizer(input).Tokenize()
	return initParser(tokens).Parse()
}

func main() {
	accts, err := Read(test)
	if err != nil {
		panic(err)
	}
	for _, acct := range accts {
		fmt.Println(acct)
	}
}

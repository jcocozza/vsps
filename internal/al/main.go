package main

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
)

const (
	test = "/Users/josephcocozza/Repositories/vsps/internal/al/test.al"
	test2 = "/Users/josephcocozza/Repositories/vsps/internal/al/vsps.al"
)

func Parser(input string) (internal.Accounts, error) {
	tokens := initTokenizer(input).Tokenize()
	fmt.Println(tokens)
	return initParser(tokens).Parse()
}

func main() {
	accts, err := Read(test2)
	if err != nil {
		panic(err)
	}
	for _, acct := range accts {
		fmt.Println(acct)
	}
	fmt.Println("there are ", len(accts), " accounts")
}

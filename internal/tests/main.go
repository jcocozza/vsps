package main

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
)

const (
	test = "test.al"
	test2 = "vsps.foo"
)


func main() {
	accts, err := internal.Read(test2)
	if err != nil {
		panic(err)
	}
	for _, acct := range accts {
		fmt.Println(acct)
	}
	fmt.Println("there are ", len(accts), " accounts")
}

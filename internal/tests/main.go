package main

import (
	"fmt"

	"github.com/jcocozza/vsps/internal"
)

const (
	test = "test.al"
	test2 = "vsps.al"
)


func main() {
	accts, err := internal.ReadAndParse(test)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, acct := range accts {
		fmt.Println(acct.Password)
	}
	fmt.Println("there are ", len(accts), " accounts")
}

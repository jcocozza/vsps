package main

import (
	"os"

	"github.com/jcocozza/vsps/internal"
)

func Read(path string) (internal.Accounts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	strData := string(data)
	accts := Parser(strData)
	return accts, nil
}

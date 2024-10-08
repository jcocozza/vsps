package internal

import (
	"fmt"
)

const (
	// 4 spaces for nesting
	nestStr       = "    "
	tagNamePrefix = "al"
)

// return an account as a slice of bytes
//
// if an account does not have a name, or if it has no data, throw an error
func Marshal(in Account) ([]byte, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("account name cannot be empty")
	}
	if in.Username == "" && in.Password == "" && len(in.Other) == 0 {
		return nil, fmt.Errorf("account cannot have no data associated with it")
	}
	str := in.Name + ":\n"
	if in.Username != "" {
		str += fmt.Sprintf("%s%s: %s\n", nestStr, "username", in.Username)
	}
	if in.Password != "" {
		str += fmt.Sprintf("%s%s: %s\n", nestStr, "password", in.Password)
	}
	for key, val := range in.Other {
		str += fmt.Sprintf("%s%s: %s\n", nestStr, key, val)
	}
	return []byte(str), nil
}

func Unmarshal(data []byte, accounts *Accounts) error {
	a, err := Parser(string(data))
	if err != nil {
		return err
	}
	accounts = &a
	return nil
}

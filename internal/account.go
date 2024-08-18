package internal

import (
	"fmt"
	"os"
)

type Account struct {
	Name     string
	Username string
	Password string
	Other    map[string]string
}

func (acct Account) HasOtherField(name string) bool {
	if _, ok := acct.Other[name]; ok {
		return true
	}
	return false
}

func (acct Account) AddOtherField(key, value string) error {
	if _, ok := acct.Other[key]; !ok {
		acct.Other[key] = value
	} else {
		return fmt.Errorf(fmt.Sprintf("unable to add account field. %s already has field %s", acct.Name, key))
	}
	return nil
}

func (acct Account) UpdateOtherField(fieldName, newFieldValue string) {
	acct.Other[fieldName] = newFieldValue
}

func (acct Account) DeleteOtherField(name string) {
	delete(acct.Other, name)
}

func (acct Account) CopyPassword() error {
	err := Copy(acct.Password)
	if err != nil {
		return err
	}
	return nil
}

func (acct Account) CopyUsername() error {
	err := Copy(acct.Username)
	if err != nil {
		return err
	}
	return nil
}

// Account writer wrapper
func (acct Account) Writer(filepath string, isEncrypted bool, masterpass string) error {
	if isEncrypted {
		return acct.writeEncrypted(filepath, masterpass)
	} else {
		return acct.write(filepath)
	}
}

// Write an account to the passed file
func (acct Account) write(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	acctAl, err := Marshal(acct)
	if err != nil {
		return err
	}

	_, err0 := file.Write(acctAl)
	if err0 != nil {
		return err0
	}
	return nil
}

// Write an account encrypted
func (acct Account) writeEncrypted(filepath, masterpass string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	acctAl, err := Marshal(acct)
	if err != nil {
		return err
	}

	encryptedAcctAl, err := Encryptor(masterpass, acctAl)
	if err != nil {
		return err
	}
	_, err0 := file.Write(encryptedAcctAl)
	if err0 != nil {
		return err0
	}
	return nil
}

// print the account in proper format for terminal
func (acct Account) Print() {
	fmt.Printf("%s\n", acct.Name)
	fmt.Printf("  username: %s\n", acct.Username)
	fmt.Printf("  password: %s\n", acct.Password)
	for key, value := range acct.Other {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

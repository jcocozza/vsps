package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Account struct {
	Name     string            `yaml:"-"`
	Username string            `yaml:"username"`
	Password string            `yaml:"password"`
	Other    map[string]string `yaml:",inline,omitempty"`
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

func (a Account) MarshalYAML() (interface{}, error) {
	data := make(map[string]interface{})

	// Add Name field with nested fields
	accountData := make(map[string]interface{})
	accountData["username"] = a.Username
	accountData["password"] = a.Password

	for key, value := range a.Other {
		accountData[key] = value
	}

	data[a.Name] = accountData

	return data, nil
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

	acctYaml, err := yaml.Marshal(acct)
	if err != nil {
		return err
	}

	_, err0 := file.Write(acctYaml)
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

	acctYaml, err := yaml.Marshal(acct)
	if err != nil {
		return err
	}

	encryptedAcctYaml, err := Encryptor(masterpass, acctYaml)
	if err != nil {
		return err
	}
	_, err0 := file.Write(encryptedAcctYaml)
	if err0 != nil {
		return err0
	}
	return nil
}

// print the account in proper format for terminal
func (acct Account) Print() {
	fmt.Printf("%s\n", acct.Name)
	fmt.Printf("  Username: %s\n", acct.Username)
	fmt.Printf("  Password: %s\n", acct.Password)
	for key, value := range acct.Other {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

package internal

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Accounts map[string]*Account

// TODO: combine/simplify Load Accounts functions

// Simple Wrapper for account loading
func AccountLoader(filepath string, isEncrypted bool, masterpass string) (Accounts, error) {
	if isEncrypted {
		return loadEncryptedAccounts(filepath, masterpass)
	} else {
		return loadAccounts(filepath)
	}
}

// read in accounts from the passed file path
func loadAccounts(filepath string) (Accounts, error) {
	accounts := make(Accounts)

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err0 := yaml.Unmarshal(bytes, &accounts)
	if err0 != nil {
		return nil, err0
	}
	// need to manually set account name
	for key, account := range accounts {
		account.Name = key
		accounts[key] = account

		if account.Other == nil {
			account.Other = make(map[string]string)
		}
	}
	return accounts, nil
}

// read in encrypted accounts from the passed file path
func loadEncryptedAccounts(filepath, masterpass string) (Accounts, error) {
	accounts := make(Accounts)

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	decryptedBytes, err := Decryptor(masterpass, bytes)
	if err != nil {
		return nil, err
	}

	err0 := yaml.Unmarshal(decryptedBytes, &accounts)
	if err0 != nil {
		return nil, err0
	}
	// need to manually set account name
	for key, account := range accounts {
		account.Name = key
		accounts[key] = account

		if account.Other == nil {
			account.Other = make(map[string]string)
		}
	}
	return accounts, nil
}

// Check if the account exists in the list of accounts
func (accts Accounts) Exists(acctName string) bool {
	if _, ok := accts[acctName]; ok {
		return true
	}
	return false
}

// Get an account
func (accts Accounts) Get(name string) (*Account, error) {
	if acct, ok := accts[name]; ok {
		return acct, nil
	} else {
		return nil, fmt.Errorf(fmt.Sprintf("unable to get account. account %s does not exist", name))
	}
}

// return a list of account names
func (accts *Accounts) List() []string {
	lst := []string{}
	for key := range *accts {
		lst = append(lst, key)
	}
	return lst
}

// Add an account
func (accts Accounts) Add(acct Account) error {
	// if acct is not in accounts list, add it
	// otherwise return an error
	if _, ok := accts[acct.Name]; !ok {
		accts[acct.Name] = &acct
	} else {
		return fmt.Errorf(fmt.Sprintf("unable to add account. account %s already exists", acct.Name))
	}
	return nil
}

// Remove an account
func (accts Accounts) Remove(name string) error {
	// if acct is in account list, remove it
	// otherwise return an error
	if _, ok := accts[name]; ok {
		delete(accts, name)
	} else {
		return fmt.Errorf(fmt.Sprintf("unable to remove account. account %s does not exist", name))
	}
	return nil
}

// Update an account
func (accts Accounts) UpdateAccount(acctToUpdate string, new *Account) error {
	if !accts.Exists(acctToUpdate) {
		return fmt.Errorf(fmt.Sprintf("cannot update %s. This account does not exist", acctToUpdate))
	}
	if acctToUpdate != new.Name {
		// remove the old one if we need to update the key
		delete(accts, acctToUpdate)
	}
	// set to a new one
	accts[new.Name] = new
	return nil
}

// TODO: combine/simplify both write functions

func (accts Accounts) Writer(filepath string, encrypted bool, masterpass string) error {
	if encrypted {
		return accts.writeEncrypted(filepath, masterpass)
	} else {
		return accts.write(filepath)
	}
}

// write the accounts to the passed filepath
func (accts Accounts) write(filepath string) error {
	// I am currently to lazy to implement this properly
	// right now it rewrites the entire file
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	// for each account write it to the file
	for _, value := range accts {
		err := value.write(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}

// write accounts encrypted
func (accts Accounts) writeEncrypted(filepath, masterpass string) error {
	// I am currently to lazy to implement this properly
	// right now it rewrites the entire file
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	// for each account write it to the file
	for _, value := range accts {
		err := value.writeEncrypted(filepath, masterpass)
		if err != nil {
			return err
		}
	}
	return nil
}

// search for an account
// looks for an account based on its name
func (accts Accounts) Search(search string) []string {
	lst := accts.List()

	result := []string{}
	for _, name := range lst {
		if strings.Contains(name, search) {
			result = append(result, name)
		}
	}
	return result
}

func (accts Accounts) FindSimilar(search string) []string {
	// levenshtein distance threshold
	const threshold int = 3

	lst := accts.List()

	result := []string{}
	for _, name := range lst {
		levenDist := levenshteinDistance(search, name)
		if levenDist <= threshold {
			result = append(result, name)
		}
	}
	return result
}

// check for duplicate passwords
//
// return a map of password : acct names
// does not include accounts with no password
func (accts Accounts) CheckDuplicatePasswords() map[string][]string {
	passwordMap := make(map[string][]string)

	for name, acct := range accts {
		if acct.Password == "" {
			continue
		}
		// check if password is in map.
		// if so, add to list
		// otherwise create a new entry in map
		if lst, ok := passwordMap[acct.Password]; ok {
			passwordMap[acct.Password] = append(lst, name)
		} else {
			passwordMap[acct.Password] = []string{name}
		}
	}

	for password, acctList := range passwordMap {
		if len(acctList) == 1 {
			delete(passwordMap, password)
		}
	}

	return passwordMap
}

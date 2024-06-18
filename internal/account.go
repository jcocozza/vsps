package internal

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Account struct {
  Name     string `yaml:"-"`
  Username string `yaml:"username"`
  Password string `yaml:"password"`
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
 
  // necessary to properly write in the yaml format we want
  /* 
  <acct-name>:
    username: <username>
    password: <password>
  */
  type wrappedAcct struct {
    Acct map[string]Account `yaml:",inline"`
  }

  acctWrapper := wrappedAcct{
    Acct: map[string]Account{
      acct.Name: acct,
    },
  }

  acctYaml, err := yaml.Marshal(acctWrapper)
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
 
  // necessary to properly write in the yaml format we want
  /* 
  <acct-name>:
    username: <username>
    password: <password>
  */
  type wrappedAcct struct {
    Acct map[string]Account `yaml:",inline"`
  }

  acctWrapper := wrappedAcct{
    Acct: map[string]Account{
      acct.Name: acct,
    },
  }

  acctYaml, err := yaml.Marshal(acctWrapper)
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
}

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

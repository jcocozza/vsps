package internal

import "os"

const accountsFile string = ".vsps.yaml"
const encryptedAccountsFile string = ".vsps.encrypted.yaml"
// get the path of the vsps yaml
func GetFilePath(encrypted bool) (string, error) {
  home, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }
 
  path := home + "/"
  if encrypted {
    return path + encryptedAccountsFile, nil 
  }
  return path + accountsFile, nil
}

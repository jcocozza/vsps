package internal

import "os"

const accountsFile string = ".vsps.yaml"

// get the path of the vsps yaml
func GetFilePath() (string, error) {
  home, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }
  return home + "/" + accountsFile, nil
}

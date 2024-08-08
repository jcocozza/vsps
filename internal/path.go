package internal

import "os"

const accountsFile string = ".vsps.al"
const encryptedAccountsFile string = ".vsps.encrypted.al"

// get the path of the vsps file
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

// Write the file to the users download folder
func Backup(masterpass string, encrypted bool) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	downloads := home + "/Downloads/vsps-backup.al"
	vspsPath, err := GetFilePath(encrypted)
	if err != nil {
		return err
	}

	accts, err := AccountLoader(vspsPath, encrypted, masterpass)
	if err != nil {
		return err
	}
	// writer will always write an unencrypted file
	err = accts.Writer(downloads, false, "")
	if err != nil {
		return nil
	}
	return nil
}

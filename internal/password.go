package internal

import (
	"crypto/rand"
	"math/big"
)

const (
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

// create the character set based on criteria
func createCharSet(includeSpecialChars, includeDigits, includeCapitals bool) string {
  charSet := lowerLetters 

  if includeSpecialChars {
    charSet += specialChars
  }
  if includeDigits {
    charSet += digits
  }
  if includeCapitals {
    charSet += upperLetters
  }
  return charSet
}

// pick a random element from the passed character set
func pickRandom(charSet string) (byte, error) {
  max := big.NewInt(int64(len(charSet)))
  num, err := rand.Int(rand.Reader, max)
  if err != nil {
    return 0, err
  }
  return charSet[num.Int64()], nil
}

// generate password based on length
func GeneratePassword(length int, specialChars bool, digits bool, capitals bool) (string, error) {
  charSet := createCharSet(specialChars, digits, capitals)

  password := make([]byte, length)

  for i := range password {
      char, err := pickRandom(charSet)
      if err != nil {
        return "",err
      }
      password[i] = char
  }
  return string(password), nil
}

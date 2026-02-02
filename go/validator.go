package diccionario

import (
	"errors"
	"unicode"
)

// ValidateWord checks if a word is valid.
// A valid word is a single string of unbroken alpha characters
// (no numbers, spaces, or special characters).
func ValidateWord(word string) (valid bool, err error) {
	if len(word) == 0 {
		return false, errors.New("word is empty")
	}

	for _, r := range word {
		if !unicode.IsLetter(r) {
			return false, errors.New("word contains non-alpha characters")
		}
	}

	return true, nil
}

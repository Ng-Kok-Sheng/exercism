// Package isogram is a small library to check if a string is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks whether a given string is an isogram.
// Returns true if it is an isogram, else return false
func IsIsogram(word string) bool {
	lowerCasedWord := strings.ToLower(word)

	for _, character := range lowerCasedWord {
		if !unicode.IsLetter(character) {
			continue
		}

		if strings.Count(lowerCasedWord, string(character)) > 1 {
			return false
		}
	}

	return true
}

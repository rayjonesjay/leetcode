package leetcode
/*
link to problem : https://leetcode.com/problems/string-to-integer-atoi/
*/
import (
	"strings"
	"unicode"
)

func MyAtoi(s string) int {
	// remove leading and trailing spaces
	newString := strings.TrimSpace(s)

	// check if the returned string is empty
	if newString == "" {
		return 0
	}

	// assume all the numbers are positive
	isNegative := false

	// check the sign
	if newString[0] == '-' {
		isNegative = true
	} else if newString[0] == '+' {
		isNegative = false
	} else {
		isNegative = false
	}

	// a function to remove the sign if it exist
	newString = removeSign(newString)

	// check if newString is empty after removing sign : instances where the user passed "+" or "-" only remove sign would return an empty string
	if newString == "" {
		return 0
	}

	// result variable to hold the final int equivalence of newString
	var result int = 0
	for _, runeNumber := range newString {
		if !unicode.IsDigit(runeNumber) {
			break
		}
		result = result*10 + int(runeNumber-'0')
	}
	if isNegative {
		return result * -1
	}
	return result
}

// removeSign() removes the sign (- or +) if the string contains it
func removeSign(s string) string {
	// if the first byte is a '+' or '-' remove it from s
	if s[0] == '-' || s[0] == '+' {
		return s[1:]
	}
	return s
}

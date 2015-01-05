// If the value satisfies the condition, return true.
package validation

import (
	"unicode/utf8"
)

func MinLength(str string, min int) bool {
	return utf8.RuneCountInString(str) >= min
}

func MaxLength(str string, max int) bool {
	return utf8.RuneCountInString(str) <= max
}

func Length(str string, min, max int) bool {
	l := utf8.RuneCountInString(str)
	return min <= l && l <= max
}

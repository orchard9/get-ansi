package ansi

import (
	"regexp"
	"unicode/utf8"
)

// ansiRegex matches standard ANSI escape sequences (CSI).
// It matches ESC [ followed by optional parameters and ending with a letter.
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

// Strip removes all ANSI escape sequences from a string.
func Strip(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

// VisibleLen returns the character count of the string excluding ANSI codes.
// Note: This counts runes, not visual width (columns).
func VisibleLen(s string) int {
	stripped := Strip(s)
	return utf8.RuneCountInString(stripped)
}

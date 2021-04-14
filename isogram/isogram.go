// determine if a string is an isogram
package isogram

import "strings"

//determine if there are letters that repeat and return true if not
func IsIsogram(s string) bool {
	s = strings.ToLower(s)

	if s == "" {
		return true
	}

	runeMap := make(map[rune]bool)

	for i, c := range s {
		if s[i] != '-' && s[i] != ' ' {
			if runeMap[c] {
				return false
			}
			runeMap[c] = true
		}
	}

	return true
}

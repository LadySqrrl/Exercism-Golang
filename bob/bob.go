//obnoxious teenager responds to comments/questions
package bob

import (
	"strings"
	"unicode"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Response based on input
func Hey(input string) string {

	upper := IsUpper(input)

	if (strings.HasSuffix(input, "?")) && upper {
		return "Calm down, I know what I'm doing!"
	}

	if !(strings.HasSuffix(input, "?")) && upper {
		return "Whoa, chill out!"
	}

	if (strings.HasSuffix(input, "?")) && !upper {
		return "Sure."
	}

	if input == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

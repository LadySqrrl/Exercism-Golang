//determine if a sentence contains at least 1 instance of each letter
package pangram

import "strings"

//identify each unique character and determine if the full string is a pangram
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	encountered := map[byte]bool{}

	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			encountered[s[i]] = true
		}
	}

	return len(encountered) == 26
}

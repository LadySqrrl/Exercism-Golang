//determine if a string is a valid number using the Luhn algorithm
package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

//complete the Luhn algorith on the string
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	sum := 0

	if len(s) < 2 {
		return false
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		curr := s[i]
		currInt, err := strconv.Atoi(string(curr))
		if err != nil {
			fmt.Println("error converting to int")
		}
		if (len(s)%2 == 0 && i%2 == 0) || (len(s)%2 != 0 && i%2 != 0) {
			currInt *= 2
			if currInt > 9 {
				currInt -= 9
			}
		}
		sum += currInt
	}

	return sum%10 == 0
}

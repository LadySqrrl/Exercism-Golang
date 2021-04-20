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

	if len(s) < 2 {
		return false
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}

	intArr := make([]int, len(s))

	for i := range s {
		curr := s[i]
		currInt, err := strconv.Atoi(string(curr))
		if err != nil {
			fmt.Println("error converting to int")
		}
		intArr[i] = currInt
	}

	if len(intArr)%2 == 0 {
		for i := 0; i < len(intArr); i += 2 {
			currInt := intArr[i]
			currInt *= 2
			if currInt > 9 {
				currInt -= 9
			}
			intArr[i] = currInt
		}
	} else {
		for i := 1; i < len(intArr); i += 2 {
			currInt := intArr[i]
			currInt *= 2
			if currInt > 9 {
				currInt -= 9
			}
			intArr[i] = currInt
		}
	}

	sum := 0

	for i := range intArr {
		sum += intArr[i]
	}

	return sum%10 == 0
}

//Take a string and return its acronym
package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

// return the first letter of each word in a string
func Abbreviate(s string) (acronym string) {

	s = strings.Replace(s, "'", "", -1)

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}

	processedString := reg.ReplaceAllString(s, " ")

	strSplit := strings.Split(strings.Title(processedString), " ")

	for _, word := range strSplit {
		if len(word) > 0 {
			acronym += string(word[0])
		}
	}

	fmt.Println(acronym)

	return acronym
}

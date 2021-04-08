// calculate the scrabble score of a given string
package scrabble

import (
	"strings"
)

var letterScores = map[int][]rune{
	1:  {'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	2:  {'D', 'G'},
	3:  {'B', 'C', 'M', 'P'},
	4:  {'F', 'H', 'V', 'W', 'Y'},
	5:  {'K'},
	8:  {'J', 'X'},
	10: {'Q', 'Z'},
}

func Score(s string) int {
	if s == "" {
		return 0
	}

	score := 0

	s = strings.ToUpper(s)

	var scoringMap = make(map[rune]int)

	for s, letters := range letterScores {
		for _, l := range letters {
			scoringMap[l] = s
		}
	}

	for _, c := range s {
		i, ok := scoringMap[c]
		if ok {
			score += i
		}
	}

	return score
}

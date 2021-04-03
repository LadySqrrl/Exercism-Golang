package raindrops

import "strconv"

func Convert(input int) string {
	rainSound := ""
	matches := 0

	if input%3 == 0 {
		rainSound += "Pling"
		matches++
	}

	if input%5 == 0 {
		rainSound += "Plang"
		matches++
	}

	if input%7 == 0 {
		rainSound += "Plong"
		matches++
	}

	if matches == 0 {
		rainSound = strconv.Itoa(input)
	}

	return rainSound

}

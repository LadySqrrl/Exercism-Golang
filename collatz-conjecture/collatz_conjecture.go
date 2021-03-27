//determine how many steps it takes to get to 1
package collatzconjecture

//determine if input is even or odd
func isEven(input int) bool {
	return input%2 == 0
}

//do the collatz calculations and return how many steps to get to 1
func CollatzConjecture(input int) int {
	count := 0

	if input <= 0 {
		return 0
	}

	for input != 1 {
		if isEven(input) {
			input = input / 2
		} else {
			input = (input * 3) + 1
		}
		count++
	}

	return count
}

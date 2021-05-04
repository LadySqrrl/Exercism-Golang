//Package grains calculates how many grains are on a particular square of a chessboard and the total grains
package grains

import "errors"

//Square returns the number of grains on a given square
func Square(input int) (uint64, error) {
	previousSquare := 1
	currentSquare := 1

	if input < 1 || input > 64 {
		return uint64(0), errors.New("out of bounds")
	}

	for i := 1; i < input; i++ {
		currentSquare = previousSquare * 2
		previousSquare = currentSquare
	}

	return uint64(currentSquare), nil
}

//Total returns the total number of grains on the board
func Total() uint64 {
	previousSquare := 1
	currentSquare := 1
	total := 1

	for i := 1; i < 65; i++ {
		currentSquare = previousSquare * 2
		previousSquare = currentSquare
		total += currentSquare
	}

	return uint64(total)
}

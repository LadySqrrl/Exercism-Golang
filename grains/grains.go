//Package grains calculates how many grains are on a particular square of a chessboard and the total grains
package grains

import "fmt"

//Square returns the number of grains on a given square
func Square(input int) (uint64, error) {

	if input < 1 || input > 64 {
		return 0, fmt.Errorf("bad square %d (must be 1-64)", input)
	}

	return 1 << (input - 1), nil
}

//Total returns the total number of grains on the board
func Total() uint64 {

	return 1<<64 - 1
}

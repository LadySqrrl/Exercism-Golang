//calculate the hamming distance between two DNA strands
package hamming

import "errors"

//if strings are the same length, return number of differences, otherwise reutrn error
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("different lengths")
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}

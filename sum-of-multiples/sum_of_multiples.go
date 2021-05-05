//Package summultiples finds the sum of all the unique multiples of particular numbers up to but not including that number
package summultiples

//SumMultiples returns the sum of all the unique multiples up to the limit
func SumMultiples(limit int, divisors ...int) (sum int) {
	sum = 0
	encountered := map[int]bool{}

	for i := range divisors {
		for j := 1; j < limit; j++ {
			if divisors[i] != 0 && j%divisors[i] == 0 && !encountered[j] {
				sum += j
				encountered[j] = true
			}
		}
	}

	return sum
}

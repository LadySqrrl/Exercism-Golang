//determine the difference between the sum of squares and the square of sum of set of ints
package diffsquares

//get the square of the sum
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return (sum * sum)
}

//get the sum of the squares
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}

	return sum
}

//determine the absolute difference
func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)
}

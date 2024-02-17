package diffsquares

func SquareOfSum(n int) int {
	sumOfSequence := n * (n + 1) / 2
	return sumOfSequence * sumOfSequence
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

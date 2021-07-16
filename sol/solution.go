package sol

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	ways := 0
	prevTwo := 2
	prevOne := 3
	for stairs := 3; stairs <= n; stairs++ {
		ways = prevTwo + prevOne
		prevTwo = prevOne
		prevOne = ways
	}

	return ways
}

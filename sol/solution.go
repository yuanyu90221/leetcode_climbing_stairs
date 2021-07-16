package sol

func climbStairs(n int) int {
	// create space to store result
	ways := make([]int, n)
	// step 1
	ways[0] = 1
	ways[1] = 2

	for stairs := 3; stairs <= n; stairs++ {
		ways[stairs-1] = ways[stairs-2] + ways[stairs-3]
	}
	return ways[n-1]
}

package stairs

func climbStairs(n int) int {

	if n < 3 {
		return n
	}

	n1, n2 := 1, 2

	for i := 2; i < n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

package uniquepaths

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)

	for i := range len(dp) {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := range m {
		for j := range n {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

package uniquepaths2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	rows, cols := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, rows)

	for i := range len(dp) {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = 1

	for i := range rows {
		for j := range cols {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}

		}
	}
	return dp[rows-1][cols-1]
}

package minpathsum

func minPathSum(grid [][]int) int {

	rows, cols := len(grid), len(grid[0])

	dp := make([][]int, rows)
	for i := range rows {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = grid[0][0]

	for row := range rows {
		for col := range cols {
			if row == 0 && col == 0 {
				continue
			} else if row == 0 {
				dp[row][col] += grid[row][col] + dp[row][col-1]
			} else if col == 0 {
				dp[row][col] += grid[row][col] + dp[row-1][col]
			} else {
				dp[row][col] += grid[row][col] + min(dp[row-1][col], dp[row][col-1])
			}
		}
	}

	return dp[rows-1][cols-1]
}

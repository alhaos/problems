package lcs

func longestCommonSubsequence(text1 string, text2 string) int {

	rows := len(text1) + 1
	cols := len(text2) + 1

	dp := make([][]int, rows)
	for i := range rows {
		dp[i] = make([]int, cols)
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if text1[i-1] != text2[j-1] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return dp[rows-1][cols-1]
}

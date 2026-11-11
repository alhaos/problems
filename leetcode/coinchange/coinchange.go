package coinchange

func coinChange(coins []int, amount int) int {

	overflow := amount + 1

	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = overflow
	}

	for _, coin := range coins {
		for currentAmount := coin; currentAmount < amount+1; currentAmount++ {
			dp[currentAmount] = min(dp[currentAmount], dp[currentAmount-coin]+1)
		}
	}

	if dp[amount] < overflow {
		return dp[amount]
	}

	return -1
}

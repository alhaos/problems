package maxprofit

func maxProfit(prices []int) int {

	l := len(prices)

	if l == 0  {
		return 0
	}

	best, minPrice := 0, prices[0]

	for i := 1; i < l; i++ {
		best = max(best, prices[i]-minPrice)
		minPrice = min(prices[i], minPrice)
	}

	return best
}

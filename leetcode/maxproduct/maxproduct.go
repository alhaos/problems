package maxproduct

func maxProduct(nums []int) int {
	best := nums[0]
	maxCur := nums[0]
	minCur := nums[0]
	for i := 1; i < len(nums); i++ {
		curMinCur := min(minCur*nums[i], min(maxCur*nums[i], nums[i]))
		curMaxCur := max(maxCur*nums[i], max(minCur*nums[i], nums[i]))
		minCur, maxCur = curMinCur, curMaxCur
		best = max(best, maxCur)
	}

	return best
}

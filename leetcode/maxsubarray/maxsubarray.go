package maxsubarray

func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		cur = max(cur+nums[i], nums[i])
		best = max(best, cur)
	}

	return best
}

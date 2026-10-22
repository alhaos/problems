package robber

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	sum1 := nums[0]
	sum2 := max(nums[0], nums[1])
	sum3 := 0

	for i := 2; i < len(nums); i++ {
		take := nums[i] + sum1
		skip := sum2
		sum3 = max(take, skip)
		sum1 = skip
		sum2 = sum3
	}

	return sum3
}

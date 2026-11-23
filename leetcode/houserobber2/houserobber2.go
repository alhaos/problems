package houserobber2

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	return max(robLine(nums[1:]), robLine(nums[:len(nums)-1]))
}

func robLine(nums []int) int {
	prev2 := 0
	prev1 := 0

	for _, num := range nums {
		curr := max(prev1, prev2+num)
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

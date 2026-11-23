package houserobber

func rob(nums []int) int {
	prev2 := 0
	prev1 := 0

	for _, num := range nums {
		curr := max(prev1, prev2+num)
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

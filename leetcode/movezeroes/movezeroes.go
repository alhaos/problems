package movezeroes

func moveZeroes(nums []int) {
	var w int
	for i := range len(nums) {
		if nums[i] != 0 {
			nums[i], nums[w] = nums[w], nums[i]
			w++
		}
	}
}

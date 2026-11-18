package removeduplicates

func removeDuplicates(nums []int) int {
	w := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[w] {
			w++
			nums[w] = nums[i]
		}
	}
	return w + 1
}

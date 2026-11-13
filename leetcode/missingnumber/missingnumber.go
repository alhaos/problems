package missingnumber

func missingNumber(nums []int) int {
	x := len(nums)
	for i, n := range nums {
		x ^= i
		x ^= n
	}
	return x
}

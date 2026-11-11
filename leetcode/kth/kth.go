package kth

func findKthLargest(nums []int, k int) int {

	h := From(nums)

	var result int

	for range k {
		result, _ = h.Pop()
	}

	return result

}

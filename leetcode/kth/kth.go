package kth

func findKthLargest(nums []int, k int) int {

	h := New()

	for _, n := range nums {
		h.Push(n)
	}

	var result int

	for range k {
		result, _ = h.Pop()
	}

	return result

}

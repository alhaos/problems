package intersectii

func intersect(nums1 []int, nums2 []int) []int {

	counts := make(map[int]int)

	for _, n := range nums1 {
		counts[n]++
	}

	var result []int

	for _, n := range nums2 {
		if counts[n] > 0 {
			result = append(result, n)
			counts[n]--
		}
	}

	return result
}

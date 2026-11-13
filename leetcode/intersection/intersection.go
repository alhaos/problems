package intersection

func intersection(nums1 []int, nums2 []int) []int {

	hs1 := make(map[int]struct{})

	for _, n := range nums1 {
		hs1[n] = struct{}{}
	}

	var result []int

	for _, n := range nums2 {
		if _, exist := hs1[n]; exist {
			result = append(result, n)
			delete(hs1, n)
		}
	}

	return result
}

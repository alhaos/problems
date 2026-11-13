package intersectionofmultiplearrays

import "slices"

func intersection(nums [][]int) []int {

	m := make(map[int]int)

	for _, num := range nums {
		for _, n := range num {
			m[n]++
		}
	}

	var result []int

	for k, v := range m {
		if v == len(nums) {
			result = append(result, k)
		}
	}

	slices.Sort(result)

	return result
}

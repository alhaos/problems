package permutations

import "slices"

// Ehrlich's swap algorithm (Knuth, TAOCP 7.2.1.2, Algorithm E)
func permute(nums []int) [][]int {

	clone := slices.Clone(nums)
	// results slice
	var res [][]int

	// additional slice b
	b := make([]int, len(clone))

	// fill b slice with indexes
	for i := range len(b) {
		b[i] = i
	}

	// init additional slice c
	c := make([]int, len(clone)+1)

	// init indexes
	var j, k int

	// append base slice first step
	for {
		// append clone current clone to results
		res = append(res, slices.Clone(clone))
		k = 1
		for c[k] == k {
			c[k] = 0
			k++
		}

		if k == len(clone) {
			return res
		}

		c[k]++
		clone[0], clone[b[k]] = clone[b[k]], clone[0]
		j = 1
		k--

		for j < k {
			b[j], b[k] = b[k], b[j]
			j++
			k--
		}
	}
}

func permuteBacktrack(nums []int) [][]int {

	var result [][]int

	var fn func(nums []int, path []int)
	fn = func(nums []int, path []int) {

		if len(nums) == 0 {
			result = append(result, path)
			return
		}

		for i := range nums {
			rest := make([]int, 0)
			rest = append(rest, nums[:i]...)
			rest = append(rest, nums[i+1:]...)
			fn(rest, append(slices.Clone(path), nums[i]))
		}
	}

	fn(nums, nil)

	return result
}

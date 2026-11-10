package subsets

import "slices"

func subsets(nums []int) [][]int {

	res := [][]int{}
	res = append(res, []int{})

	var fn func(position int, path []int)

	fn = func(position int, path []int) {
		for i := position; i < len(nums); i++ {
			next := append(slices.Clone(path), nums[i])
			res = append(res, next)
			fn(i+1, next)
		}
	}

	fn(0, nil)

	return res
}

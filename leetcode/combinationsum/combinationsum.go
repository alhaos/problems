package combinationsum

import (
	"slices"
)

func combinationSum(candidates []int, target int) [][]int {

	cp := slices.Clone(candidates)

	slices.Sort(cp)

	var res [][]int

	var fn func(start int, taken []int, remaining int)

	fn = func(start int, taken []int, remaining int) {
		if remaining < 0 {
			return
		}

		if remaining == 0 {
			res = append(res, taken)
		}

		for i := start; i < len(cp); i++ {
			if cp[i] > remaining {
				return
			}

			fn(i, append(slices.Clone(taken), cp[i]), remaining-cp[i])
		}
	}

	fn(0, nil, target)

	return res
}

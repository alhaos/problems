package permutationsii

import "slices"

func permuteUnique(nums []int) [][]int {

	numsCp := slices.Clone(nums)
	slices.Sort(numsCp)

	var res [][]int
	var path []int
	used := make([]bool, len(numsCp))

	var backtrack func()

	backtrack = func() {

		if len(path) == len(numsCp) {
			res = append(res, slices.Clone(path))
			return
		}

		for i := range numsCp {

			if used[i] || (i > 0 && numsCp[i-1] == numsCp[i] && !used[i-1]) {
				continue
			}

			used[i] = true
			path = append(path, numsCp[i])
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()

	return res
}

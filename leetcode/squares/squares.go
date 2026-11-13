package sortedsquares

import "slices"

func sortedSquares_(nums []int) []int {

	cp := slices.Clone(nums)

	for i := range cp {
		cp[i] = cp[i] * cp[i]
	}

	slices.Sort(cp)

	return cp
}

func sortedSquares(nums []int) []int {

	result := make([]int, len(nums))

	l, r := 0, len(nums)-1

	for pos := len(nums) - 1; pos >= 0; pos-- {
		lSq, rSq := nums[l]*nums[l], nums[r]*nums[r]
		if lSq > rSq {
			result[pos] = lSq
			l++
		} else {
			result[pos] = rSq
			r--
		}
	}

	return result
}

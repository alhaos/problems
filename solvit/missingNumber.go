package solvit

import "math"

func missingNumber(nums []int) int {

	m := map[int]struct{}{}
	nMin := math.MaxInt
	nMax := math.MinInt

	for _, n := range nums {
		m[n] = struct{}{}
		if n < nMin {
			nMin = n
		}
		if n > nMax {
			nMax = n
		}
	}

	for i := nMin; i < nMax; i++ {
		if _, exist := m[i]; !exist {
			return i
		}
	}

	panic("not found")
}

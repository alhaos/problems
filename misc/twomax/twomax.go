package twomax

import (
	"errors"
	"math"
)

// TwoMax returns the two largest elements in nums (max1 >= max2),
// scanning the slice exactly once.
// If len(nums) < 2, ok is false.
func TwoMax(nums []int) (int, int, error) {

	if len(nums) < 2 {
		return 0, 0, errors.New("minimum two elements required")
	}

	max1 := math.MinInt
	max2 := math.MinInt

	for _, n := range nums {
		if n > max1 {
			max2 = max1
			max1 = n
			continue
		}
		if n > max2 {
			max2 = n
		}
	}

	return max1, max2, nil
}

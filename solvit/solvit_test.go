package solvit

import "testing"

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		desc     string
		nums     []int
		expected int
	}{
		{
			desc:     "#1",
			nums:     []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := missingNumber(tC.nums)
			if result != tC.expected {
				t.Errorf("Unexpected results for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

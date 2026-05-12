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

func TestAngleClock(t *testing.T) {
	testCases := []struct {
		desc     string
		hour     int
		minutes  int
		expected float64
	}{
		{
			desc:     "#1",
			hour:     12,
			minutes:  30,
			expected: 165,
		},
		{
			desc:     "#2",
			hour:     3,
			minutes:  30,
			expected: 75,
		},
		{
			desc:     "#3",
			hour:     3,
			minutes:  15,
			expected: 7.5,
		},
		{
			desc:     "#4",
			hour:     1,
			minutes:  57,
			expected: 76.5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := angleClock(tC.hour, tC.minutes)
			if result != tC.expected {
				t.Errorf("Unexpected result for test %s, expected: %g, but got: %g", tC.desc, tC.expected, result)
			}
		})
	}
}

package lis

import "slices"

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}

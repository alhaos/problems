package slidingWindow

func maxSum(nums []int, k int) int {
	sum := 0
	for i := range k {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

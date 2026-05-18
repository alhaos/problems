package leetcode

func threeSum(nums []int) [][]int {

	var results [][]int

	for i := range nums {
		for j := range nums {
			for k := range nums {
				if i == j {
					continue
				}

				if i == k {
					continue
				}

				if j == k {
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					results = append(results, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return results
}

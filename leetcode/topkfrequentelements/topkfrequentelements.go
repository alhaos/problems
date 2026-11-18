package topkfrequentelements

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	buckets := make([][]int, len(nums)+1)

	for k, v := range m {
		buckets[v] = append(buckets[v], k)
	}

	var result []int
	for i := len(buckets) - 1; i >= 0; i-- {
		result = append(result, buckets[i]...)
		if len(result) >= k {
			return result
		}
	}

	return nil
}

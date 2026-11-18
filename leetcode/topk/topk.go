package topk

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	bucket := make([][]int, len(nums)+1)

	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}

	var res []int

	for i := len(bucket) - 1; i >= 0; i-- {
		res = append(res, bucket[i]...)
		if len(res) == k {
			break
		}
	}

	return res
}

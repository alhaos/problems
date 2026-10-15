package insertinterval

func Insert(intervals [][]int, newInterval []int) [][]int {

	var result [][]int

	for i, interval := range intervals {

		if newInterval[1] < interval[0] {
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		}

		if interval[1] < newInterval[0] {
			result = append(result, interval)
			continue
		}

		newInterval[0] = min(newInterval[0], interval[0])
		newInterval[1] = max(newInterval[1], interval[1])
	}

	result = append(result, newInterval)

	return result
}

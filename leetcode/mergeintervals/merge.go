package mergeintervals

import "sort"

func Merge(intervals [][]int) [][]int {

	sort.Slice(
		intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		},
	)

	var result [][]int
	buffer := intervals[0]

	for _, interval := range intervals[1:] {

		if interval[1] < buffer[1] {
			continue
		}

		if buffer[1] >= interval[0] {
			buffer[1] = interval[1]
			continue
		}
		result = append(result, buffer)
		buffer = interval
	}

	result = append(result, buffer)

	return result
}

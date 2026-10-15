package meetingrooms

import (
	"cmp"
	"slices"
)

func MinMeetingRooms(intervals [][]int) int {

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	h := New()

	for _, interval := range intervals {

		if peek, exist := h.Peek(); peek <= interval[0] && exist {
			h.Pop()
		}

		h.Push(interval[1])
	}

	return h.Len()
}

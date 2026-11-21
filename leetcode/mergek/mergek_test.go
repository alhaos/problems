package mergek

import (
	"slices"
	"testing"
)

func build(vals ...int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(l *ListNode) []int {
	res := []int{}
	for ; l != nil; l = l.Next {
		res = append(res, l.Val)
	}
	return res
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"example", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"no lists", [][]int{}, []int{}},
		{"one empty list", [][]int{{}}, []int{}},
		{"several empty", [][]int{{}, {}, {}}, []int{}},
		{"single list", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"empty among full", [][]int{{1, 3}, {}, {2}}, []int{1, 2, 3}},
		{"duplicates", [][]int{{2, 2}, {2}, {2, 2}}, []int{2, 2, 2, 2, 2}},
		{"negatives", [][]int{{-3, 0}, {-5, -1}}, []int{-5, -3, -1, 0}},
		{"uneven lengths", [][]int{{1}, {2, 3, 4, 5, 6}, {0, 7}}, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{"disjoint ranges", [][]int{{7, 8}, {1, 2}, {4, 5}}, []int{1, 2, 4, 5, 7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				lists[i] = build(vals...)
			}
			got := toSlice(mergeKLists(lists))
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

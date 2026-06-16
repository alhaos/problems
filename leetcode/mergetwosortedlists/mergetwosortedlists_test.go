package mergetwosortedlists

import "testing"

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func listToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{2}, []int{1}, []int{1, 2}},
	}

	for _, tt := range tests {
		got := listToSlice(mergeTwoLists(makeList(tt.l1), makeList(tt.l2)))
		if len(got) != len(tt.expected) {
			t.Errorf("l1=%v l2=%v: got %v, want %v", tt.l1, tt.l2, got, tt.expected)
			continue
		}
		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("l1=%v l2=%v: got %v, want %v", tt.l1, tt.l2, got, tt.expected)
				break
			}
		}
	}
}

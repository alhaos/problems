package reverselinkedlist

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

func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		got := listToSlice(reverseList(makeList(tt.input)))
		if len(got) != len(tt.expected) {
			t.Errorf("input %v: got %v, want %v", tt.input, got, tt.expected)
			continue
		}
		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("input %v: got %v, want %v", tt.input, got, tt.expected)
				break
			}
		}
	}
}

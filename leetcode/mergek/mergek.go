package mergek

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type nodeHeap []*ListNode

// value receiver
func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// pointer receiver
func (h *nodeHeap) Push(x any) {
	*h = append((*h), x.(*ListNode))
}

func (h *nodeHeap) Pop() any {
	old := *h
	l := len(*h)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	var h nodeHeap

	for _, l := range lists {
		if l != nil {
			heap.Push(&h, l)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		node := heap.Pop(&h).(*ListNode)
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}

	return dummy.Next
}

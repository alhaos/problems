package leetcode

// Given the head of a singly linked list and two integers left and right where
// left <= right, reverse the nodes of the list from position left to position
// right, and return the reversed list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head.Next == nil {
		return head
	}

	currentNode = head
	var reverseIntervalHead ListNode

	for i:= 0; head.Next != nil; i++ {
		if i > left

	}

}

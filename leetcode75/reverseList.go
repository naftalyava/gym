package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	var tmp *ListNode

	if head == nil {
		return nil
	}

	for head != nil {
		tmp = head
		head = head.Next
		tmp.Next = newHead
		newHead = tmp
	}

	return newHead
}

func main() {

}

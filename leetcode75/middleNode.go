package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	var slow *ListNode = head
	var quick *ListNode = head

	quick = quick.Next

	for quick != nil {
		slow = slow.Next

		if quick.Next != nil {
			quick = quick.Next.Next
		} else {
			quick = nil
		}
	}

	return slow
}

func main() {

}

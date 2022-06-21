package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var slow *ListNode = head
	var fast *ListNode = head.Next

	// traverse list twice and find some node after cycle
	for fast != nil && slow != fast {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}

	if fast == nil {
		return nil
	}

	// at this point fast == slow
	// iterate with fast to find length
	cycle_length := 1
	fast = fast.Next

	for slow != fast {
		cycle_length++
		fast = fast.Next
	}

	// iterate fast cycle_length
	fast = head
	slow = head
	for i := cycle_length; i > 0; i-- {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast

}

func main() {

}

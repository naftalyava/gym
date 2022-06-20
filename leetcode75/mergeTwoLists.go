package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var newList *ListNode
	var curr *ListNode
	var big_list *ListNode = list1
	var small_list *ListNode = list2

	if list1 == nil || list2 == nil {
		return newList
	}

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			small_list = list1
			big_list = list2
		} else {
			small_list = list2
			big_list = list1
		}

		if newList == nil {
			newList = small_list
			curr = small_list
		} else {
			curr.Next = small_list
			curr = curr.Next
		}

		if small_list == list1 {
			list1 = list1.Next
		} else {
			list2 = list2.Next
		}
	}

	small_list.Next = big_list

	return newList
}

func main() {

}

package main

func middleNode(head *ListNode) *ListNode {
	elem := make([]*ListNode, 0)

	for current := head; current != nil; current = current.Next {
		elem = append(elem, current)
	}

	mid := (len(elem) - 1) / 2

	if len(elem)%2 == 1 {
		return elem[mid]
	}

	return elem[mid+1]
}

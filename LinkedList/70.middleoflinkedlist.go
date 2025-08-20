package LinkedList

import "new_apps/rust_build/melonnu/Trees"

func middleNode(head *Trees.ListNode) *Trees.ListNode {
	elem := make([]*Trees.ListNode, 0)

	for current := head; current != nil; current = current.Next {
		elem = append(elem, current)
	}

	mid := (len(elem) - 1) / 2

	if len(elem)%2 == 1 {
		return elem[mid]
	}

	return elem[mid+1]
}

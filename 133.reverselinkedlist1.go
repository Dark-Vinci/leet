package main

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for c := head; c != nil; c = c.Next {
		current := &ListNode{Val: c.Val}
		current.Next = prev
		prev = current
	}

	return prev
}

func reverseListRecur(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}

	newHead := reverseListRecur(h.Next)

	h.Next.Next = h
	h.Next = nil

	return newHead
}

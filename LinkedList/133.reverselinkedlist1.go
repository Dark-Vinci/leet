package LinkedList

import "new_apps/rust_build/melonnu/Trees"

func reverseList(head *Trees.ListNode) *Trees.ListNode {
	var prev *Trees.ListNode

	for c := head; c != nil; c = c.Next {
		current := &Trees.ListNode{Val: c.Val}
		current.Next = prev
		prev = current
	}

	return prev
}

func reverseListRecur(h *Trees.ListNode) *Trees.ListNode {
	if h == nil || h.Next == nil {
		return h
	}

	newHead := reverseListRecur(h.Next)

	h.Next.Next = h
	h.Next = nil

	return newHead
}

package LinkedList

import "new_apps/rust_build/melonnu/Trees"

func swapNodes(head *Trees.ListNode, k int) *Trees.ListNode {
	var (
		front *Trees.ListNode
		back  *Trees.ListNode
		l     = 0
		j     = 0
		i     = 0
	)

	for c := head; c != nil; c = c.Next {
		i++
		j++

		if i == k {
			front = c
		}
	}

	if front == nil {
		return nil
	}

	for c := head; c != nil; c = c.Next {
		l++

		if l == j-k+1 {
			back = c
		}
	}

	front.Val, back.Val = back.Val, front.Val

	return head
}

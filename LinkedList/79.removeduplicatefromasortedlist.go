package LinkedList

import "new_apps/rust_build/melonnu/Trees"

func deleteDuplicates(head *Trees.ListNode) *Trees.ListNode {
	if head == nil {
		return head
	}

	var recurse func(prev, curr *Trees.ListNode)

	recurse = func(prev, curr *Trees.ListNode) {
		if curr == nil {
			return
		}

		if prev != nil && prev.Val == curr.Val {
			prev.Next = curr.Next
			recurse(prev, prev.Next)
			return
		}

		prev = curr
		curr = curr.Next

		recurse(prev, curr)
	}

	recurse(nil, head)

	return head
}

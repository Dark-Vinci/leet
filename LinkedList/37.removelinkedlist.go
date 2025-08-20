package LinkedList

import "new_apps/rust_build/melonnu/Trees"

func removeElements(head *Trees.ListNode, val int) *Trees.ListNode {
	var r func(a, prev *Trees.ListNode)

	r = func(a, prev *Trees.ListNode) {
		if a == nil {
			return
		}

		if a.Val == val && prev == nil {
			head = a.Next
			r(a.Next, nil)
		} else if a.Val == val && prev != nil {
			prev.Next = a.Next
			r(a.Next, prev)
		} else {
			r(a.Next, a)
		}
	}

	r(head, nil)

	return head
}

package main

func reorderList(head *ListNode) {
	var (
		values = make([]int, 0)
		prev   *ListNode
	)

	for c := head; c != nil; c = c.Next {
		values = append(values, c.Val)
	}

	l, r := 0, len(values)-1

	for l <= r {
		left := &ListNode{Val: values[l]}
		right := &ListNode{Val: values[r]}

		if l != r {
			left.Next = right
		}

		if l == 0 {
			*head = *left
		}

		if prev != nil {
			prev.Next = left
		}

		prev = right

		l++
		r--
	}
}

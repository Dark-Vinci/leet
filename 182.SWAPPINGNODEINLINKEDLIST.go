package main

func swapNodes(head *ListNode, k int) *ListNode {
	var (
		front *ListNode
		back  *ListNode
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

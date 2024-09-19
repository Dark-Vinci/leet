package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		newHead         *ListNode
		newHeadTail     *ListNode
		reverseHead     *ListNode
		reverseTail     *ListNode
		lastHead        *ListNode
		lastHeadLast    *ListNode
		isFullyReversed = false
		count           = 0
	)

	for c := head; c != nil; c = c.Next {
		count++
		current := &ListNode{Val: c.Val}

		switch {
		case count < left:
			{
				if newHead == nil {
					newHead = current
					newHeadTail = newHead
				} else {
					newHeadTail.Next = current
					newHeadTail = current
				}
			}

		case isFullyReversed:
			if lastHead == nil {
				lastHead = current
				lastHeadLast = current
			} else {
				lastHeadLast.Next = current
				lastHeadLast = current
			}

		case count >= left && count <= right:
			if count == right {
				isFullyReversed = true
			}

			if reverseTail == nil {
				reverseTail = current
			}

			current.Next = reverseHead
			reverseHead = current
		}
	}

	if newHead == nil && reverseHead != nil {
		reverseTail.Next = lastHead
		return reverseHead
	}

	if reverseHead != nil {
		newHeadTail.Next = reverseHead
		reverseTail.Next = lastHead
	}

	return newHead
}

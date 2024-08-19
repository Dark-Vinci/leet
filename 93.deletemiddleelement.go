package main

func deleteMiddle(head *ListNode) *ListNode {
	count := 0

	for current := head; current != nil; current = current.Next {
		count++
	}

	if count == 1 {
		return nil
	}

	if count == 2 {
		head.Next = nil
		return head
	}

	c := head
	i, half := 0, count/2

	for {
		if i+1 == half {
			c.Next = c.Next.Next
			break
		}

		i++
		c = c.Next
	}

	return head
}

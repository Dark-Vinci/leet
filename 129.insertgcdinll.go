package main

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var newHead *ListNode
	var prev *ListNode

	for c := head; c != nil; c = c.Next {
		if newHead == nil {
			newHead = &ListNode{Val: c.Val}
			prev = newHead

			if c.Next != nil {
				g := gcd(c.Val, c.Next.Val)

				gList := &ListNode{Val: g}
				prev.Next = gList
				prev = prev.Next
			}
			continue
		}

		if c.Next == nil {
			prev.Next = &ListNode{Val: c.Val}
			break
		}

		// add current
		prev.Next = &ListNode{Val: c.Val}
		prev = prev.Next

		// add gcd
		g := gcd(c.Val, c.Next.Val)

		gList := &ListNode{Val: g}
		prev.Next = gList
		prev = prev.Next
	}

	return newHead
}

func gcd(a, b int) int {
	result := 1
	count := 1

	for result <= a && result <= b {
		if a%result == 0 && b%result == 0 {
			count = result
		}

		result++
	}

	return count
}

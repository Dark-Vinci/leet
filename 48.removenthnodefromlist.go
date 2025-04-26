package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var recurse func(node, prev *ListNode) int

	recurse = func(node, prev *ListNode) int {
		if node == nil {
			return 0
		}

		count := recurse(node.Next, node)

		if count+1 == n {
			if prev != nil {
				prev.Next = node.Next
			} else {
				head = head.Next
			}
		}

		return count + 1
	}

	_ = recurse(head, nil)

	return head
}

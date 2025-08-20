package Trees

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := length(head)

	result := make([]*ListNode, 0)

	if count <= k {
		i := 0

		for c := head; c != nil; c = c.Next {
			a := *c
			a.Next = nil

			result = append(result, &a)
			i++
		}

		for i < k {
			result = append(result, nil)
			i++
		}

		return result
	}

	flat := count / k
	remainder := count - (flat * k)
	index := 0

	current, start := head, head

	for current != nil {
		next := current.Next
		if remainder != 0 && index != flat {
			index++
			current = next
			continue
		}

		if remainder != 0 && index == flat {
			c := current
			c.Next = nil

			result = append(result, start)
			start = next
			current = next
			remainder--
			index = 0
			continue
		}

		if index == flat-1 {
			c := current
			c.Next = nil

			result = append(result, start)
			start = next
			current = next
			index = 0
			continue
		}

		current = next
		index++
	}

	return result
}

func length(h *ListNode) int {
	count := 0

	for c := h; c != nil; c = c.Next {
		count += 1
	}

	return count
}

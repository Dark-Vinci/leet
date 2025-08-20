package Sorting

import "new_apps/rust_build/melonnu/Trees"

func mergeKLists(lists []*Trees.ListNode) *Trees.ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]

	for _, node := range lists[1:] {
		head := new(Trees.ListNode)
		decoy := head

		for result != nil && node != nil {
			if result.Val < node.Val {
				decoy.Next = &Trees.ListNode{Val: result.Val}
				result = result.Next
			} else {
				decoy.Next = &Trees.ListNode{Val: node.Val}
				node = node.Next
			}

			decoy = decoy.Next
		}

		for result != nil {
			decoy.Next = &Trees.ListNode{Val: result.Val}

			result = result.Next
			decoy = decoy.Next
		}

		for node != nil {
			decoy.Next = &Trees.ListNode{Val: node.Val}

			node = node.Next
			decoy = decoy.Next
		}

		result = head.Next
	}

	return result
}

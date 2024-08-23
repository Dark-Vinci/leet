package main

func modifiedList(nums []int, head *ListNode) *ListNode {
	var prev *ListNode
	numsDb := make(map[int]struct{})

	for _, v := range nums {
		numsDb[v] = struct{}{}
	}

	var recur func(h *ListNode)

	recur = func(h *ListNode) {
		if h == nil {
			return
		}

		if _, ok := numsDb[h.Val]; !ok {
			prev = h
		} else {
			if prev == nil {
				head = head.Next
			} else {
				prev.Next = h.Next
			}
		}

		recur(h.Next)
	}

	recur(head)

	return head
}

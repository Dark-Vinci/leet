package main

import "slices"

func pairSum(head *ListNode) int {
	db := make([]int, 0)

	for current := head; current != nil; current = current.Next {
		db = append(db, current.Val)
	}

	l := len(db)

	maximus := make([]int, 0)

	for i := 0; i < l/2; i++ {
		evilTwin := db[l-1-i]

		maximus = append(maximus, evilTwin+db[i])
	}

	return slices.Max(maximus)
}

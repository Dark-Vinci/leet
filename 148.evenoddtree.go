package main

import (
	"slices"
)

func isEvenOddTree(root *TreeNode) bool {
	var (
		q      = []*TreeNode{root}
		isEven = true
	)

	for len(q) > 0 {
		l := len(q)
		db := make([]int, 0)

		for i := 0; i < l; i++ {
			current := q[0]
			q = q[1:]

			if (isEven && current.Val%2 == 0) || (!isEven && current.Val%2 != 0) {
				return false
			}

			db = append(db, current.Val)

			if current.Left != nil {
				q = append(q, current.Left)
			}

			if current.Right != nil {
				q = append(q, current.Right)
			}
		}

		if len(db) == 1 {
			isEven = !isEven
			continue
		}

		if len(dedupe(db)) != len(db) {
			return false
		}

		if (isEven && !slices.IsSorted(db)) || (!isEven && slices.IsSorted(db)) {
			return false
		}

		isEven = !isEven
	}

	return true
}

func dedupe(arr []int) []int {
	mp := make(map[int]struct{})
	result := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = struct{}{}
	}

	for k := range mp {
		result = append(result, k)
	}

	return result
}

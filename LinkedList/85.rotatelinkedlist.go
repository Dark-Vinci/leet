package LinkedList

import "new_apps/rust_build/melonnu/Trees"

func revRight(head *Trees.ListNode, k int) *Trees.ListNode {
	db := make([]int, 0)

	for current := head; current != nil; current = current.Next {
		db = append(db, current.Val)
	}

	if len(db) == 0 {
		return nil
	}

	k %= len(db)

	var rev func(i, j int)

	rev = func(i, j int) {
		for i < j {
			db[i], db[j] = db[j], db[i]
			i++
			j--
		}
	}

	rev(0, len(db)-1)
	rev(0, k-1)
	rev(k, len(db)-1)

	nodes := make([]*Trees.ListNode, len(db))

	for i := 0; i < len(db); i++ {
		nodes[i] = &Trees.ListNode{
			Val: db[i],
		}
	}

	for i := range nodes {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
	}

	return nodes[0]
}

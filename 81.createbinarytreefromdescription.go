package main

func createBinaryTree(descriptions [][]int) *TreeNode {
	if len(descriptions) == 0 {
		return nil
	}

	db := make(map[int]*TreeNode)
	visibility := make(map[int]bool)

	for _, description := range descriptions {
		value, child, position := description[0], description[1], description[2]

		if _, ok := db[value]; !ok {
			db[value] = &TreeNode{Val: value}
		}

		if _, ok := db[child]; !ok {
			db[child] = &TreeNode{Val: child}
		}

		if position == 0 {
			db[value].Right = db[child]
		} else {
			db[value].Left = db[child]
		}

		visibility[child] = true
	}

	for _, val := range descriptions {
		if _, ok := visibility[val[0]]; !ok {
			return db[val[0]]
		}
	}

	return nil
}

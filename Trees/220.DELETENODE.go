package Trees

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var (
		dfs    func(left bool, parent, node *TreeNode)
		result = make(map[*TreeNode]struct{})
		db     = make(map[int]struct{})
	)

	for _, val := range to_delete {
		db[val] = struct{}{}
	}

	dfs = func(left bool, parent, node *TreeNode) {
		if node == nil {
			return
		}

		if _, ok := db[node.Val]; ok {
			if _, ko := result[node]; ko {
				delete(result, node)
			}

			if parent != nil {
				if left {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}

			if node.Left != nil {
				result[node.Left] = struct{}{}
				dfs(true, nil, node.Left)
			}

			if node.Right != nil {
				result[node.Right] = struct{}{}
				dfs(false, nil, node.Right)
			}
		} else {
			dfs(true, node, node.Left)
			dfs(false, node, node.Right)
		}
	}

	result[root] = struct{}{}

	dfs(false, nil, root)

	res := make([]*TreeNode, 0)

	for key := range result {
		res = append(res, key)
	}

	return res
}

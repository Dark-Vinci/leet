package Trees

import (
	"fmt"
	"strconv"
)

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	db := make([]string, 0)

	var dfs func(node *TreeNode, values string)

	dfs = func(node *TreeNode, values string) {
		if node == nil {
			return
		}

		values += fmt.Sprintf("%v", node.Val)

		if node.Left == nil && node.Right == nil {
			db = append(db, values)
			return
		}

		dfs(node.Left, values)
		dfs(node.Right, values)
	}

	dfs(root, "")

	sum := 0

	for _, j := range db {
		jj, _ := strconv.Atoi(j)

		sum += jj
	}

	return sum
}

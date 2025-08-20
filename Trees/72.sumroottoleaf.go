package Trees

import (
	"fmt"
	"strconv"
)

func sumRootToLeaf(root *TreeNode) int {
	db := make([]string, 0)

	var dfs func(r *TreeNode, prev string)

	dfs = func(r *TreeNode, prev string) {
		if r == nil {
			return
		}

		prev += fmt.Sprintf("%v", r.Val)

		if r.Left == nil && r.Right == nil {
			db = append(db, prev)
			return
		}

		dfs(r.Left, prev)
		dfs(r.Right, prev)
	}

	dfs(root, "")

	sum := 0

	for _, v := range db {
		numb, _ := strconv.ParseInt(v, 2, 64)

		sum += int(numb)
	}

	return sum
}

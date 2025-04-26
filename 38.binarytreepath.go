package main

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	result := make([]string, 0)

	var dfs func(a *TreeNode, path string, start bool)

	dfs = func(a *TreeNode, path string, start bool) {
		if a == nil {
			return
		}

		if start {
			path = fmt.Sprintf("%v", a.Val)
		} else {
			path += fmt.Sprintf("->%v", a.Val)
		}

		if a.Left == nil && a.Right == nil {
			result = append(result, path)
			return
		}

		dfs(a.Left, path, false)
		dfs(a.Right, path, false)
	}

	dfs(root, "", true)

	return result
}

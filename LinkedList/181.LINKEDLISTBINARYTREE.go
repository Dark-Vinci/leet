package LinkedList

import (
	"cmp"
	"new_apps/rust_build/melonnu/Trees"
)

func isSubPath(head *Trees.ListNode, root *Trees.TreeNode) bool {
	var dfs func(c *Trees.ListNode, h *Trees.TreeNode) bool
	dfs = func(c *Trees.ListNode, h *Trees.TreeNode) bool {
		if (h == nil && c == nil) || (c == nil && h != nil) {
			return true
		}

		if (c != nil && h == nil) || (h.Val != c.Val) {
			return false
		}

		// LEFT [NEXT LIST]
		if dfs(c.Next, h.Left) {
			return true
		}

		// RIGHT [NEXT LIST]
		if dfs(c.Next, h.Right) {
			return true
		}

		return false
	}

	if root == nil {
		return false
	}

	if dfs(head, root) {
		return true
	}

	l, r := isSubPath(head, root.Left), isSubPath(head, root.Right)

	return cmp.Or(l, r)
}

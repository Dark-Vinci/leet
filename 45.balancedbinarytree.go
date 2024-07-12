package main

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    var dfs func(node *TreeNode) int

    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        l, r := dfs(node.Left), dfs(node.Right)

        if l == -1 || r == -1 || abs(l - r) > 1 {
            return -1
        }

        if l > r {
            return 1 + l
        }

        return 1 + r
    }

    return dfs(root) >= 0
}

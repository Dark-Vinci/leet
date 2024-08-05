package main

import "slices"

// BT AND BST
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func(r *TreeNode) ([]int, *TreeNode)

    dfs = func (r *TreeNode) ([]int, *TreeNode) {
        if r == nil {
            return []int{}, nil
        }

        left, lNode := dfs(r.Left)

        if lNode != nil {
            return []int{}, lNode
        }

        right, rNode := dfs(r.Right)

        if rNode != nil {
            return []int{}, rNode
        }

        curr := make([]int, 0)

        curr = append(curr, r.Val)

        curr = slices.Concat(left, right, curr)

        if slices.Contains(curr, p.Val) && slices.Contains(curr, q.Val) {
            return curr, r
        }

        return curr, nil
    }

    _, result := dfs(root)

    return result
}

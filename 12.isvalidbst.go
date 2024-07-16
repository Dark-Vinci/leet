package main

import "math"

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    var isBST func(r *TreeNode, min, max int) bool

    isBST = func(r *TreeNode, min, max int) bool {
        if r == nil {
            return true
        }

        if r.Val >= max || r.Val <= min {
            return false
        }

        return isBST(r.Left, min, r.Val) && isBST(r.Right, r.Val, max)
    }

    return isBST(root, math.MinInt, math.MaxInt)
}


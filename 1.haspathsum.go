package main
 
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type ListNode struct {
    Val int
    Next *ListNode
}


func hasPathSum(root *TreeNode, targetSum int) bool {
    var pathSum func(a *TreeNode, sum int) bool

    pathSum = func(a *TreeNode, sum int) bool {
        if a == nil {
            return false
        }

        sum += a.Val

        if a.Left == nil && a.Right == nil {
            return sum == targetSum
        }

        return pathSum(a.Left, sum) || pathSum(a.Right, sum)
    }

    return pathSum(root, 0)
}
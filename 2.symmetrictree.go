package main

func isSymmetric(root *TreeNode) bool {
    var isSym func(r1 *TreeNode, r2 *TreeNode) bool

    isSym = func(r1 *TreeNode, r2 *TreeNode) bool {
        if r1 == nil && r2 == nil {
            return true
        }

        if r1 == nil || r2 == nil {
            return false
        }

        if r1.Val != r2.Val {
            return false
        }

        return isSym(r1.Left, r2.Right) && isSym(r1.Right, r2.Left)
    }

    return isSym(root, root)
}

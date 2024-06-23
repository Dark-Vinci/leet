package main

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    l := inorderTraversal(root.Left)
    r := inorderTraversal(root.Right)

    res := make([]int, 0)

    for _, v := range l {
        res = append(res, v)
    }

    res = append(res, root.Val)

    for _, v := range r {
        res = append(res, v)
    }

    return res
}

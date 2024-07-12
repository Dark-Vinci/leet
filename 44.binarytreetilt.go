package main

func findTilt(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var dfs func (node *TreeNode) (*TreeNode, int)

    dfs = func (node *TreeNode) (*TreeNode, int) {
        if node == nil {
            return nil, 0
        }

        if node.Left == nil && node.Right == nil {
            return &TreeNode{ Val: 0 }, node.Val
        }

        lNode, lValues := dfs(node.Left)
        rNode, rValues := dfs(node.Right)

        diff := abs(lValues - rValues)

        newVal := node.Val + rValues + lValues

        return &TreeNode{
            Val: diff,
            Left: lNode,
            Right: rNode,
        },
        newVal
    }

    node, _ := dfs(root)

    var dfsSum func(node *TreeNode) int

    dfsSum = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        return node.Val + dfsSum(node.Left) + dfsSum(node.Right)
    }

    return dfsSum(node)
}

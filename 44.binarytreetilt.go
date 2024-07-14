package main

func findTilt(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    sum := 0

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
        
        sum += diff

        newVal := node.Val + rValues + lValues

        return &TreeNode{
            Val: diff,
            Left: lNode,
            Right: rNode,
        },
        newVal
    }

    dfs(root)

    return sum
}

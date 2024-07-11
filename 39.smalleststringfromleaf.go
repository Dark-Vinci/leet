package main

import "sort"

func smallestFromLeaf(root *TreeNode) string {
    if root == nil {
        return ""
    }

    db := make([]string, 0)

    var dfs func(node *TreeNode, bag []byte)

    dfs = func(node *TreeNode, bag []byte) {
        if node == nil {
            return
        }

        val := byte(97 + node.Val)

        bag = append([]byte{val}, bag...)

        if node.Left == nil && node.Right == nil {
            db = append(db, string(bag))
            return
        }

        dfs(node.Left, bag)
        dfs(node.Right, bag)
    }

    dfs(root, []byte{})

    sort.Strings(db)

    return db[0]
}

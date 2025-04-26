package main

func sortedListToBST(head *ListNode) *TreeNode {
	linkedListSlice := make([]int, 0)

	for current := head; current != nil; current = current.Next {
		linkedListSlice = append(linkedListSlice, current.Val)
	}

	var dfs func(l, r int) *TreeNode

	dfs = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		mid := l + (r-l)/2

		left := dfs(l, mid-1)
		right := dfs(mid+1, r)

		return &TreeNode{
			Val:   linkedListSlice[mid],
			Left:  left,
			Right: right,
		}
	}

	return dfs(0, len(linkedListSlice)-1)
}

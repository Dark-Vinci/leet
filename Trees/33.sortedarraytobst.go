package Trees

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
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
			Val:   nums[mid],
			Left:  left,
			Right: right,
		}
	}

	return dfs(0, len(nums)-1)
}

package Trees

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var (
		dfs         func(i int) bool
		visited     = make(map[int]struct{})
		parentCount = make([]int, n)
		root        = -1
	)

	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			parentCount[leftChild[i]]++

			if parentCount[leftChild[i]] > 1 {
				return false
			}
		}

		if rightChild[i] != -1 {
			parentCount[rightChild[i]]++

			if parentCount[rightChild[i]] > 1 {
				return false
			}
		}
	}

	for i := 0; i < n; i++ {
		if parentCount[i] == 0 {
			if root != -1 {
				return false
			}

			root = i
		}
	}

	if root == -1 {
		return false
	}

	dfs = func(i int) bool {
		if i >= n {
			return true
		}

		if i < 0 {
			return true
		}

		if _, ok := visited[i]; ok {
			return false
		}

		visited[i] = struct{}{}

		left, right := dfs(leftChild[i]), dfs(rightChild[i])

		return left && right
	}

	if !dfs(root) {
		return false
	}

	for i := 0; i < n; i++ {
		if _, ok := visited[i]; !ok {
			return false
		}
	}

	return true
}

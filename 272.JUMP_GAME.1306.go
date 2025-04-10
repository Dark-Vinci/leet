package main

func canReach1(arr []int, start int) bool {
	var (
		used = make(map[int]struct{})
		dfs  func(i int) bool
	)

	dfs = func(i int) bool {
		if i >= len(arr) || i < 0 {
			return false
		}

		if _, ok := used[i]; ok {
			return false
		}

		used[i] = struct{}{}

		if arr[i] == 0 {
			return true
		}

		if dfs(i + arr[i]) {
			return true
		}

		return dfs(i - arr[i])
	}

	return dfs(start)
}

package dynamicProgramming

func splitArraySameAverage__(nums []int) bool {
	var (
		sum  = 0
		dfs  func(sum int, i int, count int) bool
		memo = make(map[[3]int]bool)
		l    = len(nums)
	)

	for _, v := range nums {
		sum += v
	}

	avg := float64(sum) / float64(l)

	dfs = func(sum int, i int, count int) bool {
		k := [3]int{sum, i, count}

		if v, ok := memo[k]; ok {
			return v
		}

		if i >= l || count >= l {
			memo[k] = false
			return false
		}

		if count != 0 && float64(sum)/float64(count) == avg {
			memo[k] = true
			return memo[k]
		}

		if dfs(sum+nums[i], i+1, count+1) {
			memo[k] = true
			return memo[k]
		}

		memo[k] = dfs(sum, i+1, count)
		return memo[k]
	}

	return dfs(0, 0, 0)
}

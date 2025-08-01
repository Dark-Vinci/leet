package dynamicProgramming

func findNumberOfLIS(nums []int) int {
	var (
		dp   func(curr int, prev int, memo [][][2]int) (int, int)
		memo = make([][][2]int, len(nums))
	)

	for i := range len(nums) {
		memo[i] = make([][2]int, len(nums))

		for j := range len(nums) {
			memo[i][j] = [2]int{-1, -1}
		}
	}

	dp = func(curr int, prev int, memo [][][2]int) (int, int) {
		if curr == len(nums) {
			return 0, 1
		}

		if v := memo[curr][prev+1]; v[0] != -1 {
			return v[0], v[1]
		}

		dontl, dontc := dp(curr+1, prev, memo)
		dol, doc := 0, 0

		if prev == -1 || nums[curr] > nums[prev] {
			dol, doc = dp(curr+1, curr, memo)
			dol += 1
		}

		if dol > dontl {
			memo[curr][prev+1] = [2]int{dol, doc}
		} else if dol == dontl {
			memo[curr][prev+1] = [2]int{dol, dontc + doc}
		} else {
			memo[curr][prev+1] = [2]int{dontl, dontc}
		}

		return memo[curr][prev+1][0], memo[curr][prev+1][1]
	}

	_, result := dp(0, -1, memo)

	return result
}

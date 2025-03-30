package main

import "fmt"

func rob11(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var dp func(i, j int) int

	dp = func(i, j int) int {
		dp1, dp2 := 0, 0

		for ii := i; ii <= j; ii++ {
			mx := max(dp2+nums[ii], dp1)
			dp1, dp2 = mx, dp1
		}

		return dp1
	}

	return dp(0, len(nums)-1)
}

func rob1(nums []int) int {
	var (
		dfs  func(sum int, i int, pickedPrev bool) int
		memo = make(map[string]int)
	)

	dfs = func(sum int, i int, pickedPrev bool) int {
		key := fmt.Sprintf("|%v|%v|%v|", sum, i, pickedPrev)

		if v, ok := memo[key]; ok {
			return v
		}

		if i == len(nums) {
			return sum
		}

		if i == 0 {
			p, n := dfs(sum+nums[i], i+1, true), dfs(sum, i+1, false) //pick, dont pick
			memo[key] = max(p, n)
			return memo[key]
		}

		if pickedPrev {
			memo[key] = dfs(sum, i+1, false) // dont pick
			return memo[key]
		}

		p, n := dfs(sum+nums[i], i+1, true), dfs(sum, i+1, false) // pick, dont pick
		memo[key] = max(p, n)

		return memo[key]
	}

	return dfs(0, 0, false)
}

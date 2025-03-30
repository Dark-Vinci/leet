package main

import "fmt"

func rob22(nums []int) int {
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

	return max(dp(0, len(nums)-2), dp(1, len(nums)-1))
}

func rob2(nums []int) int {
	var (
		dfs  func(sum int, i int, pickedPrev, pickedInitial bool) int
		memo = make(map[string]int)
	)

	dfs = func(sum int, i int, pickedPrev, pickedInitial bool) int {
		key := fmt.Sprintf("|%v|%v|%v|%v|", sum, i, pickedPrev, pickedInitial)

		if v, ok := memo[key]; ok {
			return v
		}

		if i == len(nums) {
			memo[key] = sum
			return memo[key]
		}

		if i == len(nums)-1 && pickedInitial {
			memo[key] = dfs(sum, i+1, false, pickedInitial) // dont pick
			return memo[key]
		}

		if i == 0 {
			r, nr := dfs(sum+nums[i], i+1, true, true), dfs(sum, i+1, false, false) //pick| dont pick
			memo[key] = max(r, nr)

			return memo[key]
		}

		if pickedPrev {
			memo[key] = dfs(sum, i+1, false, pickedInitial) // dont pick
			return memo[key]
		}

		r, nr := dfs(sum+nums[i], i+1, true, pickedInitial), dfs(sum, i+1, false, pickedInitial) // pick | dont pick
		memo[key] = max(r, nr)

		return memo[key]
	}

	return dfs(0, 0, false, false)
}

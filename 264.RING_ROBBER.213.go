package main

import "fmt"

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

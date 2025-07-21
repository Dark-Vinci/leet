package dynamicProgramming

import "slices"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	slices.Sort(nums)

	var dp func(prev, current int, memo map[[2]int][]int) []int

	dp = func(prev, current int, memo map[[2]int][]int) []int {
		if current == len(nums) {
			return []int{}
		}

		key := [2]int{prev, current}
		if v, ok := memo[key]; ok {
			return v
		}

		dont := dp(prev, current+1, memo)

		if prev == -1 || nums[current]%nums[prev] == 0 || nums[prev]%nums[current] == 0 {
			do := dp(current, current+1, memo)

			next := append([]int{nums[current]}, do...)

			if len(next) > len(dont) {
				dont = next
			}
		}

		memo[key] = dont

		return dont
	}

	return dp(-1, 0, map[[2]int][]int{})
}

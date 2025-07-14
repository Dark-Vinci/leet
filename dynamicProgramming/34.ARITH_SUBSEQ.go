package dynamicProgramming

import "sort"

//https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference/description/
// difficulty; medium
// lesson; -> you can use memo and still get TLE, that's why tabulation is better

func longestSubsequence(arr []int, difference int) int {
	var (
		dfs    func(i int) int
		pos    = make(map[int][]int)
		memo   = make(map[int]int)
		result = 0
	)

	for i, v := range arr {
		pos[v] = append(pos[v], i)
	}

	dfs = func(i int) int {
		if v, ok := memo[i]; ok {
			return v
		}

		next := arr[i] + difference
		result := 1

		if indexes, ok := pos[next]; ok {
			j := sort.Search(len(indexes), func(k int) bool {
				return indexes[k] > i
			})

			if j < len(indexes) {
				result = 1 + dfs(indexes[j])
			}
		}

		memo[i] = result

		return result
	}

	for i := range arr {
		result = max(result, dfs(i))
	}

	return result
}

// TLE
func longestSubsequence__(arr []int, difference int) int {
	var dp func(prev, i int, memo map[[2]int]int) int

	dp = func(prev, i int, memo map[[2]int]int) int {
		if i >= len(arr) {
			return 0
		}

		key := [2]int{prev + 1, i}

		if v, ok := memo[key]; ok {
			return v
		}

		dont := dp(prev, i+1, memo)

		do := 0

		if prev == -1 || arr[i]-arr[prev] == difference {
			do = 1 + dp(i, i+1, memo)
		}

		memo[key] = max(do, dont)

		return memo[key]
	}

	return dp(-1, 0, map[[2]int]int{})
}

// TLE
func longestSubsequence_(arr []int, difference int) int {
	var dp func(i int, prev int) int
	memo := make(map[[2]int]int)

	dp = func(i int, prev int) int {
		if i >= len(arr) {
			return 1
		}

		key := [2]int{i, prev}

		if v, ok := memo[key]; ok {
			return v
		}

		if arr[i]-prev != difference {
			memo[key] = dp(i+1, prev)
		} else {
			memo[key] = 1 + dp(i+1, arr[i])
		}

		return memo[key]
	}

	result := 1

	for i, v := range arr {
		if result > len(arr)-i+1 {
			break
		}

		result = max(result, dp(i+1, v))
	}

	return result
}

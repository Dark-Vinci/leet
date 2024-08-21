package main

import (
	"fmt"
	"slices"
)

// THIS IS A COMBINATION SOLUTION THAT WON'T WORK FOR THE PROBLEM
func incremovableSubarrayCount(nums []int) int {
	result := 0

	var dfs func([]int, int)

	dfs = func(db []int, count int) {
		if count == len(nums) {
			if isIncreasing(db) {
				fmt.Println(db)
				result++
			}

			return
		}

		cp := slices.Clone(db)
		cp = append(cp, nums[count])

		count++

		dfs(cp, count)
		dfs(db, count)
	}

	dfs([]int{}, 0)

	return result
}

func isIncreasing(a []int) bool {
	for i := 0; i < len(a); i++ {
		if i > 0 && a[i-1] >= a[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))
}

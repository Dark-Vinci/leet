package main

import (
	"fmt"
	"slices"
)

func incremovableSubarrayCount(nums []int) int {
	result := 0
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			transient := make([]int, 0)

			for k := 0; k < l; k++ {
				if k < i || k > j {
					transient = append(transient, nums[k])
				}
			}

			if isIncreasing(transient) {
				result++
			}
		}
	}

	return result
}

func isIncreasing(a []int) bool {
	for i, j := 0, 1; i < len(a)-1; i, j = i+1, j+1 {
		if a[i] >= a[j] {
			return false
		}
	}

	return true
}

// THIS IS A COMBINATION SOLUTION THAT WON'T WORK FOR THE PROBLEM
func incremovableSubarrayCountNOT(nums []int) int {
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

//
//func main() {
//	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))
//}

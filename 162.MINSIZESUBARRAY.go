package main

import (
	"fmt"
	"math"
	"slices"
)

// THIS WAS A STRUGGLE BROTHER
func minSubArrayLen(target int, nums []int) int {
	l, result, pref := len(nums), math.MaxInt, slices.Clone(nums)
	pref = append([]int{0}, pref...)

	for i := 1; i <= l; i++ {
		if pref[i] >= target {
			return 1
		}

		pref[i] += pref[i-1]
	}

	for i := l; i >= 0; i-- {
		c := pref[i]

		for j := 0; j < i; j++ {
			change := c - pref[j]
			if change >= target {
				result = min(result, i-j)
			}

			if change < target {
				break
			}
		}
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}

func minSubArrayLen0(target int, nums []int) int {
	result := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return 1
		}

		for j := i + 1; j < len(nums); j++ {
			fmt.Println(nums[i : j+1])
			if sum(nums[i:j+1]) == target {
				result = min(result, j-i+1)
				continue
			}
		}
	}

	fmt.Println(result)

	if result == math.MaxInt {
		return 0
	}

	return result
}

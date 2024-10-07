package main

import "slices"

func triangularSum(nums []int) int {
	result := slices.Clone(nums)

	for len(result) > 1 {
		for i := 0; i < len(result); i++ {
			if i > 0 {
				result[i-1] += result[i] % 10
			}
		}

		result = result[:len(result)-1]
	}

	return result[0] % 10
}

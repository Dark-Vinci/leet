package main

import "slices"

func lengthOfLIS(nums []int) int {
	result := make([]int, 0)

	for _, v := range nums {
		if len(result) == 0 || result[len(result)-1] < v {
			result = append(result, v)
			continue
		}

		index, _ := slices.BinarySearch(result, v)
		result[index] = v
	}

	return len(result)
}

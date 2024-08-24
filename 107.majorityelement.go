package main

import "math"

func majorityElement(nums []int) int {
	db := make(map[int]int)

	for _, v := range nums {
		db[v] += 1
	}

	maximum := math.MinInt
	result := math.MinInt

	for k, v := range db {
		if v > maximum {
			maximum = v
			result = k
		}
	}

	return result
}

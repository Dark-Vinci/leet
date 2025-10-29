package arrays

import "new_apps/rust_build/melonnu/Trees"

func maximumJumps(nums []int, target int) int {
	table, l := make([]int, len(nums)), len(nums)

	for i := 0; i < len(table); i++ {
		table[i] = -1
	}

	table[0] = 0

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if Trees.abs(nums[i]-nums[j]) <= target && table[j] != -1 {
				table[i] = max(table[i], table[j]+1)
			}
		}
	}

	return table[l-1]
}

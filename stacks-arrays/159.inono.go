package stacks_arrays

import (
	"math"
)

// TIME LIMIT EXCEEDED
func minSum(nums []int, target int) int {
	var dfs func(sum int, used map[int]struct{})
	result := math.MaxInt

	dfs = func(sum int, used map[int]struct{}) {
		if sum >= target {
			result = min(result, len(used))
			return
		}

		if len(used) == len(nums) {
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := used[i]; !ok {
				sum += nums[i]
				used[i] = struct{}{}
				dfs(sum, used)
				delete(used, i)
				sum -= nums[i]
			}
		}
	}

	dfs(0, map[int]struct{}{})

	if result == math.MaxInt {
		return 0
	}

	return result
}

//func minSubArrayLen(target int, nums []int) int {
//	result := math.MaxInt
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == target {
//			return 1
//		}
//
//		for j := i + 1; j < len(nums); j++ {
//			fmt.Println(nums[i : j+1])
//			if sum(nums[i:j+1]) == target {
//				result = min(result, j-i+1)
//				continue
//			}
//		}
//	}
//
//	fmt.Println(result)
//
//	if result == math.MaxInt {
//		return 0
//	}
//
//	return result
//}

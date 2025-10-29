package arrays

import (
	"fmt"
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)

	var (
		result = make([][]int, 0)
		db     = make(map[string]struct{})
	)

	for i := 0; i <= len(nums)-4; i++ {
		a := nums[i]

		for j := len(nums) - 1; j > i; j-- {
			b := nums[j]

			l, r := i+1, j-1

			for l < r {
				c, d := nums[l], nums[r]

				if a+b+c+d < target {
					l++
					continue
				}

				if a+b+c+d > target {
					r--
					continue
				}

				arr := []int{a, b, c, d}

				arrStr := fmt.Sprintf("%v%v%v%v", a, b, c, d)

				if _, ok := db[arrStr]; ok {
					l++
					r--
					continue
				}

				db[arrStr] = struct{}{}

				// equal at this point
				result = append(result, arr)
				l++
				r--
			}
		}
	}

	return result
}

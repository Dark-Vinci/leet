package dynamicProgramming

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision_(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprintf("%v", nums[0])
	}

	if len(nums) == 2 {
		return fmt.Sprintf("%v/%v", nums[0], nums[1])
	}

	var result strings.Builder

	result.WriteString(fmt.Sprintf("%v/(", nums[0]))

	for i := 1; i < len(nums)-1; i++ {
		result.WriteString(fmt.Sprintf("%v/", nums[i]))
	}

	result.WriteString(fmt.Sprintf("%v)", nums[len(nums)-1]))

	return result.String()
}

type res struct {
	maxStr string
	minStr string
	mx, mn float64
}

func optimalDivision(nums []int) string {
	var (
		memo   = make(map[[2]int]res)
		helper func(i, j int) res
	)

	helper = func(i, j int) res {
		key := [2]int{i, j}

		if v, ok := memo[key]; ok {
			return v
		}

		if i == j {
			mx, mxStr := float64(nums[i]), strconv.Itoa(nums[i])

			memo[key] = res{mx: mx, mn: mx, minStr: mxStr, maxStr: mxStr}

			return memo[key]
		}

		result := res{mn: 1e9, mx: -1e9}

		for k := i; k < j; k++ {
			left, right := helper(i, k), helper(k+1, j)

			mx := left.mx / right.mn
			mxStr := ""
			if k+1 == j {
				mxStr += fmt.Sprintf("%v/%v", left.maxStr, right.minStr)
			} else {
				mxStr += fmt.Sprintf("%v/(%v)", left.maxStr, right.minStr)
			}

			if mx > result.mx {
				result.mx = mx
				result.maxStr = mxStr
			}

			mn := left.mn / right.mx
			mnStr := ""
			if k+1 == j {
				mnStr += fmt.Sprintf("%v/%v", left.minStr, right.maxStr)
			} else {
				mnStr += fmt.Sprintf("%v/(%v)", left.minStr, right.maxStr)
			}

			if mn < result.mn {
				result.mn = mn
				result.minStr = mnStr
			}
		}

		memo[key] = result

		return memo[key]
	}

	return helper(0, len(nums)-1).maxStr
}

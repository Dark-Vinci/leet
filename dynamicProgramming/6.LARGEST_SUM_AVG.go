package dynamicProgramming

import (
	"math"
)

func largestSumOfAverages(nums []int, k int) float64 {
    var (
        memo = make(map[[2]int]float64)
        helper func(k, i int) float64
        l = len(nums)
    )

    helper = func(k, i int) float64 {
        if k < 0 || i >= l {
            return 0.0
        }

        if k == 0 {
            return math.MinInt
        }

        key := [2]int{k, i}

        if v, ok := memo[key]; ok {
            return v
        }

        res, sum := 0.0, 0

        for j := i; j < l; j++ {
            sum += nums[j]

            ns := (float64(sum) / float64(j - i + 1)) + helper(k - 1, j + 1)

            res = max(res, ns)
        }

        memo[key] = res

        return memo[key]
    }

    return helper(k, 0)
}
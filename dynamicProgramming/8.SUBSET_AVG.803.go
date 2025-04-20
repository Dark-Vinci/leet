package dynamicProgramming

import (
	"slices"
)

func splitArraySameAverage(nums []int) bool {
    var (
        n = len(nums)
        possible = false
        totalSum = 0
        memo = make(map[[2]int]map[int]bool)
        dfs func(i, k, target int) bool
    )

    for _, v := range nums {
        totalSum += v
    }

    for k := 1; k <= n/2; k++ {
        if (totalSum * k) % n == 0 {
            possible = true
            break
        }
    }

    if !possible {
        return false
    }

    slices.Sort(nums)

    dfs = func(i, k, target int) bool {
        if k == 0 {
            return target == 0
        }

        if i >= n || k > n-i {
            return false
        }

        if memo[[2]int{i, k}] != nil {
            if v, ok := memo[[2]int{i, k}][target]; ok {
                return v
            }
        } else {
            memo[[2]int{i, k}] = make(map[int]bool)
        }

        if dfs(i+1, k-1, target-nums[i]) {
            memo[[2]int{i, k}][target] = true
            return true
        }

        if dfs(i+1, k, target) {
            memo[[2]int{i, k}][target] = true
            return true
        }

        memo[[2]int{i, k}][target] = false
        return false
    }

    for k := 1; k <= n/2; k++ {
        if (totalSum * k) % n == 0 {
            t := (totalSum * k) / n

            if dfs(0, k, t) {
                return true
            }
        }
    }

    return false
}


func splitArraySameAverage_(nums []int) bool {
    var (
        sum = 0
        dfs func(sum int, i int, count int) bool
        memo = make(map[[3]int]bool)
        l = len(nums)
    )

    for _, v := range nums {
        sum += v
    }

    avg := float64(sum) / float64(l)

    dfs = func(sum int, i int, count int) bool {
        k := [3]int{sum, i, count}

        if v, ok := memo[k]; ok {
            return v
        }

        if i >= l || count >= l  {
            memo[k] = false
            return false
        }

        if count != 0 && float64(sum) / float64(count) == avg {
            memo[k] = true
            return memo[k]
        }

        if dfs(sum + nums[i], i + 1, count + 1) {
            memo[k] = true
            return memo[k]
        }

        memo[k] = dfs(sum, i + 1, count)
        return memo[k]
    }

    return dfs(0, 0, 0)
}

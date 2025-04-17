package dynamicProgramming

func canPartition_(nums []int) bool {
    sum := 0

    for _, v := range nums {
        sum += v
    }

    if sum % 2 != 0 {
        return false
    }

    var (
        h = sum / 2
        dp = make([]bool, h + 1)
    )

    dp[0] = true

    for _, v := range nums {
        for j := h; j >= v; j-- {
            dp[j] = dp[j] || dp[j - v]
        }
    }

    return dp[h]
}

func canPartition(nums []int) bool {
    sum := 0

    for _, v := range nums {
        sum += v
    }

    if sum % 2 != 0 {
        return false
    }

    var (
        h = sum / 2
        memo = make(map[[2]int]bool)
        dfs func(i int, s int) bool
    )

    dfs = func(i int, s int) bool {
        if i >= len(nums) {
            return false
        }

        key := [2]int{i, s}

        if v, ok := memo[key]; ok {
            return v
        } 

        if s > h {
            memo[key] = false
            return false
        }

        if s == h {
            memo[key] = true
            return true
        }

        add, dont := dfs(i + 1, s + nums[i]), dfs(i + 1, s)

        if add || dont {
            memo[key] = true
            return true
        }

        memo[key] = false
        return false
    }

    return dfs(0, 0)
}

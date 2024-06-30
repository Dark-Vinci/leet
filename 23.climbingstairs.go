package main

func climbStairs(n int) int {
    if n < 3 {
        return n
    }

    var climb func(n int, memo map[int]int) int

    climb = func(n int, memo map[int]int) int {
        if n < 3 {
            return n
        }

        if v, ok := memo[n]; ok {
            return v
        }

        memo[n] = climb(n - 2, memo) + climb(n - 1, memo)

        return memo[n]
    }

    return climb(n, map[int]int{})
}

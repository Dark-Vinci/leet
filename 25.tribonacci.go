package main

func tribonacci(n int) int {
    var trib func(n int, memo map[int]int) int

    trib = func(n int, memo map[int]int) int {
        if v, ok := memo[n]; ok {
            return v
        }

        if n == 0 {
            return 0
        }

        if n <= 2 {
            return 1
        }

        memo[n] = trib(n - 1, memo) + trib(n - 2, memo) + trib(n - 3, memo)

        return memo[n]
    } 

    return trib(n, map[int]int{})
}

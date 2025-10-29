package arrays

import "fmt"

func maxProfit2(prices []int) int {
	var (
		dfs func(i int, canBuy bool) int
		dp  = make(map[string]int)
	)

	dfs = func(i int, canBuy bool) int {
		if i >= len(prices) {
			return 0
		}

		key := fmt.Sprintf("%v|%v", i, canBuy)

		if v, ok := dp[key]; ok {
			return v
		}

		cool := dfs(i+1, canBuy)
		op := 0 // buy/sell

		if canBuy {
			op = dfs(i+1, !canBuy) - prices[i]
		} else {
			op = dfs(i+2, !canBuy) + prices[i]
		}

		dp[key] = max(op, cool)

		return dp[key]
	}

	return dfs(0, true)
}

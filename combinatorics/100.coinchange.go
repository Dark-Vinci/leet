package combinatorics

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

// WITH ALL MY OPTIMIZATION, THIS WON'T HAVE DESIRED TIME LIMIT
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if sum(coins) == amount {
		return len(coins)
	}

	slices.SortFunc(coins, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	for i := 0; i < len(coins); i++ {
		if amount%coins[i] == 0 && coins[i] != 1 {
			return amount / coins[i]
		}
	}

	result := math.MaxInt

	cache := make(map[string]struct{})

	var dfs func(pocket []int)

	dfs = func(pocket []int) {
		str := ""

		for _, v := range pocket {
			str += fmt.Sprintf("%v", v)
		}

		if _, ok := cache[str]; ok {
			return
		} else {
			cache[str] = struct{}{}
		}

		if sum(pocket) > amount {
			return
		}

		if sum(pocket) == amount && len(pocket) < result {
			result = len(pocket)
		}

		for i := 0; i < len(coins); i++ {
			cp := append(pocket, coins[i])
			fmt.Println(coins)
			dfs(cp)
		}
	}

	dfs([]int{})

	if result == math.MaxInt {
		return -1
	}

	return result
}

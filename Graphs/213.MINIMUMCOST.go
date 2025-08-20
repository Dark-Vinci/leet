package Graphs

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var (
		list = make(map[int][]int)
		prc  = make(map[[2]int]int)
		dfs  func(from, stop int) int
		memo = make(map[[2]int]int)
	)

	for _, val := range flights {
		u, v, p := val[0], val[1], val[2]
		list[u] = append(list[u], v)
		prc[[2]int{u, v}] = p
	}

	dfs = func(from, stop int) int {
		if stop > k+1 {
			return math.MaxInt
		}

		if from == dst {
			return 0
		}

		key := [2]int{from, stop}

		if val, ok := memo[key]; ok {
			return val
		}

		mn := math.MaxInt

		for _, next := range list[from] {
			cost := dfs(next, stop+1)

			if cost != math.MaxInt {
				mn = min(mn, prc[[2]int{from, next}]+cost)
			}
		}

		memo[key] = mn
		return mn
	}

	result := dfs(src, 0)

	if result == math.MaxInt {
		return -1
	}

	return result
}

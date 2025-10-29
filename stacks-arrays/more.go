package arrays

import (
	"fmt"
	"math"
)

func minimumJumps(forbidden []int, a int, b int, x int) int {
	mn := math.MaxInt

	var dfs func(back bool, val int, count int)

	fMap, used := make(map[int]struct{}), make(map[int]struct{})

	for _, v := range forbidden {
		fMap[v] = struct{}{}
	}

	dfs = func(back bool, val int, count int) {
		if val < 0 || count > mn {
			return
		}

		if _, ok := fMap[val]; ok {
			return
		}

		if val == x {
			mn = min(mn, count)
			return
		}

		if _, ok := used[val]; ok {
			return
		}

		used[val] = struct{}{}

		if !back {
			dfs(true, val-b, count+1)
		}

		dfs(false, val+a, count+1)
	}

	dfs(false, 0, 0)

	return mn
}

func main() {
	res := minimumJumps([]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7)

	fmt.Println(res)
}

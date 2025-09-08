package dynamicProgramming

import (
	"math"
	"slices"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	var (
		r    = len(robot)
		fac  = make([]int, 0)
		memo = make([][]int64, r+1)
		dfs  func(rInd, fInd int, memo [][]int64) int64
		m    int
	)

	for _, f := range factory {
		for i := 0; i < f[1]; i++ {
			fac = append(fac, f[0])
		}
	}

	m = len(fac)

	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int64, m+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	dfs = func(rInd, fInd int, memo [][]int64) int64 {
		if rInd < 0 {
			return 0
		}

		if fInd < 0 {
			return math.MaxInt64
		}

		if v := memo[rInd][fInd]; v != -1 {
			return v
		}

		dont, do := dfs(rInd, fInd-1, memo), dfs(rInd-1, fInd-1, memo)

		if do != math.MaxInt64 {
			do += int64(abs(robot[rInd] - fac[fInd]))

			dont = min(do, dont)
		}

		memo[rInd][fInd] = dont

		return memo[rInd][fInd]
	}

	slices.Sort(robot)
	slices.Sort(fac)

	return dfs(r-1, m-1, memo)
}

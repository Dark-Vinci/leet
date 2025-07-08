package dynamicProgramming

// lcs
// https://leetcode.com/problems/knight-dialer/
// medium
// lesson: use array of arrays in place of map[[3]int]int map of arrays

func knightDialer(n int) int {
	pattern := [4][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{-1, 0, -1},
	}

	dirs := [8][2]int{
		{1, -2},
		{2, -1},
		{-1, -2},
		{-2, -1},
		{-2, 1},
		{-1, 2},
		{1, 2},
		{2, 1},
	}

	var (
		dp     func(x, y, z int) int
		result = 0
	)

	mem := [4][3][5001]int{}

	dp = func(x, y, z int) int {
		if y >= len(pattern) || x >= len(pattern[0]) || x < 0 || y < 0 || pattern[y][x] == -1 {
			return 0
		}

		if v := mem[y][x][z]; v != 0 {
			return v
		}

		if z == n {
			return 1
		}

		zz := z + 1
		result := 0

		for _, dir := range dirs {
			xx, yy := x+dir[0], y+dir[1]
			result = (result + dp(xx, yy, zz)) % 1000000007
		}

		mem[y][x][z] = result

		return mem[y][x][z]
	}

	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			if pattern[i][j] != -1 {
				result = (result + dp(j, i, 1)) % 1000000007
			}
		}
	}

	return result
}

// THIS TLE AT HIGH VALUES
func knightDialer_(n int) int {
	pattern := [4][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{-1, 0, -1},
	}

	dirs := [8][2]int{
		{1, -2},
		{2, -1},
		{-1, -2},
		{-2, -1},
		{-2, 1},
		{-1, 2},
		{1, 2},
		{2, 1},
	}

	var (
		dp     func(x, y, z int, memo map[[3]int]int) int
		result = 0
	)

	dp = func(x, y, z int, memo map[[3]int]int) int {
		if y >= len(pattern) || x >= len(pattern[0]) || x < 0 || y < 0 || pattern[y][x] == -1 {
			return 0
		}

		key := [3]int{x, y, z}

		if v, ok := memo[key]; ok {
			return v
		}

		if z == n {
			return 1
		}

		zz := z + 1
		result := 0

		for _, dir := range dirs {
			xx, yy := x+dir[0], y+dir[1]
			result = (result + dp(xx, yy, zz, memo)) % 1000000007
		}

		memo[key] = result

		return memo[key]
	}

	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			if pattern[i][j] != -1 {
				result = (result + dp(j, i, 1, map[[3]int]int{})) % 1000000007
			}
		}
	}

	return result
}

package dynamicProgramming

func numberOfWays1(s string) int64 {
	var (
		memo = make([][3][4]int64, len(s))
		dp   func(prev rune, curr int, count int, memo [][3][4]int64) int64
	)

	for i := 0; i < len(s); i++ {
		memo[i] = [3][4]int64{}

		for j := 0; j < 3; j++ {
			memo[i][j] = [4]int64{-1, -1, -1, -1}
		}
	}

	dp = func(prev rune, curr int, count int, memo [][3][4]int64) int64 {
		if count == 3 {
			return 1
		}

		if len(s) == curr {
			return 0
		}

		p := 2
		if prev == '0' {
			p = 0
		} else if prev == '1' {
			p = 1
		}

		if v := memo[curr][p][count]; v != -1 {
			return v
		}

		dont, do := dp(prev, curr+1, count, memo), int64(0)

		if prev == '2' || rune(s[curr]) != prev {
			do = dp(rune(s[curr]), curr+1, count+1, memo)
		}

		memo[curr][p][count] = do + dont

		return do + dont
	}

	return dp('2', 0, 0, memo)
}

// TLE
func numberOfWays__(s string) int64 {
	var dp func(prev int, curr int, count int, memo map[[3]int]int64) int64

	dp = func(prev int, curr int, count int, memo map[[3]int]int64) int64 {
		if count == 3 {
			return 1
		}

		if len(s) == curr {
			return 0
		}

		key := [3]int{prev, curr, count}
		if v, ok := memo[key]; ok {
			return v
		}

		dont, do := dp(prev, curr+1, count, memo), int64(0)

		count += 1

		if prev == -1 || s[curr] != s[prev] {
			do = dp(curr, curr+1, count, memo)
		}

		memo[key] = do + dont

		return memo[key]
	}

	return dp(-1, 0, 0, map[[3]int]int64{})
}

// func numberOfWays(s string) int64 {
//     var dp func(prev int, curr int, count int, memo map[[3]int]int64) int64

//     dp = func(prev int, curr int, count int, memo map[[3]int]int64) int64 {
//         if count == 3 {
//             return 1
//         }

//         if len(s) == curr {
//             return 0
//         }

//         key := [3]int {prev, curr, count}
//         if v, ok := memo[key]; ok {
//             return v
//         }

//         dont, do := dp(prev, curr + 1, count, memo), int64(0)

//         count += 1

//         if prev == -1 || s[curr] != s[prev] {
//             do = dp(curr, curr + 1, count, memo)
//         }

//         memo[key] = int64(do + dont)

//         return memo[key]
//     }

//     memo := make([][][3]int, len(s))

//     for i := 0; i < len(s); i++ {
//         memo[i] = make([][3]int, len(s))

//         for j := 0; j < len(s); j++ {
//             memo[i][j] = [3]int{-1, -1, -1}
//         }
//     }

//     return int64(dp(-1, 0, 0, map[[3]int]int64{}))
// }

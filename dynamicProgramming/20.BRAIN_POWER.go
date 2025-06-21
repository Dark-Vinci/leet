package dynamicProgramming

func mostPoints(questions [][]int) int64 {
	var helper func(i int, memo map[int]int64) int64

	helper = func(i int, memo map[int]int64) int64 {
		if i >= len(questions) {
			return int64(0)
		}

		if v, ok := memo[i]; ok {
			return v
		}

		point, brain := questions[i][0], questions[i][1]

		do, dont := int64(point)+helper(i+brain+1, memo), helper(i+1, memo)

		memo[i] = max(do, dont)

		return memo[i]
	}

	return helper(0, map[int]int64{})
}

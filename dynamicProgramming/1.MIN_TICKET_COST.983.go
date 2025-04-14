package dynamicProgramming

func mincostTickets(days []int, costs []int) int {
	var (
		daysMap = make(map[int]struct{})
		helper  func(day int, memo map[int]int) int
		l       = len(days)
	)

	for _, v := range days {
		daysMap[v] = struct{}{}
	}

	helper = func(day int, memo map[int]int) int {
		if day > days[l-1] {
			return 0
		}

		if v, ok := memo[day]; ok {
			return v
		}

		if _, ok := daysMap[day]; !ok {
			memo[day] = helper(day+1, memo)
			return memo[day]
		}

		cost0 := helper(day+1, memo) + costs[0]
		cost1 := helper(day+7, memo) + costs[1]
		cost2 := helper(day+30, memo) + costs[2]

		memo[day] = min(cost0, cost1, cost2)

		return memo[day]
	}

	return helper(1, map[int]int{})
}

func mincostTickets_(days []int, costs []int) int {
	var (
		dp      = make([]int, 366)
		daysMap = make(map[int]struct{})
	)

	for _, v := range days {
		daysMap[v] = struct{}{}
	}

	for i := 1; i < 366; i++ {
		if _, ok := daysMap[i]; !ok {
			dp[i] = dp[i-1]
			continue
		}

		dp[i] = min(
			dp[i-1]+costs[0],
			dp[max(0, i-7)]+costs[1],
			dp[max(0, i-30)]+costs[2],
		)
	}

	return dp[365]
}

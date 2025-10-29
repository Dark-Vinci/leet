package arrays

func goodDaysToRobBank(security []int, time int) []int {
	var (
		n      = len(security)
		result = make([]int, 0)
		left   = make([]int, n)
		right  = make([]int, n)
	)

	if time == 0 {
		for i := 0; i < n; i++ {
			result = append(result, i)
		}

		return result
	}

	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			result = append(result, i)
		}
	}

	return result
}

// TLE
func goodDaysToRobBank1(security []int, time int) []int {
	var (
		dp     func(curr int) bool
		result = make([]int, 0)
	)

	dp = func(curr int) bool {
		left, right := 0, 0

		for i := 1; i <= curr; i++ {
			if security[i-1] >= security[i] {
				left++
			}
		}

		for i := curr; i < len(security)-1; i++ {
			if security[i] <= security[i+1] {
				right++
			}
		}

		if right >= time && left >= time {
			return true
		}

		return false
	}

	for i := time; i < len(security)-time; i++ {
		if dp(i) {
			result = append(result, i)
		}
	}

	return result
}

package main

func waysToSplit(nums []int) int {
	n, result := len(nums), 0

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for l := 0; l < n-2; l++ {
		upper := searchT(nums, l, n-1, true)
		lower := searchT(nums, l, upper, false)

		if upper > lower {
			result = (result + upper - lower) % 1000000007
		}
	}

	return result
}

func searchT(nums []int, i, j int, upper bool) int {
	l, r := i+1, j

	for l < r {
		m := l + (r-l)/2
		mid := nums[m] - nums[i]

		if upper {
			right := nums[j] - nums[m]

			if mid <= right {
				l = m + 1
			} else {
				r = m
			}
		} else {
			if mid < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
	}

	return l
}

// TLE
func waysToSplit___(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]

	le := len(nums) - 1

	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	result, end := 0, len(nums)

	for l := 0; l <= end-2; l++ {
		for r := l + 1; r <= end-1; r++ {
			lSum := sums[l]
			mSum := sums[r] - sums[l]
			rSum := sums[le] - sums[r]

			if lSum <= mSum && mSum <= rSum {
				result++
			}
		}
	}

	return result
}

// TLE
func waysToSplit_(nums []int) int {
	result, end := 0, len(nums)

	for l := 1; l <= end-2; l++ {
		for r := l + 1; r <= end-1; r++ {
			lSum, mSum, rSum := 0, 0, 0

			for _, v := range nums[0:l] {
				lSum += v
			}

			for _, v := range nums[l:r] {
				mSum += v
			}

			for _, v := range nums[r:end] {
				rSum += v
			}

			if lSum <= mSum && mSum <= rSum {
				result++
			}
		}
	}

	return result
}

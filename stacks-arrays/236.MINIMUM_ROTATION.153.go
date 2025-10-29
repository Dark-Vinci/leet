package arrays

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[r]
}

func findMin_recursive(nums []int) int {
	var recur func(s, e int) int

	recur = func(s, e int) int {
		if s == e {
			return nums[s]
		}

		m := s + (e-s)/2

		if nums[m] > nums[e] {
			return recur(m+1, e)
		}

		return recur(s, m)
	}

	return recur(0, len(nums)-1)
}

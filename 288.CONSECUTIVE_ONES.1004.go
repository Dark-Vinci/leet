package main

func longestOnes(nums []int, k int) int {
	l, result, zeros := 0, 0, 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeros += 1
		}

		for zeros > k {
			if nums[l] == 0 {
				zeros -= 1
			}

			l += 1
		}

		result = max(result, r-l+1)
	}

	return result
}

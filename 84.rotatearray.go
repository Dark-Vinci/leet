package main

func rotate(nums []int, k int) {
	numLen := len(nums)

	k %= numLen

	var rev func(i, j int)

	rev = func(i, j int) {
		for i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	rev(0, numLen-1)
	rev(0, k-1)
	rev(k, numLen-1)
}

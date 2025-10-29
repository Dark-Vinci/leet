package arrays

func moveZeroes(nums []int) {
	nNums := make([]int, 0)
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nNums = append(nNums, nums[i])
		} else {
			count++
		}
	}

	for i := 0; i < count; i++ {
		nNums = append(nNums, 0)
	}

	copy(nums, nNums)
}

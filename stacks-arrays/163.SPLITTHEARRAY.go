package arrays

func isPossibleToSplit(nums []int) bool {
	db := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		db[nums[i]] += 1

		if db[nums[i]] > 2 {
			return false
		}
	}

	return true
}

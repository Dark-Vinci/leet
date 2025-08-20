package combinatorics

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	result = append(result, []int{nums[0]})

	for j := 1; j < len(nums); j++ {
		updated := make([][]int, 0)

		for _, permutation := range result {
			for i := 0; i <= len(permutation); i++ {
				before := permutation[:i]
				after := permutation[i:]

				newPermutation := make([]int, 0)

				newPermutation = append(newPermutation, before...)
				newPermutation = append(newPermutation, nums[j])
				newPermutation = append(newPermutation, after...)

				updated = append(updated, newPermutation)
			}
		}

		result = updated
	}

	return result
}

package stacks_arrays

func canThreePartsEqualSum(arr []int) bool {
	sum := arr[0]

	for i := 1; i < len(arr); i++ {
		sum += arr[i]
	}

	if sum%3 != 0 {
		return false
	}

	track, count, div3 := 0, 0, sum/3

	for i := 0; i < len(arr); i++ {
		track += arr[i]

		if track == div3 {
			count++
			track = 0
		}
	}

	return count >= 3
}

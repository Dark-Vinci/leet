package stacks_arrays

func maxContainers(n int, w int, maxWeight int) int {
	var (
		result        = 0
		cell          = n * n
		currentWeight = 0
	)

	for i := 0; i < cell; i++ {
		currentWeight += w

		if currentWeight <= maxWeight {
			result++
		}
	}

	return result
}

package arrays

func dailyTemperatures(temperatures []int) []int {
	var (
		result = make([]int, len(temperatures))
		l      = len(temperatures)
	)

	result[l-1] = -1

	for j := l - 2; j >= 0; j-- {
		curr := temperatures[j]

		nextIndex := -1

		if curr == temperatures[j+1] {
			result[j] = result[j+1]
			continue
		}

		for i := j + 1; i < l; i++ {
			if temperatures[i] > curr {
				nextIndex = i
				break
			}
		}

		result[j] = nextIndex
	}

	for i := 0; i < len(result); i++ {
		if result[i] == -1 {
			result[i] = 0
		} else {
			result[i] = result[i] - i
		}
	}

	return result
}

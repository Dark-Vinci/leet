package arrays

func findPoisonedDuration(timeSeries []int, duration int) int {
	count := 0

	for i := 0; i < len(timeSeries); i++ {
		dur := timeSeries[i] + duration - 1

		if i == len(timeSeries)-1 {
			count += duration
			continue
		}

		if dur < timeSeries[i+1] {
			count += duration
		} else {
			count += timeSeries[i+1] - timeSeries[i]
		}
	}

	return count
}

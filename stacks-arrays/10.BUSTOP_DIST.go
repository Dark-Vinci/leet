package arrays

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination {
		return 0
	}

	loop := append(distance, distance...)

	if start > destination {
		start, destination = destination, start
	}

	f, b := 0, 0

	for i := start; i < destination; i++ {
		f += loop[i]
	}

	for i := destination; i < len(distance)+start; i++ {
		b += loop[i]
	}

	return min(f, b)
}

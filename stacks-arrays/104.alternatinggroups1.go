package stacks_arrays

func numberOfAlternatingGroups(colors []int) int {
	var (
		n     = len(colors)
		last  = n - 1
		count = 0
	)

	for i := 0; i < n; i++ {
		var (
			a = colors[i]
			b int
			c int
		)

		if i+1 > last {
			b = colors[i-last]
		} else {
			b = colors[i+1]
		}

		if i+2 > last {
			c = colors[i+1-last]
		} else {
			c = colors[i+2]
		}

		if b != a && b != c {
			count++
		}
	}

	return count
}

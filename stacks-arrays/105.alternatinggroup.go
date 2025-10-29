package arrays

// Preferred solution
func numberOfAlternatingGroups1(colors []int, k int) int {
	var (
		count  = 1
		i      = 1
		result = 0
	)

	colors = append(colors, colors[:k-1]...)

	for i < len(colors) {
		if colors[i] != colors[i-1] {
			count++
		} else {
			count = 1
		}

		if count >= k {
			result++
		}

		i++
	}

	return result
}

// TOO CRUDE [EXCEED TIME LIMIT]
func numberOfAlternatingGroups11(colors []int, k int) int {
	var (
		count = 0
		n     = len(colors)
		last  = n - 1
	)

O:
	for i := 0; i < n; i++ {
		a := []int{colors[i]}

		if i+k > last {
			// append from i + 1
			for j := i + 1; j <= last; j++ {
				a = append(a, colors[j])
			}

			// append after
			for j := 0; j < (i+k)-last-1; j++ {
				a = append(a, colors[j])
			}
		} else {
			for j := i + 1; j < i+k; j++ {
				a = append(a, colors[j])
			}
		}

		for j := 0; j < len(a); j++ {
			if j > 0 && j+1 < len(a) {
				if a[j-1] == a[j] || a[j] == a[j+1] {
					continue O
				}
			}
		}

		count++
	}

	return count
}

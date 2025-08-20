package Grid

func generateMatrix(n int) [][]int {
	var (
		left, right = 0, n - 1
		top, bot    = 0, n - 1
		val         = 1
		result      = make([][]int, n)
	)

	for j := 0; j < n; j++ {
		result[j] = make([]int, n)
	}

	for left <= right {
		// left to right
		for i := left; i < right+1; i++ {
			result[top][i] = val
			val++
		}
		top++

		// top to down
		for i := top; i < bot+1; i++ {
			result[i][right] = val
			val++
		}
		right--

		// left to right
		for i := right; i >= left; i-- {
			result[bot][i] = val
			val++
		}
		bot--

		// down to up
		for i := bot; i >= top; i-- {
			result[i][left] = val
			val++
		}
		left++
	}

	return result
}

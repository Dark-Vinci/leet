package stacks_arrays

func maxArea(height []int) int {
	left, right, result := 0, len(height)-1, 0

	for left < right {
		w := right - left
		h := min(height[left], height[right])
		water := w * h

		result = max(result, water)

		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}
	}

	return result
}

package Sorting

func mergeSort(arr []int) []int {
	var recurse func(a []int) []int

	recurse = func(a []int) []int {
		if len(a) <= 1 {
			return a
		}

		l, r := 0, len(a)

		mid := l + (r-l)/2

		left := recurse(a[:mid])
		right := recurse(a[mid:])

		result := merge(left, right)

		return result
	}

	return recurse(arr)
}

func merge(a, b []int) []int {
	result := make([]int, 0)

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		result = append(result, a[i])
	}

	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return result
}

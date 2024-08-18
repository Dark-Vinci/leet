package main

func nextPermutation(arr []int) {
	l := len(arr)

	i := l - 2

	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}

	if i >= 0 {
		for j := l - 1; j > i; j-- {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
				break
			}
		}
	}

	for j, k := i+1, l-1; j < k; j, k = j+1, k-1 {
		arr[j], arr[k] = arr[k], arr[j]
	}
}

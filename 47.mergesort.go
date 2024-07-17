package main

func quickSort(arr []int) []int {
	var recurse func(a []int) [] int
	
	recurse = func(a []int) []int {
		if len(a) <= 1 {
			return a
		}
		
		pivot := a[len(a) - 1]
		
		left := finds(a, pivot, true)
		
		sortedLeft := recurse(left)
		
		right := finds(a, pivot, false)
		
		sortedRight := recurse(right)
		
		res := make([]int, 0)
		
		res = append(res, sortedLeft...)
		res = append(res, pivot)
		res = append(res, sortedRight...)
		
		return res
	}
	
	return recurse(arr)
}

func finds(arr []int, pivot int, min bool) []int {
	result := make([]int, 0)
	
	for _, j := range arr {
		if min && j < pivot {
			result = append(result, j)
		} else if !min && j > pivot {
			result = append(result, j)
		}
	}
	
	
	return result
}

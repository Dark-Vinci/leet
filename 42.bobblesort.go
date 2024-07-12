package main

func bobbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	var recurse func(a []int, i, j, end int)
	
	recurse = func(a []int, i, j, end int) {
		if i == end {
			return
		}
		
		pOne := i + 1
		
		if pOne < end {
			if a[i] > a[pOne] {
				a[i], a[pOne] = a[pOne], a[i]
				i++
				
				recurse(a, i, 0, end)
			}
		}
	}
	
	recurse(arr, 0, 0, len(arr) - 1)
	
	return arr
}
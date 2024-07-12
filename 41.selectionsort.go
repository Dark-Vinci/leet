package main

import (
	"math"
)

func selectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	var recurse func(arr []int, end int)
	
	recurse = func(arr []int, end int) {
		if end <= 0 {
			return
		}
		
		l := largest(arr[:end+1])
		
		arr[end], arr[l] = arr[l], arr[end]
		
		end--
		
		recurse(arr, end)
	}
	
	recurse(arr, len(arr) - 1)
	
	return arr
}

func largest(arr []int) int {
	larg := math.MinInt
	index := -1
	
	for k, v := range arr {
		if v > larg {
			larg = v
			index = k
		}
	}
	
	return index
}

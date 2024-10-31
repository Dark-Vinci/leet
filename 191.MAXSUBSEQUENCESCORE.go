package main

import (
	"container/heap"
	"math"
	"slices"
	"sort"
)

// Pair SORT
type Pair struct {
	n1, n2 []int
}

func (p Pair) Len() int {
	return len(p.n1)
}

func (p Pair) Swap(i, j int) {
	p.n1[i], p.n1[j] = p.n1[j], p.n1[i]
	p.n2[i], p.n2[j] = p.n2[j], p.n2[i]
}

func (p Pair) Less(i, j int) bool {
	return p.n2[i] > p.n2[j]
}

// HEAP
type myHeap []int

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m myHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *myHeap) Push(val any) {
	*m = append(*m, val.(int))
}

func (m *myHeap) Pop() any {
	n := len(*m)
	c := (*m)[n-1]
	*m = (*m)[:n-1]

	return c
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	sort.Sort(Pair{nums1, nums2})

	h := &myHeap{}

	maxSum := 0

	for i := 0; i < k; i++ {
		heap.Push(h, nums1[i])
		maxSum += nums1[i]
	}

	mx := maxSum * nums2[k-1]

	for i := k; i < len(nums1); i++ {
		maxSum -= heap.Pop(h).(int)
		maxSum += nums1[i]
		heap.Push(h, nums1[i])
		mx = max(mx, maxSum*nums2[i])
	}

	return int64(mx)
}

// EXCEED TIME LIMIT
func maxScore0(nums1 []int, nums2 []int, k int) int64 {
	comb, result := make([][]int, 0), int64(0)

	var dfs func(top []int, i int)
	dfs = func(top []int, i int) {
		if len(top) > k {
			return
		}

		if i == len(nums1) {
			if len(top) == k {
				t := slices.Clone(top)
				comb = append(comb, t)
			}

			return
		}

		dfs(top, i+1)

		top = append(top, i)

		dfs(top, i+1)
	}

	dfs([]int{}, 0)

	helper := func(arr []int) int64 {
		n1, n2 := 0, math.MaxInt

		for i := 0; i < len(arr); i++ {
			n1 += nums1[arr[i]]
			n2 = min(n2, nums2[arr[i]])
		}

		return int64(n1 * n2)
	}

	for i := 0; i < len(comb); i++ {
		result = max(result, helper(comb[i]))
	}

	return result
}

package main

import "container/heap"

type nyHeap [][2]int

func (m nyHeap) Len() int {
	return len(m)
}

func (m nyHeap) Less(i, j int) bool {
	return m[i][0] < m[j][1]
}

func (m nyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *nyHeap) Push(x any) {
	*m = append(*m, x.([2]int))
}

func (m *nyHeap) Pop() any {
	l := len(*m)

	x := (*m)[l-1]
	*m = (*m)[:l-1]
	return x
}

func shortestPath(l int, edges [][3]int, source int) map[int]int {
	adj := make(map[int][][2]int)

	for i := range l {
		adj[i] = [][2]int{}
	}

	for _, edge := range edges {
		s, d, w := edge[0], edge[1], edge[2]
		adj[s] = append(adj[s], [2]int{d, w})
	}

	result := make(map[int]int)

	hh := nyHeap{{0, source}}
	heap.Init(&hh)

	for len(hh) > 0 {
		curr := heap.Pop(&hh).([2]int)
		w, n := curr[0], curr[1]

		if _, ok := result[n]; ok {
			continue
		}

		result[n] = w

		for _, v := range adj[n] {
			if _, ok := result[v[1]]; !ok {
				heap.Push(&hh, [2]int{v[0] + w, v[1]})
			}
		}
	}

	for i := range l {
		if _, ok := result[i]; !ok {
			result[i] = -1
		}
	}

	return result
}

package main

import "container/heap"

type pq [][2]int // number; freq

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i][1] > p[j][1]
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(val any) {
	*p = append(*p, val.([2]int))
}

func (p *pq) Pop() any {
	val := (*p)[len(*p)-1]

	*p = (*p)[:len(*p)-1]

	return val
}

func rearrangeBarcodes(barcodes []int) []int {
	var (
		prev = [2]int{0, -1}
		db   = make(map[int]int)
		p    = new(pq)
	)

	heap.Init(p)

	for _, val := range barcodes {
		db[val] += 1
	}

	for key, value := range db {
		heap.Push(p, [2]int{key, value})
	}

	for i := range len(barcodes) {
		curr := heap.Pop(p).([2]int)

		if prev[1] != -1 {
			heap.Push(p, prev)
		}

		barcodes[i] = curr[0]
		curr[1]--
		prev = curr
	}

	return barcodes
}

package Graphs

import (
	"container/heap"
	"math"
)

type item struct {
	node, distance, index int
}

type PQ []*item

func (p PQ) Len() int {
	return len(p)
}

func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index, p[j].index = p[j].index, p[i].index
}

func (p PQ) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p *PQ) Push(val any) {
	l := len(*p)

	value := val.(*item)
	value.index = l

	*p = append(*p, value)
}

func (p *PQ) Pop() any {
	old := *p

	l := len(old)

	last := old[l-1]

	old[l-1] = nil
	last.index = -1

	*p = old[:l-1]

	return last
}

func dijk(val map[int][][2]int, start int) map[int]int {
	result := make(map[int]int)

	for i := range val {
		result[i] = math.MaxInt
	}

	result[start] = 0

	startNode := &item{node: start, distance: 0}

	h := &PQ{startNode}

	heap.Init(h)

	for h.Len() > 0 {
		mn := heap.Pop(h).(*item)

		n, cost := mn.node, mn.distance

		if cost > result[n] {
			continue
		}

		for _, chil := range val[n] {
			newNode, newNodeWeight := chil[0], chil[1]

			if cost+newNodeWeight < result[newNode] {
				result[newNode] = cost + newNodeWeight

				heap.Push(h, &item{node: newNode, distance: result[newNode]})
			}
		}

	}

	return result
}

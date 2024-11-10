package main

import (
	"container/heap"
	"math"
)

type Edge struct {
	node, weight int
}

type Graph struct {
	edges map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(u, v, weight int) {
	g.edges[u] = append(g.edges[u], Edge{node: v, weight: weight})
}

type Item struct {
	index, node, distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(val any) {
	item := val.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]

	return item
}

func (pq *PriorityQueue) update(item *Item, dist int) {
	item.distance = dist
	heap.Fix(pq, item.index)
}

func Djks(g *Graph, start int) map[int]int {
	result := map[int]int{}

	for nod := range g.edges {
		result[nod] = math.MaxInt32
	}

	result[start] = 0

	pq := PriorityQueue{{-1, start, 0}}

	heap.Init(&pq)

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)
		currNode, currDist := curr.node, curr.distance

		if currDist > result[currNode] {
			continue
		}

		for _, edge := range g.edges[currNode] {
			newDist := currDist + edge.weight

			if newDist < result[edge.node] {
				result[edge.node] = newDist
				heap.Push(&pq, &Item{node: edge.node, distance: newDist})
			}
		}
	}

	return result
}

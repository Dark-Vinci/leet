package dynamicProgramming

import (
	"container/heap"
	"math"
)

type Item struct {
	effort int
	i, j   int
}

type PriorityQueue []Item

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].effort < pq[j].effort
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(a any) {
	*pq = append(*pq, a.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	last := old[n-1]
	*pq = old[:n-1]

	return last
}

func minimumEffortPath(heights [][]int) int {
	l1, l2 := len(heights), len(heights[0])

	var (
		efforts    = make([][]int, l1)
		directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		pq         = make(PriorityQueue, 0)
	)

	for i := range efforts {
		efforts[i] = make([]int, l2)

		for j := range efforts[i] {
			efforts[i][j] = math.MaxInt
		}
	}

	efforts[0][0] = 0

	heap.Init(&pq)
	heap.Push(&pq, Item{effort: 0, i: 0, j: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(Item)

		i, j, effort := current.i, current.j, current.effort

		if i == l1-1 && j == l2-1 {
			return effort
		}

		for _, dir := range directions {
			ii, jj := i+dir[0], j+dir[1]

			if ii >= 0 && jj >= 0 && ii < l1 && jj < l2 {
				absDiff := abs(heights[ii][jj] - heights[i][j])

				eff := max(absDiff, effort)

				if eff < efforts[ii][jj] {
					heap.Push(&pq, Item{effort: eff, i: ii, j: jj})
					efforts[ii][jj] = eff
				}
			}
		}
	}

	return efforts[l1-1][l2-1]
}

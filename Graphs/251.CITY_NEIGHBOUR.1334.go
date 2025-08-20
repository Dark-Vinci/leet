package Graphs

import (
	"cmp"
	"container/heap"
	"math"
	"slices"
)

type nnode struct {
	to, cost int
}

type Hp []nnode

func (h *Hp) Push(x any) {
	*h = append(*h, x.(nnode))
}

func (h *Hp) Pop() any {
	old := *h
	l := len(old)
	result := old[l-1]

	*h = old[:l-1]

	return result
}

func (h Hp) Len() int {
	return len(h)
}

func (h Hp) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h Hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	var (
		list   = toThisAdjList(edges)
		djk    func(from int) int
		result = make([]int, 0)
		mn     = math.MaxInt
	)

	djk = func(from int) int {
		var (
			start  = nnode{to: from, cost: 0}
			mp     = make(map[int]int)
			hp     = &Hp{start}
			result = 0
		)

		heap.Init(hp)

		for i := 0; i < n; i++ {
			mp[i] = math.MaxInt
		}

		mp[from] = 0

		for hp.Len() > 0 {
			c := heap.Pop(hp)
			current := c.(nnode)

			if current.cost > mp[current.to] {
				continue
			}

			for _, v := range list[current.to] {
				if v.cost+current.cost < mp[v.to] {
					mp[v.to] = v.cost + current.cost

					n := nnode{cost: mp[v.to], to: v.to}

					heap.Push(hp, n)
				}
			}
		}

		for _, val := range mp {
			if val <= distanceThreshold {
				result++
			}
		}

		return result
	}

	for i := range n {
		res := djk(i)

		if res < mn {
			mn = res
			result = []int{i}
			continue
		}

		if res == mn {
			result = append(result, i)
		}
	}

	slices.SortFunc(result, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return result[0]
}

func toThisAdjList(edges [][]int) map[int][]nnode {
	result := make(map[int][]nnode)

	for _, val := range edges {
		u, v, w := val[0], val[1], val[2]
		a, b := nnode{to: v, cost: w}, nnode{to: u, cost: w}

		result[u] = append(result[u], a)
		result[v] = append(result[v], b)
	}

	return result
}

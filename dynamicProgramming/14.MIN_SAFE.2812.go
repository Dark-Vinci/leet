package dynamicProgramming

import "container/heap"

type Jtem struct {
	dist int
	i, j   int
}

type MaxHeap []Jtem

func (pq MaxHeap) Less(i, j int) bool {
	return pq[i].dist > pq[j].dist
}

func (pq MaxHeap) Len() int {
	return len(pq)
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxHeap) Push(a any) {
	*pq = append(*pq, a.(Jtem))
}

func (pq *MaxHeap) Pop() any {
	old := *pq
	n := len(old)
	last := old[n-1]
	*pq = old[:n-1]

	return last
}


func maximumSafenessFactor(grid [][]int) int {
    var (
        l = len(grid)
        prelim func() map[[2]int] int
    )

    prelim = func() map[[2]int] int {
        var (
	       	q = [][3]int{}
	        distMap = make(map[[2]int] int)
        )

        for i := range l {
            for j := range l {
                if grid[i][j] == 1 {
                    q = append(q, [3]int{i, j, 0})
                    distMap[[2]int{i, j}] = 0
                }
            }
        }

        for len(q) > 0 {
            ll := len(q)

            for i := 0; i < ll; i++ {
                curr := q[0]
                q = q[1:]

                ii, jj, dist := curr[0], curr[1], curr[2]

                dirs := [][]int {
                    {ii, jj + 1},
                    {ii, jj - 1},
                    {ii + 1, jj},
                    {ii - 1, jj},
                }

                for _, val := range dirs {
                    x, y := val[0], val[1]
                    pos := [2]int{x, y}

                    if x >= 0 && x < l && y >= 0 && y < l {
                        if _, ok := distMap[pos]; !ok {
                            q = append(q, [3]int{x, y, dist + 1})
                            distMap[[2]int{x, y}] = dist + 1
                        }
                    }
                }
            }
        }

        return distMap
    }

    var (
        mp = prelim()
        mh = make(MaxHeap, 0)
        visited = make(map[[2]int]struct{})
    )

    heap.Init(&mh)
    heap.Push(&mh, Jtem{i: 0, j: 0, dist: mp[[2]int{0, 0}]})

    for mh.Len() > 0 {
        curr := heap.Pop(&mh).(Jtem)

        dist, i, j := curr.dist, curr.i, curr.j

        if i == l - 1 && j == l - 1 {
            return dist
        }

        dirs := [][]int {
            {i, j + 1},
            {i, j - 1},
            {i + 1, j},
            {i - 1, j},
        }

        for _, val := range dirs {
            x, y := val[0], val[1]

            if x >= 0 && y >= 0 && x < l && y < l {
                pos := [2]int{x, y}

                if _, ok := visited[pos]; !ok {
                    visited[pos] = struct{}{}

                    newDist := min(dist, mp[pos])

                    heap.Push(&mh, Jtem{i: x, j: y, dist: newDist})
                }
            }
        }
    }

    return -1
}

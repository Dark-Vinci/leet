package main

func findCircleNum(isConnected [][]int) int {

	var (
		visited = make(map[int]struct{})
		result  = 0
		adjList = make(map[int][]int)
		dfs     func(to int)
	)

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if i != j && isConnected[i][j] != 0 {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	dfs = func(to int) {
		if _, ok := visited[to]; ok {
			return
		}

		visited[to] = struct{}{}

		for i := 0; i < len(adjList[to]); i++ {
			dfs(adjList[to][i])
		}
	}

	for key := range len(isConnected) {
		if _, ok := visited[key]; !ok {
			result++
			dfs(key)
		}
	}

	return result
}

type UnionFind struct {
	par  []int
	rank []int
}

func New(n int) *UnionFind {
	p, r := make([]int, n), make([]int, n)

	for i := range n {
		p[i] = i
	}

	return &UnionFind{par: p, rank: r}
}

func (u *UnionFind) Find(a int) int {
	if u.par[a] != a {
		u.par[a] = u.Find(u.par[a])
	}

	return u.par[a]
}

func (u *UnionFind) Union(a, b int) {
	rootA, rootB := u.Find(a), u.Find(b)

	if rootA != rootB {
		rankA, rankB := u.rank[rootA], u.rank[rootB]

		if rankA > rankB {
			u.par[rootB] = rootA
		} else if rankA < rankB {
			u.par[rootA] = rootB
		} else {
			u.par[rootB] = rootA
			u.rank[rootA]++
		}
	}
}

func findCircleNumUnionFind(isConnected [][]int) int {
	var (
		uf  = New(len(isConnected))
		set = make(map[int]struct{})
	)

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if i != j && isConnected[i][j] != 0 {
				uf.Union(i, j)
			}
		}
	}

	for i := 0; i < len(uf.par); i++ {
		root := uf.Find(i)
		set[root] = struct{}{}
	}

	return len(set)
}

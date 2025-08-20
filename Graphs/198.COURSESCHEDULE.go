package Graphs

func courseSchedules(count int, prerequisites [][]int) bool {
	var (
		visit   = make(map[int]struct{})
		dfs     func(val int) bool
		adjList = make(map[int][]int)
	)
	for i := range count {
		adjList[i] = make([]int, 0)
	}
	for _, val := range prerequisites {
		a, b := val[0], val[1]
		adjList[a] = append(adjList[a], b)
	}
	dfs = func(val int) bool {
		if _, ok := visit[val]; ok {
			return false
		}
		if len(adjList[val]) == 0 {
			return true
		}
		visit[val] = struct{}{}
		curr := adjList[val]
		for _, v := range curr {
			if !dfs(v) {
				return false
			}
		}
		delete(visit, val)
		adjList[val] = []int{}
		return true
	}
	for i := range count {
		if !dfs(i) {
			return false
		}
	}
	return true
}

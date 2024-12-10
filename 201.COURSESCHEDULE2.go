package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		list     = make(map[int][]int)
		visited  = make(map[int]struct{})
		visiting = make(map[int]struct{})
		dfs      func(int) bool
		result   = make([]int, 0)
	)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		list[a] = append(list[a], b)
	}
	dfs = func(val int) bool {
		if _, ok := visiting[val]; ok {
			return false
		}
		if _, ok := visited[val]; ok {
			return true
		}
		visiting[val] = struct{}{}
		visited[val] = struct{}{}
		for i := 0; i < len(list[val]); i++ {
			if !dfs(list[val][i]) {
				return false
			}
		}
		result = append(result, val)
		delete(visiting, val)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if _, ok := visited[i]; !ok && !dfs(i) {
			return []int{}
		}
	}
	return result
}

func main() {
	arr := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	fmt.Println(findOrder(4, arr))
}

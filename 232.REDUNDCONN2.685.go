package main

import (
	"fmt"
	"slices"
)

func findRedundantDirectedConnection(edges [][]int) []int {
	var dfs func(lst map[int][]int, visited map[int]struct{}, curr int, count *int) bool
	n := len(edges)
	result := []int{}

	dfs = func(lst map[int][]int, visited map[int]struct{}, curr int, count *int) bool {
		if _, ok := visited[curr]; ok {
			return false
		}

		visited[curr] = struct{}{}
		*count += 1

		for _, conn := range lst[curr] {
			if !dfs(lst, visited, conn, count) {
				return false
			}
		}

		delete(visited, curr)

		return true
	}

	for i := n - 1; i >= 0; i-- {
		curr := edges[i]
		edges = slices.Delete(edges, i, i+1)
		list, start, valid := thisAdjList(edges, n)

		fmt.Println("LIST, START, VALID", list, start, valid, edges)

		if !valid {
			edges = slices.Insert(edges, i, curr)
			continue
		}

		visited, count := make(map[int]struct{}), 0

		if dfs(list, visited, start, &count) && count == n {
			result = curr
			break
		}

		edges = slices.Insert(edges, i, curr)

		// fmt.Println("COUNT, I", count, i, edges)
	}

	return result
}

func thisAdjList(edges [][]int, n int) (map[int][]int, int, bool) {
	graph, start := make(map[int][]int), 0
	to, valid := make(map[int]struct{}), true

	for _, edge := range edges {
		u, v := edge[0], edge[1] // from -> to

		if _, ok := to[v]; ok {
			return map[int][]int{}, int(0), false
		}

		to[v] = struct{}{}

		graph[u] = append(graph[u], v)
	}

	for i := range n {
		if _, ok := to[i+1]; !ok {
			start = i + 1
		}
	}

	for _, value := range graph {
		if len(value) > 1 {
			count := 0
			for _, v := range value {
				if _, ok := graph[v]; ok {
					count++
				}
			}

			if count > 1 {
				valid = false
			}
		}
	}

	return graph, start, valid
}

package main

import (
	"cmp"
	"fmt"
)

//func pacificAtlantic(heights [][]int) [][]int {
//	result := make([][]int, 0)
//	l1, l2 := len(heights), len(heights[0])
//
//	var check func(i, j int) bool
//
//	check = func(i, j int) bool {
//		r := true
//		for k := i; k < l2; k++ {
//			if k+1 == l2 {
//				break
//			}
//
//			if heights[j][k] < heights[j][k+1] {
//				r = false
//				break
//			}
//		}
//
//		d := true
//
//		for k := j; k < l1; k++ {
//			if k+1 == l1 {
//				break
//			}
//
//			if heights[k][i] < heights[k+1][i] {
//				d = false
//				break
//			}
//		}
//
//		if !r && !d {
//			return false
//		}
//
//		l := true
//		for k := i; k >= 0; k-- {
//			if k-1 < 0 {
//				break
//			}
//
//			if heights[j][k] < heights[j][k-1] {
//				l = false
//				break
//			}
//		}
//
//		u := true
//		for k := j; k >= 0; k-- {
//			if k-1 < 0 {
//				break
//			}
//
//			if heights[k][i] < heights[k-1][i] {
//				u = false
//				break
//			}
//		}
//
//		return cmp.Or(u, l)
//	}
//
//	for j := 0; j < l1; j++ {
//		for i := 0; i < l2; i++ {
//			if check(i, j) {
//				result = append(result, []int{j, i})
//			}
//		}
//	}
//
//	return result
//}

//func pacificAtlantic(heights [][]int) [][]int {
//	result := make([][]int, 0)
//	l1, l2 := len(heights), len(heights[0])
//	visited := make(map[[2]int]struct{})
//
//	var dfs func(dir *int, i, j int) bool
//
//	dfs = func(dir *int, i, j int) bool {
//		n := [2]int{i, j}
//		if _, ok := visited[n]; ok {
//			return false
//		}
//
//		if i >= l2 || j >= l1 || j < 0 || i < 0 {
//			return true
//		}
//
//		current := heights[j][i]
//
//		var r, d, u, l bool
//
//		visited[n] = struct{}{}
//
//		if i+1 == l2 {
//			r = true
//		} else if i+1 < l2 && heights[j][i+1] <= current {
//			nDir := 1
//			r = dfs(&nDir, i+1, j)
//			if dir != nil && *dir == 1 {
//				return r
//			}
//		}
//
//		if j+1 == l1 {
//			d = true
//		} else if heights[j+1][i] <= current {
//			nDir := 2
//			d = dfs(&nDir, i, j+1)
//
//			if dir != nil && *dir == 2 {
//				return d
//			}
//		}
//
//		if !r && !d {
//			delete(visited, n)
//			return false
//		}
//
//		if i-1 < 0 {
//			l = true
//		} else if heights[j][i-1] <= current {
//			l = dfs(i-1, j)
//
//			if dir != nil && *dir == 2 {
//				return d
//			}
//		}
//
//		if j-1 < 0 {
//			u = true
//		} else if heights[j-1][i] <= current {
//			u = dfs(i, j-1)
//
//			if dir != nil && *dir == 2 {
//				return d
//			}
//		}
//
//		delete(visited, n)
//
//		return cmp.Or(l, u)
//	}
//
//	for j := 0; j < l1; j++ {
//		for i := 0; i < l2; i++ {
//			clear(visited)
//			fmt.Println(i, j)
//			if dfs(i, j) {
//				result = append(result, []int{j, i})
//			}
//		}
//	}
//
//	return result
//}

//func

func pacificAtlanticjdjdj(heights [][]int) [][]int {
	result := make([][]int, 0)
	l1, l2 := len(heights), len(heights[0])
	cache := make(map[[2]int]bool)
	visited := make(map[[2]int]struct{})

	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		if i == l2-1 && j == 0 {
			return true
		}

		if i == 0 && j == l1-1 {
			return true
		}

		current := [2]int{i, j}

		if v, ok := cache[current]; ok {
			return v
		}

		if _, ok := visited[current]; ok {
			return false
		}

		var left, right, up, down bool

		// RIGHT
		if i+1 >= l1 {
			right = true
		} else if heights[j][i] >= heights[j][i+1] {
			right = dfs(i+1, j)
		}

		// DOWN
		if j+1 >= l2 {
			down = true
		} else if heights[j][i] >= heights[j+1][i] {
			down = dfs(i, j+1)
		}

		// UP
		if j-1 < 0 {
			up = true
		} else if heights[j][i] >= heights[j-1][i] {
			up = dfs(i, j-1)
		}

		// LEFT
		if i-1 < 0 {
			left = true
		} else if heights[j][i] >= heights[j][i-1] {
			left = dfs(i-1, j)
		}

		res := cmp.Or(right && up, down && left)

		cache[current] = res
		delete(visited, current)

		return res
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			clear(visited)
			if dfs(i, j) {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func main() {
	a := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}

	res := pacificAtlantic(a)

	fmt.Println(res)
}

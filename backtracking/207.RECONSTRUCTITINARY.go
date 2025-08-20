package backtracking

import (
	"cmp"
	"slices"
)

// SOLUTION THAT WORKS
func findItinerary(tickets [][]string) []string {
	var (
		dfs    func(port string)
		list   = make(map[string][]string)
		result = make([]string, 0)
	)

	for _, val := range tickets {
		u, v := val[0], val[1]
		list[u] = append(list[u], v)
	}

	for key := range list {
		slices.SortFunc(list[key], func(a, b string) int {
			return cmp.Compare(b, a)
		})
	}

	dfs = func(port string) {
		for len(list[port]) > 0 {
			next := list[port][len(list[port])-1]
			list[port] = list[port][:len(list[port])-1]
			dfs(next)
		}

		result = append(result, port)
	}

	dfs("JFK")

	slices.Reverse(result)

	return result
}

// TEST CASE 80/81 WILL NOT PASS{ SO SAD }
func findItineraryE(tickets [][]string) []string {
	var (
		result      = []string{"JFK"}
		list, count = AdjList(tickets)
		times       = len(tickets)
		dfs         func(from string) bool
	)

	dfs = func(from string) bool {
		if len(result) == times+1 {
			return true
		}

		if _, ok := list[from]; !ok {
			return false
		}

		for _, to := range list[from] {
			if count[[2]string{from, to}] > 0 {
				count[[2]string{from, to}]--
				result = append(result, to)

				if dfs(to) {
					return true
				}

				result = result[:len(result)-1]
				count[[2]string{from, to}]++
			}
		}

		return false
	}

	dfs("JFK")

	return result
}

func AdjList(tickets [][]string) (map[string][]string, map[[2]string]int) {
	result, count := make(map[string][]string), make(map[[2]string]int)

	for _, val := range tickets {
		u, v := val[0], val[1]

		count[[2]string{u, v}]++

		result[u] = append(result[u], v)
	}

	for _, val := range result {
		slices.Sort(val)
	}

	return result, count
}

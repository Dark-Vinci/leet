package Graphs

import (
	"cmp"
	"slices"
)

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	var (
		result  = make([]string, 0)
		watches = make(map[string]int)
		used    = make(map[int]struct{})
		nodes   = []int{id}
	)

	used[id] = struct{}{}

	for len(nodes) > 0 && level >= 0 {
		l := len(nodes)

		for i := 0; i < l; i++ {
			conn := nodes[0]
			nodes = nodes[1:]

			if level == 0 {
				for _, watch := range watchedVideos[conn] {
					watches[watch] += 1
					result = append(result, watch)
				}

				continue
			}

			for _, v := range friends[conn] {
				if _, ok := used[v]; !ok {
					nodes = append(nodes, v)
					used[v] = struct{}{}
				}
			}
		}

		level--
	}

	slices.Sort(result)

	result = slices.Compact(result)

	slices.SortFunc(result, func(a, b string) int {
		if watches[a] == watches[b] {
			return cmp.Compare(a, b)
		}

		return cmp.Compare(watches[a], watches[b])
	})

	return result
}

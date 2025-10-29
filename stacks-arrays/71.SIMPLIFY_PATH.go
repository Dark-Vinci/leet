package arrays

import (
	"cmp"
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var (
		result  = make([]string, 0)
		pathArr = strings.Split(path, "/")
	)

	for i := 0; i < len(pathArr); i++ {
		curr := pathArr[i]

		if curr == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else if curr != "." && curr != "" {
			result = append(result, fmt.Sprintf("/%v", curr))
		}
	}

	return cmp.Or(strings.Join(result, ""), "/")
}

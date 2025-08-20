package Graphs

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	var (
		ingredientIndex = make(map[string]int)
		suppliedDB      = make(map[string]struct{})
		recipeGraph     = map[string][]string{}
		result          = make([]string, 0)
		dfs             func(ind int) bool
		visited         = make(map[int]bool)
		visiting        = make(map[int]struct{})
	)

	for _, val := range supplies {
		suppliedDB[val] = struct{}{}
	}

	for i, val := range recipes {
		recipeGraph[val] = []string{}
		ingredientIndex[val] = i
	}

	for i := 0; i < len(ingredients); i++ {
		for j := 0; j < len(ingredients[i]); j++ {
			ingr, recipee := ingredients[i][j], recipes[i]

			if _, ok := recipeGraph[ingr]; ok {
				recipeGraph[recipee] = append(recipeGraph[recipee], ingr)
			}
		}
	}

	dfs = func(ind int) bool {
		if res, ok := visited[ind]; ok {
			return res
		}

		if _, ok := visiting[ind]; ok {
			return false
		}

		visiting[ind] = struct{}{}

		for i := 0; i < len(ingredients[ind]); i++ {
			ingredient := ingredients[ind][i]

			if _, ok := suppliedDB[ingredient]; !ok {
				if _, ko := recipeGraph[ingredient]; !ko {
					visited[ind] = false
					return false
				}

				if !dfs(ingredientIndex[ingredient]) {
					visited[ind] = false
					return false
				}
			}
		}

		visited[ind] = true
		delete(visiting, ind)
		return true
	}

	for i, val := range recipes {
		if dfs(i) {
			result = append(result, val)
		}
	}

	return result
}

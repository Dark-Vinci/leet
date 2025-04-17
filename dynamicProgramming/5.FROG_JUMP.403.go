package dynamicProgramming

import (
	"cmp"
)

func canCross_(stones []int) bool {
	mp := make(map[int]struct{})
	
	dp := make([]bool, len(stones))
	ks := map[int]int {0: 1}
	
	for _, v := range stones {
		mp[v] = struct{}{}
	}
	
	dp[0] = true
	
	for i := 0; i < len(stones); i++ {
		if i == 0 {
			k, _ := ks[i]
			
			if _, ok := mp[i + k]; ok {
				dp[i + k] = true
				ks[i + k] = k
			}
			
			continue
		}
		
		if !dp[i] {
			continue
		}
		
		k, _ := ks[i] //
		
		if _, ok := mp[i + k]; ok {
			dp[i + k] = true
			ks[i + k] = k
		}
		
		if _, ok := mp[i + k + 1]; ok {
			dp[i + k + 1] = true
			ks[i + k + 1] = k + 1
		}
		
		if k - 1 > 0 {
			if _, ok := mp[i + k - 1]; ok {
				dp[i + k - 1] = true
				ks[i + k - 1] = k - 1
			}
		} 
	}
	
	return dp[len(stones) - 1]
}

func canCross(stones []int) bool {
    var (
        memo = make(map[[2]int]bool)
        l = len(stones)
        dfs func(i, k int) bool
        mp = make(map[int]struct{})
    )

    for _, v := range stones {
        mp[v] = struct{}{}
    }

    dfs = func(i, k int) bool {
        key := [2]int{i, k}

        if v, ok := memo[key]; ok {
            return v
        }

        if i > stones[l - 1] || i < stones[0] {
            memo[key] = false
            return false
        }

        if _, ok := mp[i]; !ok {
            memo[key] = false
            return false
        }

        if i == stones[l - 1] {
            memo[key] = true
            return true
        }

        if i == 0 {
            memo[key] = dfs(i + k, k) // 0 + 1, 1
            return memo[key]
        }

        if k - 1 > 0 {
            a := dfs(i + k - 1, k - 1)

            if a {
                memo[key] = a
                return memo[key]
            }
        }

        b, c := dfs(i + k, k), dfs(i + k + 1, k + 1)

        memo[key] = cmp.Or(b, c)

        return memo[key]
    }

    return dfs(0, 1)
}
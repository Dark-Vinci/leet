package dynamicProgramming

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA) // since len(energyDrinkA) == len(energyDrinkB)
	dp := make([][2]int64, n)

	dp[0][0] = int64(energyDrinkA[0])
	dp[0][1] = int64(energyDrinkB[0])

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0]+int64(energyDrinkA[i]), dp[i-1][1])
		dp[i][1] = max(dp[i-1][1]+int64(energyDrinkB[i]), dp[i-1][0])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

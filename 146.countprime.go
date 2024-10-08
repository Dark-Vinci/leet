package main

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	db := make([]bool, n)
	result := 0

	for i := 2; i < n; i++ {
		if !db[i] {
			result++

			for j := i * i; j < n; j += i {
				db[j] = true
			}
		}
	}

	return result
}

func countPrimes_EXCEED_TIME(n int) int {
	count := 0
	ind := 1

	isPrime := func(n int) bool {
		if n == 1 {
			return false
		}

		if n <= 3 {
			return true
		}

		if n%2 == 0 || n%3 == 0 {
			return false
		}

		i := 5

		for i != n {
			if n%i == 0 {
				return false
			}

			i++
		}

		return true
	}

	for ind < n {
		if isPrime(ind) {
			count++
		}

		ind++
	}

	return count
}

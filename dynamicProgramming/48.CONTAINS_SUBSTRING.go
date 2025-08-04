package dynamicProgramming

const MOD = 1000000007

// powMod computes base^exp % mod using modular exponentiation
func powMod(base, exp, mod int64) int64 {
	result := int64(1)
	base = base % mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

func stringCount(n int) int {
	// Total strings: 26^n
	total := powMod(26, int64(n), MOD)

	// |A|: Strings with no 'l' = 25^n
	A := powMod(25, int64(n), MOD)

	// |B|: Strings with 0 or 1 'e' = 25^n + n * 25^(n-1)
	B := (powMod(25, int64(n), MOD) + (int64(n)*powMod(25, int64(n-1), MOD))%MOD) % MOD

	// |C|: Strings with no 't' = 25^n
	C := powMod(25, int64(n), MOD)

	// |A ∩ B|: Strings with no 'l' and 0 or 1 'e' = 24^n + n * 24^(n-1)
	AB := (powMod(24, int64(n), MOD) + (int64(n)*powMod(24, int64(n-1), MOD))%MOD) % MOD

	// |A ∩ C|: Strings with no 'l' and no 't' = 24^n
	AC := powMod(24, int64(n), MOD)

	// |B ∩ C|: Strings with no 't' and 0 or 1 'e' = 24^n + n * 24^(n-1)
	BC := (powMod(24, int64(n), MOD) + (int64(n)*powMod(24, int64(n-1), MOD))%MOD) % MOD

	// |A ∩ B ∩ C|: Strings with no 'l', no 't', and 0 or 1 'e' = 23^n + n * 23^(n-1)
	ABC := (powMod(23, int64(n), MOD) + (int64(n)*powMod(23, int64(n-1), MOD))%MOD) % MOD

	// Inclusion-Exclusion: |A ∪ B ∪ C| = |A| + |B| + |C| - |A ∩ B| - |A ∩ C| - |B ∩ C| + |A ∩ B ∩ C|
	invalid := (A + B + C - AB - AC - BC + ABC) % MOD

	// Valid strings = Total - Invalid
	result := (total - invalid + MOD) % MOD

	return int(result)
}

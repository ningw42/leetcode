func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	if m == n {
		return m
	}

	M := int(math.Ceil(math.Log2(float64(m))))
	N := int(math.Floor(math.Log2(float64(n))))

	if M > N {
		return naive(m, n)
	} else if M == N {
		if pow := int(math.Pow(2, float64(M))); pow == m {
			return naive(pow, n)
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func naive(m, n int) int {
	result := m
	for i := m + 1; i <= n; i++ {
		result = result & i
	}
	return result
}

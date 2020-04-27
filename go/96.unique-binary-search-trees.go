
// Catalan Number
func numTrees(n int) int {
	result := make([]int, n+1)
	result[0] = 1
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += result[j] * result[i-1-j]
		}
		result[i] = sum
	}

	return result[n]
}

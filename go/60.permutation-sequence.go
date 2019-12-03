/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */

// @lc code=start
func getPermutation(n int, k int) string {
	result := make([]byte, n)
	candidates := make([]int, n+1)
	factorials := make([]int, n+1)
	candidates[0] = 0
	factorials[0] = 1
	for i := 1; i <= n; i++ {
		candidates[i] = i
		factorials[i] = factorials[i-1] * i
	}

	get(result, candidates, factorials, n, k, 0)

	return string(result)
}

func get(result []byte, candidates, factorials []int, n, k, i int) {
	if i < n {
		quotient := k / factorials[n-i-1]
		remainder := k % factorials[n-i-1]
		if remainder == 0 {
			remainder = factorials[n-i-1]
			quotient--
		}
		result[i] = byte(candidates[quotient+1]) + '0'
		candidates = append(candidates[:quotient+1], candidates[quotient+2:]...)
		get(result, candidates, factorials, n, remainder, i+1)
	}
}
// @lc code=end


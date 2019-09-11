/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	// F(n) = F(n-1) + F(n-2)
	results := []int{0, 1, 2}
	if n < 3 {
		return results[n]
	} else {
		for i := 3; i <= n; i++ {
			results = append(results, results[i-1] + results[i-2])
		}
		return results[n]
	}
}


/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
	upper := int(math.Floor(math.Sqrt(float64(n))))
	var squareNumbers []int
	for i := 1; i <= upper; i++ {
		squareNumbers = append(squareNumbers, i * i)
	}

	var result int
	results := make([]int, n+1)
	results[0] = 0
	for i := 1; i <= n; i++ {
		result = math.MaxInt32
		// can be optimized
		for j := 0; j < len(squareNumbers) && squareNumbers[j] <= i; j++ {
			if count := results[i-squareNumbers[j]] + 1; count < result {
				result = count
			}
		}
		results[i] = result
	}

	return results[n]
}
// @lc code=end


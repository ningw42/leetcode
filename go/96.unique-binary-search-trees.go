/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start

// Catalan Number
// dynamic programming
// better explanation see https://leetcode.com/problems/unique-binary-search-trees/discuss/31706/Dp-problem.-10%2B-lines-with-comments
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

// @lc code=end


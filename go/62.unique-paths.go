/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
func uniquePaths(m int, n int) int {
	var result [][]int
	for i := 0; i < m; i++ {
		result = append(result, make([]int, n))
	}

	if m == 1 || n == 1 {
		return 1
	} else {
		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				var top, left int
				if i - 1 == 0 {
					top = 1
				} else {
					top = result[i-1][j]
				}
				if j - 1 == 0 {
					left = 1
				} else {
					left = result[i][j-1]
				}
				result[i][j] = top + left
			}
		}
		return result[m-1][n-1]
	}
}


/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var result [][]int
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		result = append(result, make([]int, n))
	}

	// init edge
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		} else {
			result[0][i] = 1
		}
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		} else {
			result[i][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
			} else {
				result[i][j] = result[i-1][j] + result[i][j-1]
			}
		}
	}

	return result[m-1][n-1]
}


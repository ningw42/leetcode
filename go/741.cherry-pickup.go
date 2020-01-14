/*
 * @lc app=leetcode id=741 lang=golang
 *
 * [741] Cherry Pickup
 */

// @lc code=start
func cherryPickup(grid [][]int) int {
	n := len(grid)
	results := make([][][]int, n + 1)
	for i := 0; i <= n; i++ {
		results[i] = make([][]int, n + 1)
		for j := 0; j <= n; j++ {
			results[i][j] = make([]int, n + 1)
			for k := 0; k <= n; k++ {
				results[i][j][k] = math.MinInt32
			}
		}
	}

	results[1][1][1] = grid[0][0]
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for p := 1; p <= n; p++ {
				q := i + j - p
				if q < 1 || q > n || grid[i-1][j-1] == -1 || grid[p-1][q-1] == -1 {
					continue
				}
				max := maxInt(maxInt(results[i-1][j][p], results[i-1][j][p-1]), maxInt(results[i][j-1][p], results[i][j-1][p-1]))
				if max < 0 {
					continue
				}
				if i == p && j == q {
					results[i][j][p] = max + grid[i-1][j-1]
				} else {
					results[i][j][p] = max + grid[i-1][j-1] + grid[p-1][q-1]
				}
			}
		}
	}

	if results[n][n][n] < 0 {
		return 0
	} else {
		return results[n][n][n]
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end


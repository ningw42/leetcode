/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
func numIslands(grid [][]byte) int {
	var traversed [][]bool
	for _, row := range grid {
		traversed = append(traversed, make([]bool, len(row)))
	}

	count := 0
	for i, row := range grid {
		for j, _ := range row {
			if isNewIsland(grid, traversed, i, j) {
				count++
			}
		}
	}

	return count
}

func isNewIsland(grid [][]byte, traversed [][]bool, i int, j int) bool {
	if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) {
		if grid[i][j] == '1' && !traversed[i][j] {
			traversed[i][j] = true
			isNewIsland(grid, traversed, i-1, j)
			isNewIsland(grid, traversed, i+1, j)
			isNewIsland(grid, traversed, i, j-1)
			isNewIsland(grid, traversed, i, j+1)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}


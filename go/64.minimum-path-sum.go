/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
func minPathSum(grid [][]int) int {
	var minSum [][]int
	for _, row := range grid {
		minSum = append(minSum, make([]int, len(row)))
	}

	minSum[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			left := getSum(minSum, i, j - 1) + grid[i][j]
			top := getSum(minSum, i - 1, j) + grid[i][j]
			if left > top {
				minSum[i][j] = top
			} else {
				minSum[i][j] = left
			}
		}
	}
	
	fmt.Println(minSum)

	return minSum[len(grid)-1][len(grid[0])-1]
}

func getSum(minSum [][]int, i, j int) int {
	if i < 0 || i >= len(minSum) || j < 0 || j >= len(minSum[0]) {
		return math.MaxInt32 - math.MaxInt16
	} else {
		return minSum[i][j]
	}
}


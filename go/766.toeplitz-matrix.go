/*
 * @lc app=leetcode id=766 lang=golang
 *
 * [766] Toeplitz Matrix
 */
func isToeplitzMatrix(matrix [][]int) bool {
	rowCount := len(matrix)
	if rowCount == 0 {
		return true
	}
	colCount := len(matrix[0])

	for x := 0; x < rowCount; x++ {
		if !isSame(matrix, x, 0) {
			return false
		}
	}
	for y := 1; y < colCount; y++ {
		if !isSame(matrix, 0, y) {
			return false
		}
	}

	return true
}

func isSame(matrix [][]int, x, y int) bool {
	for i, j := x, y; i < len(matrix) && j < len(matrix[0]); {
		if matrix[x][y] != matrix[i][j] {
			return false
		}
		i++
		j++
	}

	return true
}


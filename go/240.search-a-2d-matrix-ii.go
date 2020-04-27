// tilt the matrix, it looks like a BST
// search from the top right corner
// O(m+n)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if value := matrix[row][col]; target > value {
			row++
		} else if target < value {
			col--
		} else {
			return true
		}
	}
	return false
}
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	var rowLow, rowHigh, rowMid int
	rowLow = 0
	rowHigh = len(matrix) - 1
	for rowLow <= rowHigh {
		rowMid = (rowLow + rowHigh) / 2
		if mid := matrix[rowMid][0]; mid < target {
			if tail := matrix[rowMid][len(matrix[0])-1]; tail > target {
				break
			} else if tail < target {
				rowLow = rowMid + 1
			} else {
				return true
			}
		} else if mid > target {
			rowHigh = rowMid - 1
		} else {
			return true
		}
	}

	var colLeft, colRight, colMid int
	colLeft = 0
	colRight = len(matrix[0]) - 1
	for colLeft <= colRight {
		colMid = (colLeft + colRight) / 2
		if mid := matrix[rowMid][colMid]; mid < target {
			colLeft = colMid + 1
		} else if mid > target {
			colRight = colMid - 1
		} else {
			return true
		}
	}
	return false
}
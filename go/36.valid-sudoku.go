/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		// row
		if !isValid(board[i]) {
			return false
		}
		// column
		column := []byte{}
		for j := 0; j < 9; j++ {
			column = append(column, board[j][i])
		}
		if !isValid(column) {
			return false
		}
	}
	// cells
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := []byte{}
			for x := i * 3; x < (i+1) * 3; x++ {
				for y := j * 3; y < (j+1) * 3; y++ {
					cell = append(cell, board[x][y])
				}
			}
			if !isValid(cell) {
				return false
			}
		}
	}

	return true
}

func isValid(elems []byte) bool {
	m := make(map[byte]bool)
	for _, c := range elems {
		if c != '.' {
			if _, exists := m[c]; exists {
				return false
			}
			m[c] = true
		}
	}
	return true
}


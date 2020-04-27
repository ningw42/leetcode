
const ALL = 0x1FF
var bitmap map[byte]int

func solveSudoku(board [][]byte) {
	buildBitMap()
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				availableValues := findAvailableValues(board, i, j)
				for _, value := range availableValues {
					board[i][j] = value
					if solve(board) {
						return true
					}
				}
				board[i][j] = '.'
				return false
			}
		}
	}

	return true
}

func findAvailableValues(board [][]byte, i, j int) []byte {
	availableValues := ALL

	// row
	for _, c := range board[i] {
		if c != '.' {
			availableValues &= bitmap[c]
		}
	}

	// column 
	for x := 0; x < 9; x++ {
		if c := board[x][j]; c != '.' {
			availableValues &= bitmap[c]
		}
	}

	// cell
	rowLo := i / 3 * 3
	rowHi := (i / 3 + 1) * 3
	colLo := j / 3 * 3
	colHi := (j / 3 + 1) * 3
	for x := rowLo; x < rowHi; x++ {
		for y := colLo; y < colHi; y++ {
			if c := board[x][y]; c != '.' {
				availableValues &= bitmap[c]
			}
		}
	}

	// convert bitmap form into byte
	var values []byte
	currentValue := byte('1')
	for availableValues != 0 {
		if availableValues & 1 == 1 {
			values = append(values, currentValue)
		}
		currentValue++
		availableValues = availableValues >> 1
	}

	return values
}

func buildBitMap() {
	bitmap = make(map[byte]int)
	var i uint
	for i = 1; i < 10; i++ {
		bitmap[byte(i) + '0'] = ALL ^ (1<<(i-1))
	}
}
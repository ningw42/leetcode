/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */
func solve(board [][]byte) {
	// naive(board)
	reference(board)
}

func reference(board [][]byte) {
	// https://leetcode.com/problems/surrounded-regions/discuss/41612/A-really-simple-and-readable-C%2B%2B-solutionuff0conly-cost-12ms
	// 1. mark all 'O's on borders as 'escaped'
	// 2. perform DFS on all 'escaped' 'Os'
	// 3. flip all rest 'O'
	var escaped[][]bool
	for _, row := range board {
		escaped = append(escaped, make([]bool, len(row)))
	}

	// step 1
	for i, row := range board {
		for j, _ := range row {
			if i == 0 || i == len(board) - 1 || j == 0 || j == len(row) - 1 {
				// step 2
				markAsEscaped(board, escaped, i, j)
			}
		}
	}

	// step 3
	for i, row := range board {
		for j, c := range row {
			if c == 'O' && !escaped[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func markAsEscaped(board [][]byte, escaped [][]bool, x int, y int) {
	if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) {
		if board[x][y] == 'O' && !escaped[x][y] {
			escaped[x][y] = true
			markAsEscaped(board, escaped, x-1, y)
			markAsEscaped(board, escaped, x+1, y)
			markAsEscaped(board, escaped, x, y-1)
			markAsEscaped(board, escaped, x, y+1)
		} else {
			return 
		}
	} else {
		return 
	}
}

// naive method with poor performance
func naive(board [][]byte) {
	for i, row := range board {
		for j, c := range row {
			if c == 'O' {
				var traversed[][]bool
				for _, row := range board {
					traversed = append(traversed, make([]bool, len(row)))
				}
				if !flipCaptured(board, traversed, i, j) {
					for ti, trow := range traversed {
						for tj, t := range trow {
							if t {
								board[ti][tj] = 'O'
							}
						}
					}
				}
			}
		}
	}
}

func flipCaptured(board [][]byte, traversed [][]bool, i int, j int) bool {
	traversed[i][j] = true
	if i == 0 || i == len(board) - 1 || j == 0 || j == len(board[0]) - 1 {
		return false
	} else {
		for _, x := range []int{-1, 0, 1} {
			for _, y := range []int{-1, 0, 1} {
				if math.Abs(float64(x)) + math.Abs(float64(y)) == 1 {
					ni := i + x
					nj := j + y
					if !traversed[ni][nj] && board[ni][nj] == 'O' {
						if !flipCaptured(board, traversed, ni, nj) {
							return false
						}
					}
				}
			}
		}

		board[i][j] = 'X'
		return true
	}
}




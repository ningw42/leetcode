/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				board[i][j] = ' '
				if adjacentExist(board, i, j, word[1:]) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}
	return false
}

func adjacentExist(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}

	coods := [][]int{
		[]int{i, j+1},
		[]int{i, j-1},
		[]int{i-1, j},
		[]int{i+1, j},
	}
	for _, cood := range coods {
		x := cood[0]
		y := cood[1]
		if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) {
			if board[x][y] == word[0] {
				board[x][y] = ' '
				if adjacentExist(board, x, y, word[1:]) {
					return true
				}
				board[x][y] = word[0]
			}
		}
	}
	return false
}
// @lc code=end


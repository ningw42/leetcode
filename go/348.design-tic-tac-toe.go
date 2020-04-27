type TicTacToe struct {
    n int
    board [][]int
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    board := make([][]int, n)
    for i := 0; i < n; i++ {
        board[i] = make([]int, n)
    }
    return TicTacToe{n: n, board: board}
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    this.board[row][col] = player
    horizontal := true
    vertical := true
    diagonal := true
    antiDiagonal := true
    for i := 0; i < this.n; i++ {
        if this.board[row][i] != player {
            horizontal = false
        }
        if this.board[i][col] != player {
            vertical = false
        }
        if this.board[i][i] != player {
            diagonal = false
        }
        if this.board[i][this.n-1-i] != player {
            antiDiagonal = false
        }
        if !(horizontal || vertical || diagonal || antiDiagonal) {
            break
        }
    }
    if horizontal || vertical || diagonal || antiDiagonal {
        return player
    } else {
        return 0
    }
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
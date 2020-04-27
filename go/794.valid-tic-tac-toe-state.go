func validTicTacToe(board []string) bool {
    countX, countO := 0, 0
    for _, row := range board {
        for j := 0; j < len(row); j++ {
            state := row[j]
            if state == 'X' {
                countX++
            } else if state == 'O' {
                countO++
            }
        }
    }
    
    diff := countX - countO
    if !(diff == 0 || diff == 1) {
        return false
    }

    x := hasWon(board, 'X')
    o := hasWon(board, 'O')
   
    if diff == 0 {
        // O's turn, only O can win
        return !x
    }
    if diff == 1 {
        // X's turn, only X can win
        return !o
    }
    
    return !(x && o)
}

func hasWon(board []string, player byte) bool {
    // row
    for i := 0; i < 3; i++ {
        rowWin := true
        for j := 0; j < 3; j++ {
            if board[i][j] != player {
                rowWin = false
                break
            }
        }
        if rowWin {
            return true
        }
    }
    // column
    for j := 0; j < 3; j++ {
        colWin := true
        for i := 0; i < 3; i++ {
            if board[i][j] != player {
                colWin = false
                break
            }
        }
        if colWin {
            return true
        }
    }
    // diagonal
    dWin := true
    for i := 0; i < 3; i++ {
        if board[i][i] != player {
            dWin = false
            break
        }
    }
    if dWin {
        return true
    }
    // anti-diagonal
    adWin := true
    for i := 0; i < 3; i++ {
        if board[i][2-i] != player {
            adWin = false
            break
        }
    }
    return adWin
}
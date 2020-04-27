// DFS runs in O(N) but requires O(N) extra memory
// counting the top-left most cell will do the trick with no extra memory
func countBattleships(board [][]byte) int {
    count := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == '.' {continue}
            // board[i][j] == 'X'
            if i > 0 && board[i-1][j] == 'X' {continue}
            if j > 0 && board[i][j-1] == 'X' {continue}
            count++
        }
    }
    return count
}
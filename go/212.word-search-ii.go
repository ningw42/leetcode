func findWords(board [][]byte, words []string) []string {
    if len(board) == 0 {
        return nil
    }
    var ret []string
    used := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        used[i] = make([]bool, len(board[0]))
    }
    for _, word := range words {
        if findWord(board, used, word) {
            ret = append(ret, word)
        }
    }
    return ret
}

func findWord(board [][]byte, used [][]bool, word string) bool {
    w := []byte(word)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if findChar(board, used, i, j, w) {
                return true
            }
        }
    }
    return false
}

func findChar(board [][]byte, used [][]bool, i, j int, word []byte) bool {
    if len(word) == 0 {
        return true
    } else {
        used[i][j] = true
        defer func(used [][]bool, i, j int) {used[i][j] = false}(used, i, j)
        if board[i][j] == word[0] {
            coords := [][]int{
                []int{i-1,j},
                []int{i+1,j},
                []int{i,j-1},
                []int{i,j+1},
            }
            for _, coord := range coords {
                x, y := coord[0], coord[1]
                if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) {
                    if !used[x][y] && findChar(board, used, x, y, word[1:]) {
                        return true
                    }
                }
            }
            // [["a"]]
            if len(word) == 1 {
                return true
            } else {
                return false
            }
        } else {
            return false
        }
    }
}
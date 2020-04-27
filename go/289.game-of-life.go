func gameOfLife(board [][]int) {
    var neighbor int
    var live, dead [][]int
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            neighbor = countLiveNeighbors(board, i, j)
            if board[i][j] == 1 {
                if neighbor < 2 || neighbor > 3 {
                    dead = append(dead, []int{i, j})
                }
            } else {
                if neighbor == 3 {
                    live = append(live, []int{i, j})
                }
            }
        }
    }
    setStatus(board, live, 1)
    setStatus(board, dead, 0)
}

func countLiveNeighbors(board [][]int, i, j int) int {
    coords := [][]int{
        {i-1, j-1},
        {i-1, j},
        {i-1, j+1},
        {i, j-1},
        {i, j+1},
        {i+1, j-1},
        {i+1, j},
        {i+1, j+1},
    }
  
    var x, y, count int
    for _, coord := range coords {
        x, y = coord[0], coord[1]
        if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) {
            if board[x][y] == 1 {
                count++
            }
        }
    }
    
    return count
}

func setStatus(board [][]int, coords [][]int, status int) {
    for _, coord := range coords {
        board[coord[0]][coord[1]] = status
    }
}
func snakesAndLadders(board [][]int) int {
    return bfs(board)
}

func bfs(board [][]int) int {
    n := len(board)
    line := make([]int, n * n + 1)
   
    // flatten 'board' into 'line' for better readability
    index := 1
    for i := n - 1; i >= 0; i-- {
        var j, step int
        var cond func (int) bool
        if i % 2 == (n-1) % 2 {
            j, step, cond = 0, 1, func (j int) bool {return j < n}
        } else {
            j, step, cond = n - 1, -1, func (j int) bool {return j >= 0}
        }
        for ; cond(j); j = j + step {
            line[index] = board[i][j]
            index++
        }
    }

    // bfs
    visited := make([]bool, n*n+1)
    moves := make([]int, n*n+1)
    
    var current, next []int 
    var move int
    current = append(current, 1)
    for len(current) > 0 {
        for len(current) > 0 {
            from := current[0]
            current = current[1:]
            if visited[from] {
                continue
            }
            visited[from] = true
            moves[from] = move
            for to := from + 1; to <= n*n && to <= from + 6; to++ {
                if line[to] == -1 {
                    // if there is no shortcut
                    next = append(next, to)
                } else {
                    // if there is shortcut
                    next = append(next, line[to])
                }
            }
        }
        move++
        current, next = next, current
    }
    
    if moves[n*n] == 0 {
        return -1
    } else {
        return moves[n*n]
    }
}

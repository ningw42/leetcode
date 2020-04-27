// bfs
func minKnightMoves(x int, y int) int {
    x = abs(x)
    y = abs(y)
    if y < x {
        x, y = y, x
    }
    coordinates := [][]int{
        {1,2},
        {-1,2},
        {1,-2},
        {-1,-2},
        {2,1},
        {-2,1},
        {2,-1},
        {-2,-1},
    }
    var queue, next [][]int
    visited := make(map[string]struct{})
    queue = append(queue, []int{0, 0})
    visited[getCoordinateString(0, 0)] = struct{}{}
    moveCount := 0
    for len(queue) > 0 {
        for len(queue) > 0 {
            current := queue[0]
            queue = queue[1:]
           
            if current[0] == x && current[1] == y {
                return moveCount
            }
            
            for _, coordinate := range coordinates {
                a, b := current[0] + coordinate[0], current[1] + coordinate[1]
                if _, v := visited[getCoordinateString(a, b)]; !v && a >= -1 && b >= -1 {
                    next = append(next, []int{a, b})
                    // mark as visited
                    visited[getCoordinateString(a, b)] = struct{}{}
                }
            }
        }
        queue, next = next, queue
        moveCount++
    }
    // not reachable
    return 0
}

func getCoordinateString(x, y int) string {
    return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}

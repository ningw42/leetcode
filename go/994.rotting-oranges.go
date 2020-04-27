func orangesRotting(grid [][]int) int {
    minutes := 0
    next := make([][]int, len(grid))
    for i := range grid {
        next[i] = make([]int, len(grid[0]))
        copy(next[i], grid[i])
    }

    // simulates the process of rotting
    rotted := true
    for rotted {
        rotted = rotting(grid, next)
        for i := range grid {
            copy(grid[i], next[i])
        }
        minutes++
    }

    for i := range grid {
        for _, status := range grid[i] {
            if status == 1 {
                return -1
            }
        }
    }
    return minutes - 1
}

func rotting(current, next [][]int) bool {
    var rotted bool
    for i := range current {
        for j, status := range current[i] {
            if status == 2 {
                // rotting
                coods := [][]int{
                    []int{i+1, j},
                    []int{i-1, j},
                    []int{i, j+1},
                    []int{i, j-1},
                }
                for _, cood := range coods {
                    x, y := cood[0], cood[1]
                    if 0 <= x && x < len(current) && 0 <= y && y < len(current[0]) && current[x][y] == 1 {
                        next[x][y] = 2
                        rotted = true
                    }
                }
            }
        }
    }
    return rotted
}
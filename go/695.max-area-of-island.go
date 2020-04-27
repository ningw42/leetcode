func maxAreaOfIsland(grid [][]int) int {
    visited := make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ {
        visited[i] = make([]bool, len(grid[0]))
    }
   
    maxArea := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if !visited[i][j] {
                maxArea = max(maxArea, dfs(grid, visited, i, j))
            }
        }
    }
    
    return maxArea
}

func dfs(grid [][]int, visited [][]bool, i, j int) int {
    // mark as visited
    visited[i][j] = true
    if grid[i][j] == 0 {
        return 0
    } else {
        area := 1
        coords := [][]int{
            {i-1,j},
            {i+1,j},
            {i,j-1},
            {i,j+1},
        }
        for _, coord := range coords {
            x, y := coord[0], coord[1]
            if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && !visited[x][y] {
                area += dfs(grid, visited, x, y)
            }
        }
        return area
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
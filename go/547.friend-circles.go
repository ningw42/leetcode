func findCircleNum(M [][]int) int {
    visited := make([]bool, len(M))
    circle := 0
    var i int
    for i < len(M) {
        if !visited[i] {
            // dfs
            stack := []int{i}
            for len(stack) > 0 {
                // pop from stack
                from := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                visited[from] = true
                for to, isFriend := range M[from] {
                    if !visited[to] && isFriend == 1 {
                        stack = append(stack, to)
                    }
                }
            }
            circle++
        }
        // skip visited students
        for i < len(M) && visited[i] {i++}
    }
    return circle
}
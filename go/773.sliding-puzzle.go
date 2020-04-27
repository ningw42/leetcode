func slidingPuzzle(board [][]int) int {
    graph, begin, end := buildGraph(board)
    if begin == end {
        return 0
    }
    visited := make([]bool, 720)
    // bfs
    dist := 0
    queue := []int{end}
    var stash []int
    for len(queue) > 0 {
        dist++
        for len(queue) > 0 {
            current := queue[0]
            queue = queue[1:]
            visited[current] = true
            for _, to := range graph[current] {
                if !visited[to] {
                    if to == begin {
                        return dist
                    }
                    stash = append(stash, to)
                }
            }
        }
        queue, stash = stash, queue
    }
    
    return -1
}

func buildGraph(board [][]int) (map[int][]int, int, int) {
    end := []int{1, 2, 3, 4, 5, 0}
    begin := []int{}
    begin = append(begin, board[0]...)
    begin = append(begin, board[1]...)
    permutations := findAllPermutations(6, []int{}, make(map[int]struct{}))
    graph := make(map[int][]int)
    beginIndex, endIndex := -1, -1
    for i, p := range permutations {
        for j := i + 1; j < len(permutations); j++ {
            q := permutations[j]
            if transition(p, q, i, j) {
                graph[i] = append(graph[i], j)
                graph[j] = append(graph[j], i)
            }
        }
        if beginIndex == -1 && identical(p, begin) {
            beginIndex = i
        }
        if endIndex == -1 && identical(p, end) {
            endIndex = i
        }
    }
    return graph, beginIndex, endIndex
}

func transition(a, b []int, i, j int) bool {
    var diffPos []int
    containsZero := false
    for i, aa := range a {
        if b[i] != aa {
            containsZero = containsZero || (aa == 0 || b[i] == 0)
            if len(diffPos) >= 2 {
                return false
            } else {
                diffPos = append(diffPos, i)
            }
        }
    }
    if !containsZero {
        return false
    }
    // fmt.Println(i, a, j, b, diffPos)
    sort.Ints(diffPos)
    x, y := diffPos[0], diffPos[1]
    connections := map[int]map[int]int{
        0: {1:0, 3:0},
        1: {2:0, 4:0},
        2: {5:0},
        3: {4:0},
        4: {5:0},
    }
    _, exists := connections[x][y]
    return exists
}

func identical(a, b []int) bool {
    for i, aa := range a {
        if b[i] != aa {
            return false
        }
    }
    return true
}

func findAllPermutations(bound int, permutation []int, state map[int]struct{}) [][]int {
    if len(state) == bound {
        return [][]int{makeSliceCopy(permutation)}
    } else {
        var results [][]int
        for candidate := 0; candidate < bound; candidate++ {
            if _, used := state[candidate]; !used {
                state[candidate] = struct{}{}
                results = append(results, findAllPermutations(bound, append(permutation, candidate), state)...)
                delete(state, candidate)
            }
        }
        return results
    }
}

func makeSliceCopy(s []int) []int {
    d := make([]int, len(s))
    copy(d, s)
    return d
}
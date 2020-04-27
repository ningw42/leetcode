func countComponents(n int, edges [][]int) int {
    disjointSet := make([]int, n)
    for i := 0; i < n; i++ {
        disjointSet[i] = i
    }
    for _, edge := range edges {
        union(disjointSet, edge[0], edge[1])
    }
    m := make(map[int]struct{})
    for i := 0; i < n; i++ {
        m[find(disjointSet, i)] = struct{}{}
    }
    return len(m)
}

func find(disjointSet []int, i int) int {
    for disjointSet[i] != i {
        i = disjointSet[i]
    }
    return i
}

func union(disjointSet []int, i, j int) {
    rootI := find(disjointSet, i)
    rootJ := find(disjointSet, j)
    if rootI != rootJ {
        disjointSet[rootI] = rootJ
    }
}
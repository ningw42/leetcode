func getSkyline(buildings [][]int) [][]int {
    return bruteForce(buildings)
}

func bruteForce(buildings [][]int) [][]int {
    var xs []int
    for _, building := range buildings {
        xs = append(xs, building[0], building[1])
    }
    sort.Ints(xs)

    var points [][]int
    var y int
    for _, x := range xs {
        y = 0
        for _, building := range buildings {
            if building[0] <= x && x < building[1] {
                y = max(y, building[2])
            }
        }
        if len(points) == 0 || y != points[len(points)-1][1] {
            points = append(points, []int{x, y})
        }
    }
    
    return points
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
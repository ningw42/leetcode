func reconstructQueue(people [][]int) [][]int {
    // sort by h asc and k desc
    sort.Slice(people, func(i, j int) bool {
        if people[i][0] < people[j][0] {
            return true
        } else if people[i][0] > people[j][0] {
            return false
        } else {
            return people[i][1] > people[j][1]
        }
    })
    
    order := make([][]int, len(people))
    for _, p := range people {
        i := 0
        // count empty slot until meeting the k(p[1])
        emptyCount := 0
        for emptyCount < p[1] && i < len(people) {
            if order[i] == nil {
                emptyCount++
            }
            i++
        }
        // find next empty slot
        for order[i] != nil {i++}
        order[i] = p
    }
    
    return order
}
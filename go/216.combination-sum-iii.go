var results [][]int

func combinationSum3(k int, n int) [][]int {
    results = nil
    find(k, n, []int{})
    return results
}

func find(k int, n int, combination []int) {
    if k == 0 {
        if n == 0 {
            results = append(results, makeSliceCopy(combination))
        }
    } else {
        if n <= 0 {
            return
        } else {
            var i int
            if len(combination) == 0 {
                i = 1
            } else {
                i = combination[len(combination)-1] + 1
            }
            for ; i < 10; i++ {
                find(k-1, n-i, append(combination, i))
            }
        }
    }
}

func makeSliceCopy(s []int) []int {
    c := make([]int, len(s))
    copy(c, s)
    return c
}
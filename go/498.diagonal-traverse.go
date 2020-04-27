func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }
    
    m := len(matrix)
    n := len(matrix[0])
    
    var i, j, delta int
    var result []int
    for sum := 0; sum <= m + n - 1; sum++ {
        if sum & 1 == 1 {
            i = 0
            delta = 1
        } else {
            i = m - 1
            delta = -1
        }
        for 0 <= i && i < m {
            j = sum - i
            if 0 <= j && j < n {
                result = append(result, matrix[i][j])
            }
            i += delta
        }
    }
    
    return result
}

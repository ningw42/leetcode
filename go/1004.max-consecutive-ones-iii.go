func longestOnes(A []int, K int) int {
    var longest int
    var i, j int
    for j < len(A) {
        for K == 0 && A[j] == 0 {
            if A[i] == 0 {
                K++
            }
            i++
        }
        longest = max(longest, j - i + 1)
        if A[j] == 0 {
            K--
        }
        j++
    }
    return longest
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

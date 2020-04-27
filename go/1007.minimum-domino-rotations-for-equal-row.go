func minDominoRotations(A []int, B []int) int {
    n := len(A)
    a := make(map[int]int)
    b := make(map[int]int)
    common := make(map[int]int)
    nums := make(map[int]struct{})
    
    for i := 0; i < n; i++ {
        a[A[i]]++
        b[B[i]]++
        nums[A[i]] = struct{}{}
        nums[B[i]] = struct{}{}
        if A[i] == B[i] {
            common[A[i]]++
        }
    }
    
    // find possible number
    var candidates []int
    for num := range nums {
        if a[num] + b[num] - common[num] >= n {
            candidates = append(candidates, num)
        }
    }
    
    if len(candidates) == 0 {
        return -1
    }
    minFlip := math.MaxInt32
    for _, candidate := range candidates {
        if count := min(n-a[candidate], n-b[candidate]); count < minFlip {
            minFlip = count
        }
    }
    return minFlip
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
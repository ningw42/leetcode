func minDeletionSize(A []string) int {
    length := len(A[0])
    decreasingCount := 0
    for i:= 0; i<length; i++ {
        previous := byte('a')
        for _, a := range A {
            char := a[i]
            if char < previous {
                decreasingCount++
                break
            }
            previous = char
        }
        
    }
    
    return decreasingCount
}
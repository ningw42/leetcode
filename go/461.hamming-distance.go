func hammingDistance(x int, y int) int {
    diff := x ^ y
    count := 0
    
    for diff != 0 {
        if diff & 1 == 1 {
            count += 1
        }
        diff = diff >> 1
    }
    
    return count
}
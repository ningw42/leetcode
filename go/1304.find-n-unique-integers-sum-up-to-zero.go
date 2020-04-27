func sumZero(n int) []int {
    var slice []int
    for i := 0; i < n / 2; i++ {
        slice = append(slice, i+1, -(i+1))
    }
    if len(slice) < n {
        slice = append(slice, 0)
    }
    return slice
}
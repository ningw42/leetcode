func dailyTemperatures(T []int) []int {
    ret := make([]int, len(T))
    
    // decending monotone stack
    var stack []int
    for i := len(T) - 1; i >= 0; i-- {
        for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            ret[i] = 0
        } else {
            ret[i] = stack[len(stack)-1] - i
        }
        stack = append(stack, i)
    }
    
    return ret
}
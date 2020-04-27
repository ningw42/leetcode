func minRemoveToMakeValid(s string) string {
    removedIndex := make(map[int]struct{})
    var stack []int // stack stores index of '('
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            // push
            stack = append(stack, i)
        }
        if s[i] == ')' {
            if len(stack) > 0 {
                // pop
                stack = stack[:len(stack)-1]
            } else {
                removedIndex[i] = struct{}{}
            }
        }
    }
    for len(stack) > 0 {
        removedIndex[stack[len(stack)-1]] = struct{}{}
        stack = stack[:len(stack)-1]
    }
    var builder strings.Builder
    for i := 0; i < len(s); i++ {
        if _, exists := removedIndex[i]; !exists {
            builder.WriteByte(s[i])
        }
    }
    return builder.String()
}
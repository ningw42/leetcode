func longestValidParentheses(s string) int {
	begin, max := 0, 0
	var stack []int
	var length int
	for i, r := range s {
		c := byte(r)
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				begin = i + 1
			} else {
				// index := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				if len(stack) == 0 {
					length = i - begin + 1
				} else {
					length = i - stack[len(stack)-1]
				}
				if length > max {
					max = length
				}
			}
		}
	}
	// prev := len(s)
	// for len(stack) > 0 {
	// 	curr := stack[len(stack)-1]
	// 	stack = stack[0 : len(stack)-1]
	// 	if length := prev - curr - 1; length > 1 && length > max {
	// 		max = length
	// 	}
	// 	prev = curr
	// }

	return max
}

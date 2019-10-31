/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	begin, max := 0, 0
	var stack []int
	var length int
	for i, r := range s {
		c := byte(r)
		if c == '(' {
			// push '('
			stack = append(stack, i)
		} else {
			// ')'
			if len(stack) == 0 {
				begin = i + 1
			} else {
				// pop '('
				stack = stack[0 : len(stack)-1]
				if len(stack) == 0 {
					// no element left in stack
					// length is end - begin
					length = i - begin + 1
				} else {
					// there are elements in stask
					// length is end - index_of_the_top_element
					length = i - stack[len(stack)-1]
				}
				if length > max {
					max = length
				}
			}
		}
	}

	return max
}

// @lc code=end


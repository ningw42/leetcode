/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	var stack []int
	var operand int
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			operandLeft := stack[len(stack)-2]
			operandRight := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				operand = operandLeft + operandRight
			case "-":
				operand = operandLeft - operandRight
			case "*":
				operand = operandLeft * operandRight
			case "/":
				operand = operandLeft / operandRight
			}
		} else {
			operand, _ = strconv.Atoi(token)
		}
		stack = append(stack, operand)
	}
	return stack[0]
}
// @lc code=end


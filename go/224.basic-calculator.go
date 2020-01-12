/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */

// @lc code=start
var precedence map[byte]int

func calculate(s string) int {
	// init precedence
	precedence = map[byte]int{
		'+':0,
		'-':0,
		'*':1,
		'/':1,
	}
	// remove white spaces
	infix := strings.Replace(s, " ", "", -1)
	postfix := toPostfix(infix)

	var stack []int
	for _, token := range postfix {
		switch token {
		case "+":
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, l+r)
		case "-":
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, l-r)
		case "*":
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, l*r)
		case "/":
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, l/r)
		default:
			// operand
			value, _ := strconv.Atoi(token)
			stack = append(stack, value)
		}
	}

	return stack[0]
}

func toPostfix(infix string) []string {
	var rpn []string
	var operators []byte
	for i := 0; i < len(infix); {
		if isOperater(infix[i]) {
			if len(operators) == 0 {
				operators = append(operators, infix[i])
			} else if precedence[infix[i]] > precedence[operators[len(operators)-1]] {
				operators = append(operators, infix[i])
			} else {
				for j := len(operators)-1; j >= 0 && precedence[operators[j]] >= precedence[infix[i]]; j-- {
					rpn = append(rpn, string([]byte{operators[j]}))
					operators = operators[:j]
				}
				operators = append(operators, infix[i])
			}
			i++
		}
		if isOperand(infix[i]) {
			var j int
			for j = i; j < len(infix) && isOperand(infix[j]); j++ {}
			rpn = append(rpn, infix[i:j])
			i = j
		}
		if i < len(infix) && infix[i] == '(' {
			var j int
			bracket := 1
			for j = i+1; j < len(infix) && bracket != 0; j++ {
				if infix[j] == '(' {
					bracket++
				}
				if infix[j] == ')' {
					bracket--
				}
			}
			rpn = append(rpn, toPostfix(infix[i+1:j-1])...)
			i = j
		}
	}
	for i := len(operators)-1; i >= 0; i-- {
		rpn = append(rpn, string([]byte{operators[i]}))
	}
	return rpn
}

func isOperand(b byte) bool {
	return '0' <= b && b <= '9'
}

func isOperater(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/'
}
// @lc code=end


var precedence map[byte]int

func calculate(s string) int {
    precedence = map[byte]int{
        '+': 1,
        '-': 1,
        '*': 2,
        '/': 2,
    }
    return evalRPN(toRPN(strings.Replace(s, " ", "", -1)))
}

func toRPN(s string) []string {
    var i int // cursor
    var rpn []string // rpn
    var operators []byte // operator stack
    for i < len(s) {
        if s[i] == '(' {
            // bracket
            bracketCount := 1
            j := i + 1
            for ; j < len(s) && bracketCount > 0; j++ {
                if s[j] == '(' {
                    bracketCount++
                } else if s[j] == ')' {
                    bracketCount--
                }
            }
            rpn = append(rpn, toRPN(s[i+1:j-1])...)
            i = j
        } else if isOperator(s[i]) {
            // operator
            if len(rpn) == 0 && s[i] == '-' {
                // negetive integer at beginning of string
                operand, newIndex := getOperand(s, i+1)
                rpn = append(rpn, "-" + operand)
                i = newIndex
            } else {
                // operator
                // pop operators in stack with less or equal precedence compared to current operator
                for len(operators) > 0 && precedence[s[i]] <= precedence[operators[len(operators)-1]] {
                    op := operators[len(operators)-1]
                    operators = operators[:len(operators)-1]
                    rpn = append(rpn, string([]byte{op}))
                }
                // put current operator into stack
                operators = append(operators, s[i])
                if s[i+1] == '-' {
                    operand, newIndex := getOperand(s, i+2)
                    rpn = append(rpn, "-" + operand)
                    i = newIndex
                } else {
                    i++
                }
            }
        } else {
            // operand
            operand, newIndex := getOperand(s, i)
            rpn = append(rpn, operand)
            i = newIndex
        }
    }
    // pop rest operators and append to rpn
    for len(operators) > 0 {
        op := operators[len(operators)-1]
        operators = operators[:len(operators)-1]
        rpn = append(rpn, string([]byte{op}))
    }
    return rpn
}

func evalRPN(rpn []string) int {
    var stack []int
    for _, token := range rpn {
        if len(token) == 1 && isOperator(token[0]) {
            // operator
            result := eval(stack[len(stack)-2], stack[len(stack)-1], token)
            stack = stack[:len(stack)-2]
            stack = append(stack, result)
        } else {
            // operand
            operand, _ := strconv.Atoi(token)
            stack = append(stack, operand)
        }
    }
    return stack[0]
}

func isOperator(b byte) bool {
    return b == '+' || b == '-' || b == '*' || b == '/'
}

func isDigit(b byte) bool {
    return '0' <= b && b <= '9'
}

func getOperand(s string, i int) (string, int) {
    var builder strings.Builder
    j := i
    for ; j < len(s) && isDigit(s[j]); j++ {
        builder.WriteByte(s[j])
    }
    return builder.String(), j
}

func eval(a, b int, op string) int {
    switch op {
        case "+":
        return a + b
        case "-":
        return a - b
        case "*":
        return a * b
        case "/":
        return a / b
    }
    return 0
}
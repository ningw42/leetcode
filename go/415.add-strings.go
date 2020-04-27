func addStrings(num1 string, num2 string) string {
    var sum []byte
    var digit byte
    var carry byte
    minLength := min(len(num1), len(num2))
    var i int
    for ; i < minLength; i++ {
        digit = (num1[len(num1)-1-i] - '0') + (num2[len(num2)-1-i] - '0') + carry
        if digit > 9 {
            digit = '0' + digit - 10
            carry = 1
        } else {
            digit = '0' + digit
            carry = 0
        }
        // fmt.Printf("%c+%c=%c\n", num1[len(num1)-1-i], num2[len(num2)-1-i], digit)
        sum = append(sum, digit)
    }
    var rest string
    if len(num1) < len(num2) {
        rest = num2[:len(num2)-i]
    } else {
        rest = num1[:len(num1)-i]
    }
    // fmt.Println(string(sum), len(rest), rest, carry)
    for i := len(rest) - 1; i >= 0; i-- {
        digit = (rest[i] - '0') + carry
        if digit > 9 {
            digit = '0' + (digit - 10)
            carry = 1
        } else {
            digit = '0' + digit
            carry = 0
        }
        sum = append(sum, digit)
    }
    if carry > 0 {
        sum = append(sum, '1')
    }
    reverse(sum)
    return string(sum)
}

func reverse(s []byte) {
    for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
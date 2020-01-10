func decodeString(s string) string {
    if s == "" {
        return ""
    }
    // count
    var i int
    for isDigit(s[i]) {
        i++
    }
    // letters
    if i == 0 {
        for i < len(s) && isAlphabet(s[i]) {
            i++
        }
        return s[:i] + decodeString(s[i:])
    } else {
        count, _ := strconv.Atoi(s[:i])
        bracket := 1
        begin := i + 1
        for i = i+1; i < len(s) && bracket != 0; i++ {
            if s[i] == '[' {
                bracket++
            }
            if s[i] == ']' {
                bracket--
            }
        }
        str := decodeString(s[begin:i-1])
        fmt.Println(str)
        builder := strings.Builder{}
        for j := 0; j < count; j++ {
            builder.WriteString(str)
        }
        return builder.String() + decodeString(s[i:])
    }
}

func isDigit(b byte) bool {
    return '0' <= b && b <= '9'
}

func isAlphabet(b byte) bool {
    return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'
}
func reverseWords(s string) string {
    bs := []byte(s)
    begin := 0
    for i := 0; i < len(bs); i++ {
        if bs[i] == ' ' {
            reverse(bs[begin:i])
            begin = i+1
        }
    }
    reverse(bs[begin:])
    return string(bs)
}

func reverse(bs []byte) {
    for i := 0; i < len(bs) / 2; i++ {
        j := len(bs) - 1 - i
        bs[i], bs[j] = bs[j], bs[i]
    }
}
func convertToTitle(n int) string {
    var title []byte
    for n != 0 {
        digit := (n - 1) % 26
        title = append(title, byte(digit) + 'A')
        n = (n - 1) / 26
    }
    if len(title) == 0 {
        return "A"
    }
    reverse(title)
    return string(title)
}

func reverse(s []byte) {
    var i, j int
    for i = 0; i < len(s) / 2; i++ {
        j = len(s) - 1 - i
        s[i], s[j] = s[j], s[i]
    }
}
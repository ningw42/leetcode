func lastSubstring(s string) string {
    // result must be a suffix of s
    return bruteForce(s)
    // return suffixArray(s)
}

func bruteForce(s string) string {
    max := s
    for i := 1; i < len(s); i++ {
        if now := s[i:]; now > max {
            max = now
        }
    }
    return max
}

func suffixArray(s string) string {
    return ""
}
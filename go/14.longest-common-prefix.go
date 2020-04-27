func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    i := 0
    loop:
    for {
        for _, str := range strs {
            if len(str) > i {
                if str[i] != strs[0][i] {
                    break loop
                }
            } else {
                break loop
            }
        }
        i++
    }
    return strs[0][:i]
}
func lengthOfLongestSubstringTwoDistinct(s string) int {
    count := make(map[byte]int)
    i, j := 0, 0
    max := 0
    for j < len(s) {
        // expand right bound
        count[s[j]]++
        j++
        // shrink left bound
        for len(count) > 2 {
            count[s[i]]--
            if count[s[i]] == 0 {
                delete(count, s[i])
            }
            i++
        }
        // mark max length
        if length := j - i; length > max {
            max = length
        }
    }
    return max
}
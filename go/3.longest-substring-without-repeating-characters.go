func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    length := 1
    begin, end := 0, 0
    count := map[byte]int{s[0]:1}
    for end < len(s) && begin <= end {
        for end < len(s) && len(count) == end - begin + 1 {
            if l := end - begin + 1; l > length {
                length = l
            }
            end++
            if end == len(s) {
                break
            }
            count[s[end]]++
        }
        count[s[begin]]--
        if count[s[begin]] == 0 {
            delete(count, s[begin])
        }
        begin++
    }
    return length
}
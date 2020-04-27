func lengthOfLongestSubstringKDistinct(s string, k int) int {
    var i, j int
    counter := make(map[byte]int)
  
    length := 0
    for j < len(s) {
        counter[s[j]]++
        for len(counter) > k {
            counter[s[i]]--
            if counter[s[i]] == 0 {
                delete(counter, s[i])
            }
            i++
        }
        if l := j - i + 1; l > length {
            length = l
        }
        j++
    }
    
    return length
}
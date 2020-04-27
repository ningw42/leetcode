func minWindow(s string, t string) string {
    // make a map of required chars
    required := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        required[t[i]]++
    }
    // make a map of included chars
    included := make(map[byte]int)
    
    numberOfDistinctRequiredChars := len(required)
    numberOfDistinctMeetChars := 0
    
    // two pointer to find min length
    var i, j, minLen, begin int
    minLen = math.MaxInt32
    for j < len(s) {
        // expand right side of window to meet the requirement
        included[s[j]]++
        if included[s[j]] == required[s[j]] {
            numberOfDistinctMeetChars++
        }
        // shrink left side of window to find a smaller length
        for i <= j && numberOfDistinctMeetChars == numberOfDistinctRequiredChars {
            if length := j - i + 1; length < minLen {
                minLen = length
                begin = i
            }
            included[s[i]]--
            if included[s[i]] < required[s[i]] {
                numberOfDistinctMeetChars--
            }
            i++
        }
        j++
    }

    if minLen == math.MaxInt32 {
        return ""
    } else {
        return s[begin:begin+minLen]
    }
}
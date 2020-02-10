/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */

// @lc code=start
var alphabet map[byte]int

func isAlienSorted(words []string, order string) bool {
    alphabet = make(map[byte]int)
    for i := 0; i < len(order); i++ {
        alphabet[order[i]] = i
    }
    
    for i := 0; i < len(words) - 1; i++ {
        if !lessOrEqual(words[i], words[i+1]) {
            return false
        }
    }
    return true
}

func lessOrEqual(w1, w2 string) bool {
    var cursor int
    var o1, o2 int
    bound := min(len(w1), len(w2))
    for cursor = 0; cursor < bound; cursor++ {
        o1 = alphabet[w1[cursor]]
        o2 = alphabet[w2[cursor]]
        if o1 > o2 {
            return false
        }
        if o1 < o2 {
            return true
        }
    }
    
    if bound == len(w1) && bound == len(w2) {
        return true
    } else if bound == len(w1) {
        return true
    } else {
        return false
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
// @lc code=end


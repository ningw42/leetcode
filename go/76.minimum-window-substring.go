import "math"

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	tCount := count(t)
	sCount := make(map[byte]int)

	start := 0
	length := math.MaxInt32
	begin := 0
	end := 0
	for end < len(s) {
		// find a valid window
		sCount[s[end]]++
		// shink the left side of window until the window
		// no longer satisfies the condition
		for satisfy(sCount, tCount) {
			// if new valid window satisfies the condition
			// and its length is smaller previous shortest window
			if l := end - begin; l < length {
				start = begin
				length = l
			}
			sCount[s[begin]]--
			begin++
		}
		end++
	}

	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+length+1]
	}
}

func satisfy(source, target map[byte]int) bool {
	for b, c := range target {
		if source[b] < c {
			return false
		}
	}
	return true
}

func count(s string) map[byte]int {
	ret := make(map[byte]int)
	for _, r := range s {
		c := byte(r)
		if _, exists := ret[c]; exists {
			ret[c]++
		} else {
			ret[c] = 1
		}
	}
	return ret
}

// @lc code=end


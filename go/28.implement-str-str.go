/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
func strStr(haystack string, needle string) int {
	lengthOfNeedle := len(needle)
	lengthOfHaystask := len(haystack)

	if lengthOfNeedle == 0 {
		return 0
	}
    if lengthOfNeedle > lengthOfHaystask {
		return -1
	}
	for i := 0; i + lengthOfNeedle <= lengthOfHaystask; i++ {
		if haystack[i:i+lengthOfNeedle] == needle {
			return i
		}
	}

	return -1
}


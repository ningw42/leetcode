/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	var left, right byte
	for i, j := 0, len(s)-1; i < j; {
		left, right = s[i], s[j]
		if !isValidChar(left) {
			i++
			continue
		}
		if !isValidChar(right) {
			j--
			continue
		}
		if !isEqualChar(left, right) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValidChar(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func isEqualChar(a, b byte) bool {
	return makeRegular(a) == makeRegular(b)
}

func makeRegular(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 32
	} else {
		return c
	}
}
// @lc code=end


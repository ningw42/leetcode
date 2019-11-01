/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	var result int
	results := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
	results[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i >= 2 {
					result = results[i-2]
				} else {
					result = results[i-1]
				}
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' {
				if i > 2 {
					result = results[i-1] + results[i-2]
				} else {
					result = results[i-1] + 1
				}
			} else if s[i-1] == '2' {
				if '1' <= s[i] && s[i] <= '6' {
					if i > 2 {
						result = results[i-1] + results[i-2]
					} else {
						result = results[i-1] + 1
					}
				} else {
					result = results[i-1]
				}
			} else {
				result = results[i-1]
			}
		}

		results[i] = result
	}

	return results[len(s)-1]
}

// @lc code=end


/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	counts := make(map[string]int)
	wordLen := len(words[0])
	length := 0
	for _, word := range words {
		counts[word]++
		length += len(word)
	}

	var j int
	var result []int
	for i := 0; i < len(s) - length + 1; i++ {
		seen := make(map[string]int)
		for j = 0; j < length; j += wordLen {
			next := s[i+j:i+j+wordLen]
			seen[next]++
			expectedCount, existInWords := counts[next]
			if !existInWords {
				break
			} else {
				if seen[next] > expectedCount {
					break
				}
			}
		}
		if j == length {
			result = append(result, i)
		}
	}
	return result
}


/*
 * @lc app=leetcode id=859 lang=golang
 *
 * [859] Buddy Strings
 */

// @lc code=start
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		dist := make(map[byte]int)
		for _, r := range A {
			c := byte(r)
			if _, exists := dist[c]; exists {
				dist[c]++
			} else {
				dist[c] = 1
			}
		}
		for _, count := range dist {
			if count > 1 {
				return true
			}
		}
		return false
	} else {
		diffCount := 0
		var pos []int
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				diffCount++
				pos = append(pos, i)
			}
			if diffCount > 2 {
				return false
			}
		}

		return A[pos[0]] == B[pos[1]] && A[pos[1]] == B[pos[0]]
	}
}

// @lc code=end


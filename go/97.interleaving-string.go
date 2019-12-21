/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start

// dynamic programming
// there is no need to keey the entire results matrix
// just a col/row should do the job
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}

	var tmp map[string]bool
	results1 := make([]map[string]bool, len(s1)+1)
	results2 := make([]map[string]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		results1[i] = map[string]bool{s1[:i]:true}
	}
	for j := 0; j <= len(s2); j++ {
		results2[j] = map[string]bool{s2[:j]:true}
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			result := make(map[string]bool)
			if s3[i+j-1] == s2[j-1] {
				if j == 1 {
					tmp = results1[i]
				} else {
					tmp = results2[j-1]
				}
				for s := range tmp {
					result[s+s2[j-1:j]] = true
				}
			}
			if s3[i+j-1] == s1[i-1] {
				for s := range results2[j] {
					result[s+s1[i-1:i]] = true
				}
			}
			results2[j] = result
		}
	}

	_, exists := results2[len(s2)][s3]
	return exists
}

// time limit exceeded
func recursion(s1, s2, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	} else {
		return interleave(s1, s2, s3) || interleave(s2, s1, s3)
	}
}

func interleave(s1, s2, target string) bool {
	if len(s1) == 0 && len(s2) == 0 && len(target) == 0 {
		return true
	}
	for i := 0; i < len(s1) && s1[i] == target[i]; i++ {
		if interleave(s1[i+1:], s2, target[i+1:]) || interleave(s2, s1[i+1:], target[i+1:]) {
			return true
		}
	}
	return false
}
// @lc code=end


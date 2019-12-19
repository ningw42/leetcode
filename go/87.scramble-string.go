/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		if len(s1) == 1 {
			return s1 == s2
		} else {
			return checkScramble(s1, s2)
		}
	}
}

func checkScramble(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return true
	} else if len(s1) == 2 {
		return s1 == s2 || (s1[0] == s2[1] && s1[1] == s2[0])
	} else {
		for i := 1; i < len(s1); i++ {
			// doesn't switch on current root 
			if shareSameDistributionForSameSide(s1, s2, i) {
				if checkScramble(s1[:i], s2[:i]) && checkScramble(s1[i:], s2[i:]) {
					return true
				}
			}
			// switches on current root
			if shareSameDistributionForOppositeSite(s1, s2, i) {
				if checkScramble(s1[:i], s2[len(s2)-i:]) && checkScramble(s1[i:], s2[:len(s2)-i]) {
					return true
				}
			}
		}
		return false
	}
}

func shareSameDistributionForSameSide(s1, s2 string, i int) bool {
	return shareSameDistribution(getDistribution(s1[i:]), getDistribution(s2[i:])) && shareSameDistribution(getDistribution(s1[:i]), getDistribution(s2[:i]))
}

func shareSameDistributionForOppositeSite(s1, s2 string, i int) bool {
	return shareSameDistribution(getDistribution(s1[i:]), getDistribution(s2[:len(s2)-i])) && shareSameDistribution(getDistribution(s1[:i]), getDistribution(s2[len(s2)-i:]))
}

func shareSameDistribution(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for char, aCount := range a {
		if bCount, bExists := b[char]; !bExists || aCount != bCount {
			return false
		}
	}
	return true
}

func getDistribution(s string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}
// @lc code=end


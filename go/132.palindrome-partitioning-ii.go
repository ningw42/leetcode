func minCut(s string) int {
	minCuts := make([]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		minCuts[i] = math.MaxInt32
	}
	minCuts[0] = -1
	for i := 1; i <= len(s); i++ {
		min := math.MaxInt32
		for j := 0; j < i; j++ {
			if minCuts[j] != math.MaxInt32 && isPalindrome(s[j:i]) {
				if cut := minCuts[j] + 1; cut < min {
					min = cut
				}
			}
		}
		minCuts[i] = min
	}
	return minCuts[len(s)]
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func longestPalindrome(s string) string {
	// there is algorighm with time complexity O(n)
	// see https://www.google.com/search?q=Manacher
	return DynamicProgramming(s)
}

// AC with time complexity O(n^2)
// F(i, i) = true
// F(i, j) = F(i+1, j-1) && s[i] == s[j]
func DynamicProgramming(s string) string {
	var ret string
	var result bool
	var results [][]bool
	for i := 0; i < len(s); i++ {
		results = append(results, make([]bool, len(s)))
	}

	for step := 0; step < len(s); step++ {
		for start := 0; start+step < len(s); start++ {
			i := start
			j := start + step
			if step == 0 {
				result = true
			} else if step == 1 {
				result = s[i] == s[j]
			} else {
				result = results[i+1][j-1] && s[i] == s[j]
			}
			if result && j-i+1 > len(ret) {
				ret = s[i : j+1]
			}
			results[i][j] = result
		}
	}

	return ret
}

// AC with time complexity O(n^3)
func BruteForce(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			isPalindrome := true
			for a, b := i, j; a <= b; {
				if s[a] != s[b] {
					isPalindrome = false
					break
				}
				a++
				b--
			}
			if isPalindrome && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}

	return result
}

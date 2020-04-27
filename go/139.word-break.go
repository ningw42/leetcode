func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	result := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		for j := -1; j < i; j++ {
			if _, exists := dict[s[j+1:i+1]]; exists && (j == -1 || result[j]) {
				result[i] = true
				break
			}
		}
	}

	return result[len(s)-1]
}
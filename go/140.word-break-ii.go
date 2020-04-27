
// naive backtracking wont work, time limit exceeded
// dynamic programming

// type Strings []string
// func (s Strings) Len() int {
// 	return len(s)
// }
// func (s Strings) Less(i, j int) bool {
// 	return len(s[i]) > len(s[j])
// }
// func (s Strings) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

func wordBreak(s string, wordDict []string) []string {
	return DFS(s, wordDict, map[string][]string{"":[]string{""}})
}

func DFS(s string, dict []string, results map[string][]string) []string {
	if r, exists := results[s]; exists {
		return r
	}
	var result []string
	for _, word := range dict {
		if strings.HasPrefix(s, word) {
			for _, sentence := range DFS(s[len(word):], dict, results) {
				result = append(result, strings.Trim(word+" "+sentence, " "))
			}
		}
	}
	results[s] = result
	return result
}

// AC with bug
// memory limit exceeded on following case
// "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
// ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
func wordBreakDP(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	if !canBreakDP(s, dict) {
		return nil
	}

	result := make([][]string, len(s))
	for i := 0; i < len(s); i++ {
		if _, exists := dict[s[:i+1]]; exists {
			result[i] = []string{s[:i+1]}
		}
		for j := 0; j < i; j++ {
			if _, exists := dict[s[j+1:i+1]]; exists && len(result[j]) > 0 {
				for _, sentence := range result[j] {
					result[i] = append(result[i], sentence+" "+s[j+1:i+1])
				}
			}
		}
	}

	return result[len(s)-1]
}

func canBreakDP(s string, dict map[string]bool) bool {
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
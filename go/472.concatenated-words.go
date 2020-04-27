func findAllConcatenatedWordsInADict(words []string) []string {
    return dfs(words)
}

func dfs(words []string) []string {
    if len(words) == 0 {
        return nil
    }
    
    // make dictionary
    dict := make(map[string]struct{})
    for _, word := range words {
        dict[word] = struct{}{}
    }
  
    // dfs
    var results []string
    for _, word := range words {
        if word == "" {continue}
        delete(dict, word)
        if isTarget(dict, word) {
            results = append(results, word)
        }
        dict[word] = struct{}{}
    }
    
    return results
}

func isTarget(dict map[string]struct{}, word string) bool {
    if len(word) == 0 {
        return true
    }
    for i := 1; i <= len(word); i++ {
        if _, exists := dict[word[:i]]; exists {
            if isTarget(dict, word[i:]) {
                return true
            }
        }
    }
    return false
}
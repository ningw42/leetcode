func longestStrChain(words []string) int {
    // sort by length asc
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    dp := map[string]int{}
    maxChain := math.MinInt32
    wordMaxChain := math.MinInt32
    for _, word := range words {
        wordMaxChain = math.MinInt32
        for i := 0; i < len(word); i++ {
            wordMaxChain = max(wordMaxChain, dp[word[:i]+word[i+1:]] + 1) 
        }
        dp[word] = wordMaxChain
        maxChain = max(maxChain, wordMaxChain)
    }
    return maxChain
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
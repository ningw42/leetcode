func stoneGame(piles []int) bool {
    return alexAlwaysWins()
    // return dynamicProgramming(piles)
}

func alexAlwaysWins() bool {
    return true
}

func dynamicProgramming(piles []int) bool {
    // dp[i][j] = max relative score you can get given piles[i] to piles[j] inclusive
    // dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
    // since both players are playing optimally, whose turn doesn't matter
    // player has two choices for each decision, pick left most pile or right most pile
    // choose the one with max relative score
    var dp [][]int
    dp = make([][]int, len(piles))
    for i := 0; i < len(piles); i++ {
        dp[i] = make([]int, len(piles))
    }
    
    for interval := 0; interval < len(piles); interval++ {
        for i := 0; i + interval < len(piles); i++ {
            j := i + interval
            if i == j {
                dp[i][j] = piles[i]
            } else {
                dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
            }
        }
    }
    
    return dp[0][len(piles)-1] > 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
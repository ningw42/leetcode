func stoneGameIII(stoneValue []int) string {
    // dp[i] = the best score player can get from piles[i] to end
    // dp[i] = max(
    //             sum(piles, i, i) - dp[i+1], 
    //             sum(piles, i, i+1) - dp[i+2],
    //             sum(piles, i, i+2) - dp[i+3],
    //         )
    
    // padding stoneValue to get rid of edge cases
    n := len(stoneValue)
    stoneValue = append(stoneValue, 0, 0, 0)
    var dp []int
    dp = make([]int, len(stoneValue))
    for i := n - 1; i >= 0; i-- {
        sum := 0
        maxRelativeScore := math.MinInt32
        for j := i; j < i+3; j++ {
            sum += stoneValue[j]
            maxRelativeScore = max(maxRelativeScore, sum - dp[j+1])
        }
        dp[i] = maxRelativeScore
    }

    if dp[0] > 0 {
        return "Alice"
    } else if dp[0] < 0 {
        return "Bob"
    } else {
        return "Tie"
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
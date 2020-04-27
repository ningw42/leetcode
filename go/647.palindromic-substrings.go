func countSubstrings(s string) int {
    var dp [][]bool
    dp = make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    
    count := 0
    for i := 0; i < len(s); i++ {
        dp[i][i] = true
        count++
    }
    for i := 0; i < len(s) - 1; i++ {
        if s[i] == s[i+1] {
            dp[i][i+1] = true
            count++
        } else {
            dp[i][i+1] = false
        }
    }
    for interval := 2; interval < len(s); interval++ {
        for i := 0; i+interval < len(s); i++ {
            j := i + interval
            if s[i] == s[j] && dp[i+1][j-1] {
                dp[i][j] = true
                count++
            }
        }
    }
    
    return count
}
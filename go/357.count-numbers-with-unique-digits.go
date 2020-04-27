func countNumbersWithUniqueDigits(n int) int {
    // distinct[n] = # of numbers in [pow(10, n-1), pow(10, n))
    // dp[n] = # of numbers in [0, pow(10, n)) = sum(distinct[i]) for i = 1 to n
    distinct := make([]int, 11)
    distinct[1] = 10
    distinct[2] = 81 // 9*9
    for i := 8; i >= 1; i-- {
        distinct[11-i] = distinct[10-i] * i
    }
    dp := make([]int, 11)
    dp[0] = 1
    dp[1] = 10
    for i := 2; i <= 10; i++ {
        dp[i] = distinct[i] + dp[i-1]
    }
    if n <= 10 {
        return dp[n]
    } else {
        return dp[10]
    }
}
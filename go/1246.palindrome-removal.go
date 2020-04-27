func minimumMoves(arr []int) int {
    dp := make([][]int, len(arr))
    for i := 0; i < len(arr); i++ {
        dp[i] = make([]int, len(arr))
    }
  
    // length == 1
    for i := 0; i < len(arr); i++ {
        dp[i][i] = 1
    }
    
    // length == 2
    for i := 0; i < len(arr) - 1; i++ {
        if arr[i] == arr[i+1] {
            dp[i][i+1] = 1
        } else {
            dp[i][i+1] = 2
        }
    }
    
    // length >= 3
    for length := 3; length <= len(arr); length++ {
        for i := 0; i + length <= len(arr); i++ {
            j := i + length - 1
            if arr[i] == arr[j] {
                dp[i][j] = dp[i+1][j-1]
            } else {
                dp[i][j] = length
            }
            for k := i + 1; k <= j; k++ {
                dp[i][j] = min(dp[i][j], dp[i][k-1] + dp[k][j])
            }
        }
    }
    
    return dp[0][len(arr)-1]
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
func longestCommonSubsequence(text1 string, text2 string) int {
    // return backtracking(text1, text2)
    
    // m := make(map[string]int)
    // return backtrackingWithMemoization(text1, text2, m)
    
    return dynamicprogramming(text1, text2)
}

// O(M*N) where M = len(s1), N = len(s2)
// dp[i][j] = max(dp[i-1][j], dp[i][j-1]) for s1[i] != s2[j]
// dp[i][j] = max(dp[i-1][j-1]+1, dp[i-1][j], dp[i][j-1]) for s1[i] == s2[j]
func dynamicprogramming(s1, s2 string) int {
    if len(s1) == 0 || len(s2) == 0 {return 0}
    
    dp := make([][]int, len(s1))
    for i := 0; i < len(s1); i++ {
        dp[i] = make([]int, len(s2))
    }
    
    // init
    if s1[0] == s2[0] {
        dp[0][0] = 1
    } else {
        dp[0][0] = 0
    }
    for j := 1; j < len(s2); j++ {
        if dp[0][j-1] == 1 {
            dp[0][j] = 1
        } else {
            if s2[j] == s1[0] {
                dp[0][j] = 1
            }
        }
    }
    for i := 1; i < len(s1); i++ {
        if dp[i-1][0] == 1 {
            dp[i][0] = 1
        } else {
            if s1[i] == s2[0] {
                dp[i][0] = 1
            }
        }
    }
    
    for i := 1; i < len(s1); i++ {
        for j := 1; j < len(s2); j++ {
            if s1[i] == s2[j] {
                dp[i][j] = max(dp[i-1][j-1]+1, max(dp[i-1][j], dp[i][j-1]))
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    // fmt.Println(dp)
    return dp[len(s1)-1][len(s2)-1]
}

// better than naive backtracking, but still wont AC
func backtrackingWithMemoization(s1, s2 string, m map[string]int) int {
    ss := s1 + "," + s2
    if len(s1) == 0 || len(s2) == 0 {
        m[ss] = 0
        return 0
    }
    if result, exists := m[ss]; exists {
        return result
    }
    maxLength := math.MinInt32
    for i := 0; i < len(s1); i++ {
        for j := 0; j < len(s2); j++ {
            if s1[i] == s2[j] {
                maxLength = max(maxLength, backtrackingWithMemoization(s1[i+1:], s2[j+1:], m))
            }
        }
    }
    if maxLength == math.MinInt32 {
        m[ss] = 0
        return 0
    } else {
        m[ss] = maxLength + 1
        return maxLength + 1
    }
}

// wont AC
func backtracking(s1, s2 string) int {
    if len(s1) == 0 || len(s2) == 0 {
        return 0
    }
    maxLength := math.MinInt32
    for i := 0; i < len(s1); i++ {
        for j := 0; j < len(s2); j++ {
            if s1[i] == s2[j] {
                maxLength = max(maxLength, backtracking(s1[i+1:], s2[j+1:]))
            }
        }
    }
    if maxLength == math.MinInt32 {
        return 0
    } else {
        return maxLength + 1
    }
}

func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}
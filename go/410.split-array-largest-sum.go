func splitArray(nums []int, m int) int {
    return dynamicProgramming(nums, m)
}

func dynamicProgramming(nums []int, m int) int {
    // dp[i][j] = min max sum among i subarrays from nums[0] to nums[j] inclusive
    // dp[i][j] = min(max(dp[i-1][x], sum(nums, x+1, j))) for x from i - 2 to j - 1
    // sum(slice, from, to) = sum up 'from' to 'to' in 'nums' inclusive
    var dp [][]int
    dp = make([][]int, m+1)
    for i := 1; i <= m; i++ {
        dp[i] = make([]int, len(nums))
    }
    
    var sums []int
    var sum int
    sums = make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        sums[i] = sum
    }
   
    copy(dp[1], sums)
    for i := 2; i <= m; i++ {
        for j := 0; j < len(nums); j++ {
            minSum := math.MaxInt32
            for x := i - 2; x <= j - 1; x++ {
                minSum = min(minSum, max(dp[i-1][x], sums[j] - sums[x]))
            }
            dp[i][j] = minSum
        }
    }
    return dp[m][len(nums)-1]
}

// time complexity is identical to naiveDP, O(MN^3), N = len(nums), M = m
// space complexity is reduced to O(N^2), N = len(nums)
func spaceImprovedNaiveDP(nums []int, m int) int {
    sums := calculateSums(nums)
    // dp[k][i][j] = min max sum among k subarrays of nums[i] to nums[j] inclusive
    // dp[k][i][j] = min(max(dp[1][i][x], dp[k-1][x+1][j])) for x from i to j - k + 1
    // dp can be pruned into two dimensional dp[i][j]
    var dp [][]int
    dp = make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, len(nums))
        copy(dp[i], sums[i])
    }
    
    for k := 2; k <= m; k++ {
        for i := 0; i < len(nums); i++ {
            for j := i + k - 1; j < len(nums); j++ {
                minLargestSum := math.MaxInt32
                for x := i; x <= j - k + 1; x++ {
                    minLargestSum = min(minLargestSum, max(sums[i][x], dp[x+1][j]))
                }
                dp[i][j] = minLargestSum
            }
        }
    }
    return dp[0][len(nums)-1]
}

func naiveDP(nums []int, m int) int {
    // dp[k][i][j] = min max sum among k subarrays of nums[i] to nums[j] inclusive
    // dp[k][i][j] = min(max(dp[1][i][x], dp[k-1][x+1][j])) for x from i to j - k + 1
    var dp [][][]int
    dp = make([][][]int, m+1)
    for k := 1; k <= m; k++ {
        dp[k] = make([][]int, len(nums))
        for i := 0; i < len(nums); i++ {
            dp[k][i] = make([]int, len(nums))
        }
    }
   
    dp[1] = calculateSums(nums)
    for k := 2; k <= m; k++ {
        for i := 0; i < len(nums); i++ {
            for j := i + k - 1; j < len(nums); j++ {
                minLargestSum := math.MaxInt32
                for x := i; x <= j - k + 1; x++ {
                    minLargestSum = min(minLargestSum, max(dp[1][i][x], dp[k-1][x+1][j]))
                }
                dp[k][i][j] = minLargestSum
            }
        }
    }
    return dp[m][0][len(nums)-1]
}

func calculateSums(nums []int) [][]int {
    // sums[i][j] = sum of nums[i] to nums[j] inclusive
    var sums [][]int
    sums = make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        sums[i] = make([]int, len(nums))
        sums[i][i] = nums[i]
    }
    for interval := 1; interval < len(nums); interval++ {
        for i := 0; i + interval < len(nums); i++ {
            j := i + interval
            sums[i][j] = sums[i][j-1] + nums[j]
        }
    }
    return sums
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
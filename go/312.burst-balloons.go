func maxCoins(nums []int) int {
    // return dfs(nums)
    return dynamicProgramming(nums)
}

func dynamicProgramming(nums []int) int {
	// dp[i][j] = max score from i to j keep k reserved (burst i to k - 1 and k + 1 to j than k), we need to find a k with max score and store the result in dp[i][j]
	// dp[i][j] = max(dp[i][k-1] + dp[k+1][j] + nums[k] * nums[i-1] * nums[j+1]) for k = i to j
	n := len(nums)

	// padding nums, add 1 to both head and tail
	nums = append([]int{1}, append(nums, 1)...)

	var dp [][]int
	dp = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}

	for interval := 0; interval < n; interval++ {
		for i := 1; i+interval <= n; i++ {
			j := i + interval
			maxScore := 0
			for k := i; k <= j; k++ {
				maxScore = max(maxScore, dp[i][k-1]+dp[k+1][j]+nums[i-1]*nums[k]*nums[j+1])
			}
			dp[i][j] = maxScore
		}
	}

	return dp[1][n]
}

// wont AC
func dfs(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    } else {
        maxScore := math.MinInt32
        for i := 0; i < len(nums); i++ {
            product := 1
            for k := i - 1; k <= i + 1; k++ {
                if 0 <= k && k < len(nums) {
                    product *= nums[k]
                }
            }
            maxScore = max(maxScore, product + dfs(removeElement(nums, i)))
        }
        return maxScore
    }
}

func removeElement(s []int, i int) []int {
    d := make([]int, len(s)-1)
    k := 0
    for j := 0; j < len(s); j++ {
        if j == i {continue}
        d[k] = s[j]
        k++
    }
    return d
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
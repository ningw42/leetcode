func combinationSum4(nums []int, target int) int {
    return dynamicprogramming(nums, target)
}

func dynamicprogramming(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1
    for i := 1; i <= target; i++ {
        for _, num := range nums {
            if prev := i - num; prev >= 0 {
                dp[i] += dp[prev]
            }
        }
    }
    return dp[target]
}

func backtracking(nums []int, target int) int {
    if target == 0 {
        return 1
    } else if target < 0 {
        return 0
    } else {
        sum := 0
        for _, num := range nums {
            sum += backtracking(nums, target-num)
        }
        return sum
    }
}
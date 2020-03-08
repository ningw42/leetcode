/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    maxSubArraySumEndsAt := make([]int, len(nums))
    maxSubArraySumEndsAt[0] = nums[0]
    maxSum := nums[0]
    for i := 1; i < len(nums); i++ {
        if sum := maxSubArraySumEndsAt[i-1] + nums[i]; sum > nums[i] {
            maxSubArraySumEndsAt[i] = sum
        } else {
            maxSubArraySumEndsAt[i] = nums[i]
        }
        if maxSubArraySumEndsAt[i] > maxSum {
            maxSum = maxSubArraySumEndsAt[i]
        }
    }
    return maxSum
}


// @lc code=end


func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    maxSubArrayEndsAt := make([]int, len(nums))
    maxSubArrayEndsAt[0] = nums[0]
    maxSum := nums[0]
    for i := 1; i < len(nums); i++ {
        if sum := maxSubArrayEndsAt[i-1] + nums[i]; sum > nums[i] {
            maxSubArrayEndsAt[i] = sum
        } else {
            maxSubArrayEndsAt[i] = nums[i]
        }
        if maxSubArrayEndsAt[i] > maxSum {
            maxSum = maxSubArrayEndsAt[i]
        }
    }
    return maxSum
}

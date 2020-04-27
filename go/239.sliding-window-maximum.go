func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 {
        return nil
    }
    dp1 := make([]int, len(nums))
    dp2 := make([]int, len(nums))
    copy(dp1, nums)
    for i := 1; i < k; i++ {
        for j := i; j < len(nums); j++ {
            if nums[j] > dp1[j-1] {
                dp2[j] = nums[j]
            } else {
                dp2[j] = dp1[j-1]
            }
        }
        dp1, dp2 = dp2, dp1
    }
    return dp1[k-1:]
}
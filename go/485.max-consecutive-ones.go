func findMaxConsecutiveOnes(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    var maxOne, one int
    if nums[0] == 1 {
        one = 1
    }
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            if nums[i] == 1 {
                one++
            }
        } else {
            if nums[i] == 0 {
                maxOne = max(maxOne, one)
            } else {
                one = 1
            }
        }
    }
    maxOne = max(maxOne, one)
    return maxOne
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
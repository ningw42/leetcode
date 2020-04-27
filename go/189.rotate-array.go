func rotate(nums []int, k int)  {
    k = k % len(nums)
    for i := 0; i < k; i++ {
        shiftRight(nums)
    }
}

func shiftRight(nums []int) {
    t := nums[len(nums)-1]
    for i := len(nums) - 1; i > 0; i-- {
        nums[i] = nums[i-1]
    }
    nums[0] = t
}
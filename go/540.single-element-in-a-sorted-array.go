func singleNonDuplicate(nums []int) int {
    if len(nums) == 0 || len(nums) == 2 {
        return math.MaxInt32
    }
    if len(nums) == 1 {
        return nums[0]
    }
    left := 0
    right := len(nums) - 1
    mid := left + (right - left) / 2
    if nums[mid-1] == nums[mid] && nums[mid] < nums[mid+1] {
        lr := singleNonDuplicate(nums[:mid-1])
        if lr != math.MaxInt32 {
            return lr
        }
        return singleNonDuplicate(nums[mid+1:])
    } else if nums[mid-1] < nums[mid] && nums[mid] == nums[mid+1] {
        lr := singleNonDuplicate(nums[:mid])
        if lr != math.MaxInt32 {
            return lr
        }
        return singleNonDuplicate(nums[mid+2:])
    } else {
        return nums[mid]
    }
}
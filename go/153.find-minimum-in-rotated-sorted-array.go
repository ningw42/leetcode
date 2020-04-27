func findMin(nums []int) int {
    var left, right, mid int
    left = 0
    right = len(nums) - 1
    for left < right {
        mid = left + (right - left) / 2
        if nums[mid] < nums[right] {
            right = mid
        } else if nums[mid] > nums[right] {
            left = mid + 1
        } 
    }
    return nums[left]
}

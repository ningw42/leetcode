func findDuplicate(nums []int) int {
    var i, j int
    for nums[i] != nums[nums[j]] {
        i = nums[i]
        j = nums[nums[j]]
    }
    i = 0
    j = nums[nums[j]]
    for nums[i] != nums[j] {
        i = nums[i]
        j = nums[j]
    }
    return nums[i]
}
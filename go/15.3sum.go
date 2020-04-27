func threeSum(nums []int) [][]int {
    // sort
    sort.Ints(nums)
    var results [][]int
    for i := 0; i < len(nums); {
        target := -nums[i]
        j := i + 1
        k := len(nums) - 1
        for j < k {
            sum := nums[j] + nums[k]
            if sum > target {
                for j < k && nums[k] == nums[k-1] {k--}
                k--
            } else if sum < target {
                for j < k && nums[j] == nums[j+1] {j++}
                j++
            } else {
                results = append(results, []int{nums[i], nums[j], nums[k]})
                for j < k && nums[j] == nums[j+1] {j++}
                j++
                for j < k && nums[k] == nums[k-1] {k--}
                k--
            }
        }
        // skip duplicate 'i'
        for i < len(nums) - 1 && nums[i] == nums[i+1] {i++}
        i++
    }
    return results
}
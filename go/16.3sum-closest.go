func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums) // O(logn)
    var result, minDiff int
    minDiff = math.MaxInt32
    for i := 0; i < len(nums) - 2; i++ {
        if i == 0 || nums[i-1] != nums[i] {
            j := i + 1
            k := len(nums) - 1
            subTarget := target - nums[i]
            for j < k {
                sum := nums[j] + nums[k]
                if diff := DiffAbs(sum, subTarget); diff < minDiff {
                    minDiff = diff
                    result = nums[i] + nums[j] + nums[k]
                }
                if sum < subTarget {
                    for j < k && nums[j] == nums[j+1] {j++}
                    j++
                } else if sum > subTarget {
                    for j < k && nums[k] == nums[k-1] {k--}
                    k--
                } else {
                    return target
                }
            }
        }
    }
    return result
}

func DiffAbs(a, b int) int {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}
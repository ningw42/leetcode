func minSubArrayLen(s int, nums []int) int {
    var i, j, sum int
    length := math.MaxInt32
    for j < len(nums) {
        // expand left
        sum += nums[j]
        j++
        // shrink left
        for i < j && sum >= s {
            if l := j - i; l < length {
                length = l
            }
            sum -= nums[i]
            i++
        }
    }
    if length == math.MaxInt32 {
        return 0
    } else {
        return length
    }
}
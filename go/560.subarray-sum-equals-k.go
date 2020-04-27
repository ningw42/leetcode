func subarraySum(nums []int, k int) int {
    return HashMap(nums, k)
}

// TC O(n)
// SC O(n)
func HashMap(nums []int, k int) int {
    // map from 'sum' to '# of occurance of sum'
    sums := map[int]int{
        0: 1,
    }
   
    var sum, count int
    for _, num := range nums {
        sum += num
        if numberOfBegin, exists := sums[sum-k]; exists {
            count += numberOfBegin
        }
        sums[sum]++
    }
    return count
}

// TC O(n^2)
// SC O(n)
// SC can be optimized to O(1)
func NaiveDP(nums []int, k int) int {
    count := 0
    sums := make([]int, len(nums))
    for i, sum := range nums {
        sums[i] = sum
        if sum == k {
            count++
        }
    }
    for length := 2; length <= len(nums); length++ {
        for i := 0; i + length <= len(nums); i++ {
            sums[i] = sums[i] + nums[i+length-1]
            if sums[i] == k {
                count++
            }
        }
    }
    return count
}
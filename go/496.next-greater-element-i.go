func nextGreaterElement(nums1 []int, nums2 []int) []int {
    // monotone stack
    var stack []int
    pos := make(map[int]int)
    for i := len(nums2) - 1; i >= 0; i-- {
        if len(stack) == 0 {
            pos[nums2[i]] = -1
            stack = append(stack, nums2[i])
        } else {
            for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
                stack = stack[:len(stack)-1]
            }
            if len(stack) == 0 {
                pos[nums2[i]] = -1
            } else {
                pos[nums2[i]] = stack[len(stack)-1]
            }
            stack = append(stack, nums2[i])
        }
    }
    var ret []int
    for _, num := range nums1 {
        ret = append(ret, pos[num])
    }
    return ret
}